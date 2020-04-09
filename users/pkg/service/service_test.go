package service

import (
	"context"
	"testing"
)

func Test_basicUsersService_List(t *testing.T) {
	srv, ctx := setup()

	users, err := srv.List(ctx)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	t.Log("users: ",users)

}

func setup() (srv UsersService, ctx context.Context) {
	return NewBasicUsersService(), context.Background()
}
