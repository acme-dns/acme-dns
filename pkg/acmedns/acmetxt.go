package acmedns

import (
	"github.com/google/uuid"
)

func NewACMETxt() ACMETxt {
	var a = ACMETxt{}
	password := generatePassword(40)
	a.Username = uuid.New()
	a.Password = password
	a.Subdomain = uuid.New().String()
	return a
}
