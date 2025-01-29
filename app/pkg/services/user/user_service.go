package user

// CRUD

import (
	"database/sql"
	"errors"

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

func (us *UserService) GetUserBySessionID(sessionID interface{}) (entities.User, string, error) {
	user, err := us.UserRepo.GetUserBySessionID(sessionID)
	if err != nil {
		return entities.User{}, "", err
	}

	permission, err := us.UserRepo.GetPermissionByID(int(user.PermissionID))
	if err != nil {
		return user, "", err
	}

	return user, permission, nil
}

func (us *UserService) LoginAdmin(username, password string) (bool, error) {
	storedPassword, err := us.UserRepo.GetPasswordByUsername(username)
	if err != nil {
		return false, err
	}

	if storedPassword != password {
		return false, errors.New("invalid username or password")
	}

	return true, nil
}

func (us *UserService) GetUsersByPermission(idPermission int) (*entities.User, error) {
	return us.UserRepo.GetUsersByPermission(idPermission)
}

func (us *UserService) UpdateSessionCookie(sessionID interface{}) (sql.Result, error) {
	return us.UserRepo.UpdateSessionCookie(sessionID)
}
