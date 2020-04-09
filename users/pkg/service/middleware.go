package service

import (
	"context"
	log "github.com/go-kit/kit/log"
	io "ironchip.net/kit/users/pkg/io"
)

// Middleware describes a service middleware.
type Middleware func(UsersService) UsersService

type loggingMiddleware struct {
	logger log.Logger
	next   UsersService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a UsersService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next UsersService) UsersService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) List(ctx context.Context) (users []io.User, err error) {
	defer func() {
		l.logger.Log("method", "List", "users", users, "err", err)
	}()
	return l.next.List(ctx)
}
func (l loggingMiddleware) Get(ctx context.Context, id string) (user io.User, err error) {
	defer func() {
		l.logger.Log("method", "Get", "id", id, "user", user, "err", err)
	}()
	return l.next.Get(ctx, id)
}
func (l loggingMiddleware) Create(ctx context.Context, user io.User) (err error) {
	defer func() {
		l.logger.Log("method", "Create", "user", user, "err", err)
	}()
	return l.next.Create(ctx, user)
}
func (l loggingMiddleware) Delete(ctx context.Context, id string) (err error) {
	defer func() {
		l.logger.Log("method", "Delete", "id", id, "err", err)
	}()
	return l.next.Delete(ctx, id)
}
