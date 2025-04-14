package users

import "time"

type Profile struct {
	ID                int64 `json:"id" db:"profile_id"`
	Account           Account
	Username          string     `json:"username" db:"username"`
	FristName         string     `json:"first_name" db:"first_name"`
	LastName          string     `json:"surname" db:"last_name"`
	Bio               string     `json:"bio" db:"bio"`
	ProfilePicutreURL string     `json:"profile_picture_url" db:"profile_picture_url"`
	CreatedAt         *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         *time.Time `json:"updated_at" db:"updated_at"`
}
