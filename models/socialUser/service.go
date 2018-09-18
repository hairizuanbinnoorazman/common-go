package socialUser

import "golang.org/x/oauth2"

type Service interface {
	GetByID(ID string) (User, error)
	GetTokenByID(ID string) (oauth2.Token, error)
	GetByEmail(Email string) (User, error)
	Create(*User) (User, error)
	Update(*User) (User, error)
}
