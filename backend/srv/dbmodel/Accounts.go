package dbmodel

import "time"

type AccountRole struct {
	ID          int64  `json:"id" db:"role_id"`
	Name        string `json:"name" db:"role_name"`
	Description string `json:"description" db:"description"`
}

type Account struct {
	ID        int64      `json:"id" db:"account_id"`
	Email     string     `json:"email" db:"email"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	RoleID    int64      `json:"role_id" db:"role_id"`
	IsActive  bool       `json:"is_active" db:"is_active"`
}
