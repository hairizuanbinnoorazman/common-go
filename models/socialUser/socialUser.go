package socialUser

import (
	"errors"
	"time"

	"golang.org/x/oauth2"
)

var ErrTokenUnavailable = errors.New("Unable to generate token from user.")

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

func (u *User) Validate() error {
	return nil
}

func (u User) GetToken() (*oauth2.Token, error) {
	if u.RefreshToken == "" || u.AccessToken == "" || u.TokenType == "" {
		return nil, ErrTokenUnavailable
	}
	token := oauth2.Token{}
	token.RefreshToken = u.RefreshToken
	token.AccessToken = u.AccessToken
	token.Expiry = u.Expiry
	token.TokenType = u.TokenType
	return &token, nil
}
