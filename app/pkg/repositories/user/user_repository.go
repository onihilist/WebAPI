package repositories

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"time"

	"github.com/onihilist/WebAPI/pkg/entities"
	"github.com/onihilist/WebAPI/pkg/utils"
)

// UserRepository provides methods to interact with the user data.
type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CreateUser Profile creates a new user profile in the database.
func (ur *UserRepository) CreateUser(user entities.User) error {
	hash := md5.Sum([]byte(user.Password))
	hashString := hex.EncodeToString(hash[:])

	var req string
	if user.Phone != nil {
		req = `INSERT INTO users (permission_id, username, password, email, phone, creationDate, lastConnection, lastIP) VALUES (?, ?, ?, ?, ?, ?, ?, ?);`
		_, err := ur.DB.Exec(req,
			user.PermissionID,
			user.Username,
			hashString,
			user.Email,
			user.Phone,
			time.Now(),
			time.Now(),
			user.LastIP,
		)
		return err
	} else {
		req = `INSERT INTO users (permission_id, username, password, email, creationDate, lastConnection, lastIP) VALUES (?, ?, ?, ?, ?, ?, ?);`
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

	utils.LogWarning("test")
	row := ur.DB.QueryRow(query, username)
	utils.LogWarning("test finito")

	var user entities.User
	var creationDateBytes, lastConnectionBytes []byte

	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Phone, &creationDateBytes, &lastConnectionBytes, &user.LastIP)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	// Convertir les []byte en time.Time
	if len(creationDateBytes) > 0 {
		user.CreationDate, err = time.Parse("2006-01-02 15:04:05", string(creationDateBytes))
		if err != nil {
			return nil, err
		}
	}

	if len(lastConnectionBytes) > 0 {
		user.LastConnection, err = time.Parse("2006-01-02 15:04:05", string(lastConnectionBytes))
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (ur *UserRepository) DeleteUser(username string) error {
	req := `DELETE FROM users WHERE username = ?;`
	_, err := ur.DB.Exec(req, username)
	return err
}

func (ur *UserRepository) GetPasswordByUsername(username string) (string, error) {
	var password string
	err := ur.DB.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("user not found")
		}
		return "", err
	}
	return password, nil
}

func (ur *UserRepository) GetUsersByPermission(idPermission int) (*entities.User, error) {
	var user entities.User
	err := ur.DB.QueryRow("SELECT * FROM users WHERE permission_id = ?", idPermission).Scan(&user)
	if err != nil {
		if err == sql.ErrNoRows {
			return &user, errors.New("user not found")
		}
		return &user, err
	}
	return &user, nil
}

func (ur *UserRepository) GetUserBySessionID(sessionID interface{}) (entities.User, error) {
	var user entities.User
	err := ur.DB.QueryRow("SELECT username, email, phone, permission_id FROM users WHERE session_id = ?", sessionID).Scan(&user.Username, &user.Email, &user.Phone, &user.PermissionID)
	return user, err
}

func (ur *UserRepository) GetPermissionByID(permissionID int) (string, error) {
	var permission string
	err := ur.DB.QueryRow("SELECT permission FROM permissions WHERE id = ?", permissionID).Scan(&permission)
	return permission, err
}

func (ur *UserRepository) UpdateSessionCookie(session interface{}) (sql.Result, error) {
	return ur.DB.Exec(`UPDATE users SET session_id=NULL WHERE session_id=?`, session)
}
