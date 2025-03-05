package repositories

import (
	"okedigital_user_app/backend/internal/models"

	"gorm.io/gorm"
)

type RepoUser interface {
	CreateUser(*models.Users) error
	UpdateUser(id string, user *models.Users) error
	DeleteUser(id string) error
	GetUser(id string, userWhere models.UserWhere) (*models.Users, error)
	GetUsers() ([]models.Users, error)
}

type repoUser struct {
	dB *gorm.DB
}

func NewRepoUser() RepoUser {
	return &repoUser{}
}

func (r *repoUser) CreateUser(user *models.Users) error {
	return r.dB.Create(user).Error
}

func (r *repoUser) UpdateUser(id string, user *models.Users) error {
	return r.dB.Model(&models.Users{}).
		Where("id = ?", id).
		Updates(user).Error
}

func (r *repoUser) DeleteUser(id string) error {
	return r.dB.Where("id = ?", id).
		Delete(&models.Users{}).Error
}

func (r *repoUser) GetUser(id string, userWhere models.UserWhere) (*models.Users, error) {
	user := &models.Users{}
	query := r.dB
	populateUserWithWhere(query, userWhere)

	err := query.Find(user).Error
	return user, err
}

func (r *repoUser) GetUsers() ([]models.Users, error) {
	users := []models.Users{}
	err := r.dB.Find(&users).Error
	return users, err
}

func populateUserWithWhere(query *gorm.DB, userWhere models.UserWhere) {
	if userWhere.ID != "" {
		query.Where("id = ?", userWhere.ID)
	}
	if userWhere.Username != "" {
		query.Where("username = ?", userWhere.Username)
	}
}
