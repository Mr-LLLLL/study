package proxy

import (
	"log"
	"time"
)

type IUser interface {
	Login(username, password string) error
}

type User struct{}

func (u *User) Login(username, password string) error {
	return nil
}

type UserProxy struct {
	user *User
}

func (u *UserProxy) Login(username, password string) error {
	start := time.Now()

	if err := u.user.Login(username, password); err != nil {
		return err
	}

	log.Printf("user login cost time: %s", time.Now().Sub(start))
	return nil
}

func NewUserProxy(user *User) *UserProxy {
	return new(UserProxy)
}
