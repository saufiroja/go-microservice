package repository

import (
	"github.com/saufiroja/auth-service/entity"
	"github.com/saufiroja/auth-service/interfaces"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) Register(user *entity.User) error {
	err := r.DB.Model(&user).Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Login(email string) (entity.User, error) {
	user := entity.User{}
	res := r.DB.Model(&user).Where("email = ?", email).First(&user)
	if res.Error != nil {
		return user, res.Error
	}

	return user, nil
}

func (r *userRepository) ValidateToken(email string) (entity.User, error) {
	user := entity.User{}
	res := r.DB.Model(&user).Where("email = ?", email).First(&user)
	if res.Error != nil {
		return user, res.Error
	}

	return user, nil
}
