package grpc

import (
	"context"
	"errors"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc1 "github.com/go-kit/kit/transport/grpc"
	grpc "google.golang.org/grpc"
	endpoint1 "ironchip.net/kit/users/pkg/endpoint"
	pb "ironchip.net/kit/users/pkg/grpc/pb"
	service "ironchip.net/kit/users/pkg/service"
)

// New returns an AddService backed by a gRPC server at the other end
//  of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func New(conn *grpc.ClientConn, options map[string][]grpc1.ClientOption) (service.UsersService, error) {
	var listEndpoint endpoint.Endpoint
	{
		listEndpoint = grpc1.NewClient(conn, "pb.Users", "List", encodeListRequest, decodeListResponse, pb.ListReply{}, options["List"]...).Endpoint()
	}

	var getEndpoint endpoint.Endpoint
	{
		getEndpoint = grpc1.NewClient(conn, "pb.Users", "Get", encodeGetRequest, decodeGetResponse, pb.GetReply{}, options["Get"]...).Endpoint()
	}

	var createEndpoint endpoint.Endpoint
	{
		createEndpoint = grpc1.NewClient(conn, "pb.Users", "Create", encodeCreateRequest, decodeCreateResponse, pb.CreateReply{}, options["Create"]...).Endpoint()
	}

	var deleteEndpoint endpoint.Endpoint
	{
		deleteEndpoint = grpc1.NewClient(conn, "pb.Users", "Delete", encodeDeleteRequest, decodeDeleteResponse, pb.DeleteReply{}, options["Delete"]...).Endpoint()
	}

	return endpoint1.Endpoints{
		CreateEndpoint: createEndpoint,
		DeleteEndpoint: deleteEndpoint,
		GetEndpoint:    getEndpoint,
		ListEndpoint:   listEndpoint,
	}, nil
}

// encodeListRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain List request to a gRPC request.
func encodeListRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}

// decodeListResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeListResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeGetRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain Get request to a gRPC request.
func encodeGetRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}

// decodeGetResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeGetResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeCreateRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain Create request to a gRPC request.
func encodeCreateRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}

// decodeCreateResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeCreateResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeDeleteRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain Delete request to a gRPC request.
func encodeDeleteRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}

// decodeDeleteResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeDeleteResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}
