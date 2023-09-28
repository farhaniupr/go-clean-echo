package services

import (
	"clean-go-echo/library"
	"clean-go-echo/models"
	"clean-go-echo/repository"
)

type User_MethodService interface {
	ListUser() ([]models.User, error)
}

type UserService struct {
	env  library.Env
	repo repository.UserRepository
}

func ModuleUserService(env library.Env, repoUser repository.UserRepository) User_MethodService {
	return UserService{
		env:  env,
		repo: repoUser,
	}
}

func (u UserService) ListUser() (user []models.User, err error) {

	sql, err := u.repo.DB.DB()

	if sql.Ping() != nil {
		u.repo.Database.ConnectAgain(u.env)
	}
	defer sql.Close()

	return user, u.repo.DB.Debug().Table("user").Select("*").Scan(&user).Error
}
