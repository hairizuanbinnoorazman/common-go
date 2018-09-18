package socialUser

import "time"

type User struct {
	ID           string    `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	RefreshToken string    `json:"-"`
	AccessToken  string    `json:"-"`
	Expiry       time.Time `json:"-"`
	TokenType    string    `json:"-"`
	LastLoginAt  time.Time `json:"-"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}
