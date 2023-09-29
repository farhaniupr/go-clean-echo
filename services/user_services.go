package services

import (
	"clean-go-echo/library"
	"clean-go-echo/models"
	"clean-go-echo/repository"
)

type User_MethodService interface {
	ListUser() ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
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
	return user, u.repo.DB.Table("user").Select("*").Scan(&user).Error
}

func (u UserService) CreateUser(user models.User) (models.User, error) {
	return user, u.repo.DB.Table("user").Create(&user).Error
}
