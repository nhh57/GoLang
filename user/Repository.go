package user

import (
	"gorm.io/gorm"
	"log"
)

type UserRepository interface {
	FindAll() ([]User, error)
	FindById(Id int) (User, error)
	Create(user User) (User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) FindById(Id int) (User, error) {
	var users User
	err := r.db.Find(&users, Id).Error
	return users, err
}

func (r *userRepository) Create(user User) (User, error) {

	log.Default().Printf("==========Start Create User Repo==========")
	err := r.db.Create(&user).Error
	log.Default().Printf("==========Start Create User Repo==========")
	return user, err
}
