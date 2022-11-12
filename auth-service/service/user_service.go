package service

import (
	"github.com/saufiroja/auth-service/entity"
	"github.com/saufiroja/auth-service/interfaces"
	"github.com/saufiroja/auth-service/utils"
)

type UserService struct {
	UserRepository interfaces.UserRepository
}

func NewUserService(userRepository interfaces.UserRepository) interfaces.UserService {
	return &UserService{UserRepository: userRepository}
}

func (s *UserService) Register(user *entity.User) error {
	hash, _ := utils.HashPassword(user.Password)
	user.Password = hash
	return s.UserRepository.Register(user)
}

func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.UserRepository.Login(email)
	if err != nil {
		return "", err
	}
	utils.ComparePassword(user.Password, password)

	accessToken, _ := utils.GenerateAccessToken(user.ID, user.Email)

	return accessToken, nil
}
