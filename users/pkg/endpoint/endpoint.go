package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	io "ironchip.net/kit/users/pkg/io"
	service "ironchip.net/kit/users/pkg/service"
)

// ListRequest collects the request parameters for the List method.
type ListRequest struct{}

// ListResponse collects the response parameters for the List method.
type ListResponse struct {
	Users []io.User `json:"users"`
	Err   error     `json:"err"`
}

// MakeListEndpoint returns an endpoint that invokes List on the service.
func MakeListEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		users, err := s.List(ctx)
		return ListResponse{
			Err:   err,
			Users: users,
		}, nil
	}
}

// Failed implements Failer.
func (r ListResponse) Failed() error {
	return r.Err
}

// GetRequest collects the request parameters for the Get method.
type GetRequest struct {
	Id string `json:"id"`
}

// GetResponse collects the response parameters for the Get method.
type GetResponse struct {
	User io.User `json:"user"`
	Err  error   `json:"err"`
}

// MakeGetEndpoint returns an endpoint that invokes Get on the service.
func MakeGetEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		user, err := s.Get(ctx, req.Id)
		return GetResponse{
			Err:  err,
			User: user,
		}, nil
	}
}

// Failed implements Failer.
func (r GetResponse) Failed() error {
	return r.Err
}

// CreateRequest collects the request parameters for the Create method.
type CreateRequest struct {
	User io.User `json:"user"`
}

// CreateResponse collects the response parameters for the Create method.
type CreateResponse struct {
	Err error `json:"err"`
}

// MakeCreateEndpoint returns an endpoint that invokes Create on the service.
func MakeCreateEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		err := s.Create(ctx, req.User)
		return CreateResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r CreateResponse) Failed() error {
	return r.Err
}

// DeleteRequest collects the request parameters for the Delete method.
type DeleteRequest struct {
	Id string `json:"id"`
}

// DeleteResponse collects the response parameters for the Delete method.
type DeleteResponse struct {
	Err error `json:"err"`
}

// MakeDeleteEndpoint returns an endpoint that invokes Delete on the service.
func MakeDeleteEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		err := s.Delete(ctx, req.Id)
		return DeleteResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r DeleteResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// List implements Service. Primarily useful in a client.
func (e Endpoints) List(ctx context.Context) (users []io.User, err error) {
	request := ListRequest{}
	response, err := e.ListEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ListResponse).Users, response.(ListResponse).Err
}

// Get implements Service. Primarily useful in a client.
func (e Endpoints) Get(ctx context.Context, id string) (user io.User, err error) {
	request := GetRequest{Id: id}
	response, err := e.GetEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetResponse).User, response.(GetResponse).Err
}

// Create implements Service. Primarily useful in a client.
func (e Endpoints) Create(ctx context.Context, user io.User) (err error) {
	request := CreateRequest{User: user}
	response, err := e.CreateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateResponse).Err
}

// Delete implements Service. Primarily useful in a client.
func (e Endpoints) Delete(ctx context.Context, id string) (err error) {
	request := DeleteRequest{Id: id}
	response, err := e.DeleteEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteResponse).Err
}
