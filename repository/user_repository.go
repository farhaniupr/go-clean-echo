package repository

import "clean-go-echo/library"

type UserRepository struct {
	library.Env
	library.Database
}

func ModuleUserRepository(db library.Database, env library.Env) UserRepository {
	return UserRepository{
		Database: db,
		Env:      env,
	}
}
