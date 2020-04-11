package main

import (
	"context"
	"fmt"
	"reflect"

	"gopkg.in/square/go-jose.v2/json"
	userIO "ironchip.net/kit/users/pkg/io"

	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
	user "ironchip.net/kit/users/pkg/service"
)

type userService struct {
	service  user.UsersService
	ctx      context.Context
	response []byte
}

func NewUserService() *userService {
	us := &userService{}
	us.setup()
	return us
}

func (us *userService) setup() {
	us.service = user.NewBasicUsersService()
	us.ctx = context.Background()
}

func (us *userService) thereAreUsers(usersD *messages.PickleStepArgument_PickleTable) error {

	for i, row := range usersD.Rows {
		if i == 0 {
			continue
		}
		var user = userIO.User{}
		user.Name = row.Cells[0].Value
		user.Email = row.Cells[1].Value
		err := us.service.Create(us.ctx, user)
		if err != nil {
			return err
		}
	}

	return nil
}

func (us *userService) iListUsers() error {
	users, err := us.service.List(us.ctx)
	if err != nil {
		return err
	}
	fmt.Println("users: ", users)
	us.response, err = json.Marshal(users)
	if err != nil {
		return err
	}
	return nil
}

func (us *userService) theResponseShouldMatch(body *messages.PickleStepArgument_PickleDocString) error {
	var expected, actual interface{}

	// re-encode expected response
	if err := json.Unmarshal([]byte(body.Content), &expected); err != nil {
		return err
	}

	// re-encode actual response too
	if err := json.Unmarshal(us.response, &actual); err != nil {
		return err
	}

	// the matching may be adapted per different requirements.
	if !reflect.DeepEqual(expected, actual) {
		return fmt.Errorf("expected JSON does not match actual, %v vs. %v", expected, actual)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	us := NewUserService()
	s.Step(`^there are users:$`, us.thereAreUsers)
	s.Step(`^I list users$`, us.iListUsers)
	s.Step(`^the response should match:$`, us.theResponseShouldMatch)
}
