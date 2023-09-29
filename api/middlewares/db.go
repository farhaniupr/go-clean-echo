package middlewares

import (
	"clean-go-echo/constants"
	"clean-go-echo/library"
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// DatabaseTrx middleware for transactions support for database
type DatabaseTrx struct {
	handler library.RequestHandler
	db      library.Database
	env     library.Env
}

// statusInList function checks if context writer status is in provided list
func statusInList(status int, statusList []int) bool {
	for _, i := range statusList {
		if i == status {
			return true
		}
	}
	return false
}

// ModuleDatabase creates new database transactions middleware
func ModuleDatabase(
	handler library.RequestHandler,
	db library.Database,
	env library.Env,
) DatabaseTrx {
	return DatabaseTrx{
		handler: handler,
		db:      db,
		env:     env,
	}
}

// Setup sets up database transaction middleware
func (m DatabaseTrx) Setup() {

	m.handler.Echo.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			log.Println("beginning database transaction")

			m.db.ConnectAgain(m.env)

			txHandle := m.db.DB.Begin()

			defer func() {
				if r := recover(); r != nil {
					txHandle.Rollback()

					sql := CloseDb(m.db.DB)
					defer sql.Close()
				}
			}()

			c.Set(constants.DBTransaction, txHandle)
			if err := next(c); err != nil {
				log.Println("commit err : ", err.Error())
				return err
			}

			// commit transaction on success status
			if statusInList(c.Response().Status, []int{http.StatusOK, http.StatusCreated, http.StatusNoContent, http.StatusPermanentRedirect, http.StatusTemporaryRedirect}) {

				log.Println("commit database transaction")
				if err := txHandle.Commit().Error; err != nil {
					log.Println("commit err : ", err.Error())

					sql := CloseDb(m.db.DB)
					defer sql.Close()

					return err
				}
				sql := CloseDb(m.db.DB)
				defer sql.Close()

			} else {

				log.Println("rolling back database transaction")
				txHandle.Rollback()

				sql := CloseDb(m.db.DB)
				defer sql.Close()

				return nil

			}

			return nil
		}
	})
}

func CloseDb(m *gorm.DB) *sql.DB {
	sql, err := m.DB()
	if err != nil {
		log.Println(err.Error())
	}
	return sql
}
