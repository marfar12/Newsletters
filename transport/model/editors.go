package model

import (
	svcmodel "newsletter/service/model"
	"os"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
)

var (
	TokenAuth *jwtauth.JWTAuth
)

func init() {
	TokenAuth = jwtauth.New("HS256", []byte(os.Getenv("SECRET")), nil)
}

type Editor struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email" validate:"email"`
	Password string    `json:"password,omitempty"`
}

func ToSvcEditor(u Editor) svcmodel.Editor {
	return svcmodel.Editor{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
	}
}

func ToNetEditor(u svcmodel.Editor) Editor {
	return Editor{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
	}
}

func ToNetEditors(editors []svcmodel.Editor) []Editor {
	netEditors := make([]Editor, 0, len(editors))
	for _, editor := range editors {
		netEditors = append(netEditors, ToNetEditor(editor))
	}
	return netEditors
}

type Token struct {
	Token string `json:"token"`
}

func CreateToken(editor Editor) Token {
	claims := map[string]interface{}{"ID": editor.ID}
	jwtauth.SetExpiry(claims, time.Now().Add(time.Hour))
	_, tokenString, _ := TokenAuth.Encode(claims)
	return Token{Token: tokenString}
}
