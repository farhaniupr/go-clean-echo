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
	DB *gorm.DB
}

// func ModuleDatabase(env Env) Database {
// 	db := ConnectAgain()
// 	return Database{
// 		DB: db,
// 	}
// }

func ModuleDatabase(env Env) Database {
	username := env.DBUsername
	password := env.DBPassword
	host := env.DBHost
	port := env.DBPort
	dbname := env.DBName

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Database connection established")

	return Database{
		DB: db,
	}
}

func (d *Database) ConnectAgain(env Env) {
	var err error
	username := env.DBUsername
	password := env.DBPassword
	host := env.DBHost
	port := env.DBPort
	dbname := env.DBName

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	log.Println(url)

	d.DB, err = gorm.Open(mysql.Open(url), &gorm.Config{
		// Logger: logger.GetGormLogger(),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection established")
}
