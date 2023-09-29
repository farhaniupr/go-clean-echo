package repository

import (
	"clean-go-echo/library"
	"log"

	"gorm.io/gorm"
)

type UserRepository struct {
	library.Database
}

func ModuleUserRepository(db library.Database, env library.Env) UserRepository {
	return UserRepository{
		Database: db,
	}
}

// WithTrx enables repository with transaction
func (r UserRepository) WithTrx(trxHandle *gorm.DB) UserRepository {
	if trxHandle == nil {
		log.Println("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}
