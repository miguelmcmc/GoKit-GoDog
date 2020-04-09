package service

import (
	"context"
	"errors"

	"ironchip.net/kit/users/pkg/db"
	"ironchip.net/kit/users/pkg/io"
)

// UsersService describes the service.
type UsersService interface {
	List(ctx context.Context) (users []io.User, err error)
	Get(ctx context.Context, id string) (user io.User, err error)
	Create(ctx context.Context, user io.User) (err error)
	Delete(ctx context.Context, id string) (err error)
}

var (
	ErrUserNotFound = errors.New("User not found")
)

type basicUsersService struct{}

func (b *basicUsersService) List(ctx context.Context) (users []io.User, err error) {
	users = db.UsersDB
	return users, err
}
func (b *basicUsersService) Get(ctx context.Context, id string) (user io.User, err error) {
	users := db.UsersDB
	for _, userDB := range users {
		if userDB.ID == id {
			return userDB, nil
		}
	}
	return io.User{}, ErrUserNotFound
}
func (b *basicUsersService) Create(ctx context.Context, user io.User) (err error) {
	db.UsersDB = append(db.UsersDB, user)
	return err
}
func (b *basicUsersService) Delete(ctx context.Context, id string) (err error) {
	users := db.UsersDB
	for i, userDB := range users {
		if userDB.ID == id {
			db.UsersDB[i] = db.UsersDB[len(db.UsersDB)-1]
			db.UsersDB[len(db.UsersDB)-1] = io.User{}
			db.UsersDB = db.UsersDB[:len(db.UsersDB)-1]
			return nil
		}
	}
	return ErrUserNotFound
}

// NewBasicUsersService returns a naive, stateless implementation of UsersService.
func NewBasicUsersService() UsersService {
	return &basicUsersService{}
}

// New returns a UsersService with all of the expected middleware wired in.
func New(middleware []Middleware) UsersService {
	var svc UsersService = NewBasicUsersService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
