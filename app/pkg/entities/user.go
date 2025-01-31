package entities

import (
	"time"
)

type User struct {
	ID             *uint     `json:"id,omitempty"`
	PermissionID   uint      `json:"permission_id"`
	Username       string    `json:"username"`
	Password       string    `json:"-"`
	Email          string    `json:"email"`
	Phone          *string   `json:"phone,omitempty"`
	CreationDate   time.Time `json:"creationDate"`
	LastConnection time.Time `json:"lastConnection"`
	LastIP         string    `json:"lastIP"`
	SessionID      *string   `json:"session_id,omitempty"`
	AvatarURL      *string   `json:"avatar_url,omitempty"`
}
