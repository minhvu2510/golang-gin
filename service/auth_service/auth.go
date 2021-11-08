package auth_service

import "github.com/minhvu2510/golang-gin/models"

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Check() (bool, error) {
	return models.CheckAuth(a.Username, a.Password) // call sang model de query db
}
