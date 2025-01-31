package user

// CRUD

import (
	"database/sql"
	"errors"

	"github.com/onihilist/WebAPI/pkg/entities"
	userRepository "github.com/onihilist/WebAPI/pkg/repositories/user"
)

type UserService struct {
	UserRepo *userRepository.UserRepository
}

func NewUserService(repo *userRepository.UserRepository) *UserService {
	return &UserService{repo}
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

func (us *UserService) GetUserBySessionID(sessionID interface{}) (entities.User, error) {
	user, err := us.UserRepo.GetUserBySessionID(sessionID)
	if err != nil {
		return entities.User{}, err
	}

	return user, err
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

func (us *UserService) UpdateSessionCookie(sessionID interface{}, username string) (sql.Result, error) {
	return us.UserRepo.UpdateSessionCookie(sessionID, username)
}

func (us *UserService) DeleteSessionCookie(sessionID interface{}) (sql.Result, error) {
	return us.UserRepo.DeleteSessionCookie(sessionID)
}

func (us *UserService) UploadAvatar(username string, filePath string) (sql.Result, error) {
	return us.UserRepo.UploadAvatar(username, filePath)
}

func (us *UserService) DeleteAvatar(username string) (string, error) {
	return us.UserRepo.DeleteAvatar(username)
}

func (us *UserService) UpdateUsername(username string, sessionID interface{}) (sql.Result, error) {
	return us.UserRepo.UpdateUsername(username, sessionID)
}

func (us *UserService) UpdatePassword(password string, sessionID interface{}) (sql.Result, error) {
	return us.UserRepo.UpdatePassword(password, sessionID)
}

func (us *UserService) UpdateEmail(email string, sessionID interface{}) (sql.Result, error) {
	return us.UserRepo.UpdateEmail(email, sessionID)
}
