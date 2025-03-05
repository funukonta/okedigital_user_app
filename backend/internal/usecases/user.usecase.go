package usecases

import (
	"okedigital_user_app/backend/internal/dtos"
	"okedigital_user_app/backend/internal/models"
	"okedigital_user_app/backend/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UsecaseUser interface {
	CreateUser(req *dtos.UpsertUsersReq) error
	UpdateUser(id string, req *dtos.UpsertUsersReq) error
	DeleteUser(id string) error
	GetUser(id string) (*models.Users, error)
	GetUsers() ([]models.Users, error)
}

type usecaseUser struct {
	repoUser repositories.RepoUser
}

func NewUsecaseUser(repoUser repositories.RepoUser) UsecaseUser {
	return &usecaseUser{
		repoUser: repoUser,
	}
}

func (uc *usecaseUser) CreateUser(req *dtos.UpsertUsersReq) error {
	newUser := &models.Users{
		Username:  req.Username,
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser.Password = string(bcryptPassword)

	err = uc.repoUser.CreateUser(newUser)
	if err != nil {
		return err
	}

	return nil
}

func (uc *usecaseUser) UpdateUser(id string, req *dtos.UpsertUsersReq) error {
	user := &models.Users{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	if req.Password != "" {
		bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(bcryptPassword)
	}

	err := uc.repoUser.UpdateUser(id, user)
	if err != nil {
		return err
	}

	return nil
}

func (uc *usecaseUser) DeleteUser(id string) error {
	err := uc.repoUser.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

func (uc *usecaseUser) GetUser(id string) (*models.Users, error) {
	return nil, nil
}

func (uc *usecaseUser) GetUsers() ([]models.Users, error) {
	return nil, nil
}
