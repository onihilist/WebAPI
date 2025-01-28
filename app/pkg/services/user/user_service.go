package user

// CRUD

import (
	"github.com/onihilist/WebAPI/pkg/entities"
	userRepository "github.com/onihilist/WebAPI/pkg/repositories/user"
)

type UserService struct {
	UserRepo userRepository.UserRepository
}

func NewUserService(repo userRepository.UserRepository) UserService {
	return UserService{repo}
}

func (us *UserService) CreateUser(user entities.User) error {
	return us.UserRepo.CreateUser(user)
}

func (us *UserService) DeleteUser(username string) error {
	return us.UserRepo.DeleteUser(username)
}

func (us *UserService) GetUser(username string) (*entities.User, error) {
	return us.UserRepo.GetUser(username)
}
