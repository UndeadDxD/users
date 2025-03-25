package repository

import (
	"errors"
	"gorm.io/gorm"
	"users/internal/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetById(id int) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}

type userRepository struct {
	db *gorm.DB
}

type UserModel = models.User

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
func (r *userRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetById(id int) (*models.User, error) {
	{
		var user models.User
		if err := r.db.First(&user, id).Error; err != nil {
			return nil, err
		}
		return &user, nil
	}
}

func (r *userRepository) DeleteUser(id int) error {
	return r.db.Delete(&models.User{}, id).Error
}
func (r *userRepository) UpdateUser(user *models.User) error {
	var existing models.User
	if err := r.db.First(&existing, user.ID).Error; err != nil {
		return errors.New("user not found")
	}
	return r.db.Model(&existing).Updates(user).Error
}
