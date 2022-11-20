package interfaces

import (
	"github.com/saufiroja/auth-service/entity"
)

type UserRepository interface {
	Register(user *entity.User) error
	Login(email string) (entity.User, error)
	ValidateToken(email string) (entity.User, error)
}

type UserService interface {
	Register(user *entity.User) error
	Login(email, password string) (string, error)
	ValidateToken(token string) error
}
