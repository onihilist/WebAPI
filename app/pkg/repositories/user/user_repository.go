package repositories

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"time"

	"github.com/onihilist/WebAPI/pkg/entities"
)

// UserRepository provides methods to interact with the user data.
type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{db}
}

// CreateUser Profile creates a new user profile in the database.
func (ur *UserRepository) CreateUser(user entities.User) error {
	hash := md5.Sum([]byte(user.Password))
	hashString := hex.EncodeToString(hash[:])

	var req string
	if user.Phone != nil {
		req = `INSERT INTO users (permissionId, username, password, email, phone, creationDate, lastConnection, lastIP) VALUES (?, ?, ?, ?, ?, ?, ?, ?);`
		_, err := ur.DB.Exec(req,
			user.PermissionID,
			user.Username,
			hashString,
			user.Email,
			*user.Phone,
			time.Now(),
			time.Now(),
			user.LastIP,
		)
		return err
	} else {
		req = `INSERT INTO users (permissionId, username, password, email, creationDate, lastConnection, lastIP) VALUES (?, ?, ?, ?, ?, ?, ?);`
		_, err := ur.DB.Exec(req,
			user.PermissionID,
			user.Username,
			hashString,
			user.Email,
			time.Now(),
			time.Now(),
			user.LastIP,
		)
		return err
	}
}

func (ur *UserRepository) GetUser(username string) (*entities.User, error) {
	query := "SELECT id, username, password, email, phone, creationDate, lastConnection, lastIP FROM users WHERE username = ?"

	row := ur.DB.QueryRow(query, username)

	var user entities.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Phone, &user.CreationDate, &user.LastConnection, &user.LastIP)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) DeleteUser(username string) error {
	req := `DELETE FROM users WHERE username = ?;`
	_, err := ur.DB.Exec(req, username)
	return err
}
