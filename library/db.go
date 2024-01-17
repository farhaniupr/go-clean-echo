package library

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Database modal
type Database struct {
	*gorm.DB
}

// ModuleDatabase initial
func ModuleDatabase(env Env) Database {
	return Database{
		DB: ConnectDB(env),
	}
}

// ConnectAgain reconnect
func (d *Database) ConnectAgain(env Env) {

	var err error

	d.DB, err = gorm.Open(mysql.Open(configDb(env)), &gorm.Config{
		// Logger: logger.GetGormLogger(),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Println(err.Error())
	}

	log.Println("reconnect database")

}

func configDb(env Env) string {
	username := env.DBUsername
	password := env.DBPassword
	host := env.DBHost
	port := env.DBPort
	dbname := env.DBName

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)
}

func ConnectDB(env Env) *gorm.DB {
	username := env.DBUsername
	password := env.DBPassword
	host := env.DBHost
	port := env.DBPort
	dbname := env.DBName

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		// Logger: logger.GetGormLogger(),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Println(err.Error())
	}

	return db
}
