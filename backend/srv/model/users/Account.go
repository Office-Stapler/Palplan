package users

import "time"

type AccountRole struct {
	ID          int64  `json:"id" db:"role_id"`
	Name        string `json:"name" db:"role_name"`
	Description string `json:"description" db:"description"`
}

type Account struct {
	ID        int64        `json:"id"`
	Email     string       `json:"email"`
	CreatedAt *time.Time   `json:"created_at"`
	UpdatedAt *time.Time   `json:"updated_at"`
	Role      *AccountRole `json:"role"`
	IsActive  bool         `jsnon:"is_active"`
}
