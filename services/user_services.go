package services

import (
	"clean-go-echo/library"
	"clean-go-echo/models"
	"clean-go-echo/repository"
)

type User_MethodService interface {
	ListUser(limit int) ([]models.User, error)
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

func (u UserService) ListUser(limit int) (user []models.User, err error) {

	repository := library.ConnectDB(u.env)
	sql, err := repository.DB()
	if err != nil {
		return []models.User{}, err
	}

	defer sql.Close()

	if limit > 0 {
		return user, repository.Table("user").Limit(limit).Select("*").Scan(&user).Error
	} else {
		return user, repository.Table("user").Select("*").Scan(&user).Error
	}
}

func (u UserService) CreateUser(user models.User) (models.User, error) {

	repository := library.ConnectDB(u.env)
	sql, err := repository.DB()
	if err != nil {
		return models.User{}, err
	}

	defer sql.Close()

	return user, repository.Table("user").Create(&user).Error
}
