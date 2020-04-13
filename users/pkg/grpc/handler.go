package grpc

import (
	"context"
	"errors"

	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
	endpoint "ironchip.net/kit/users/pkg/endpoint"
	pb "ironchip.net/kit/users/pkg/grpc/pb"
)

// makeListHandler creates the handler logic
func makeListHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.ListEndpoint, decodeListRequest, encodeListResponse, options...)
}

// decodeListResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain List request.
// TODO implement the decoder
func decodeListRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeListResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeListResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}
func (g *grpcServer) List(ctx context1.Context, req *pb.ListRequest) (*pb.ListReply, error) {
	_, rep, err := g.list.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ListReply), nil
}

// makeGetHandler creates the handler logic
func makeGetHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetEndpoint, decodeGetRequest, encodeGetResponse, options...)
}

// decodeGetResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Get request.
// TODO implement the decoder
func decodeGetRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeGetResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}
func (g *grpcServer) Get(ctx context1.Context, req *pb.GetRequest) (*pb.GetReply, error) {
	_, rep, err := g.get.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetReply), nil
}

// makeCreateHandler creates the handler logic
func makeCreateHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateEndpoint, decodeCreateRequest, encodeCreateResponse, options...)
}

// decodeCreateResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Create request.
// TODO implement the decoder
func decodeCreateRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeCreateResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCreateResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}
func (g *grpcServer) Create(ctx context1.Context, req *pb.CreateRequest) (*pb.CreateReply, error) {
	_, rep, err := g.create.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateReply), nil
}

// makeDeleteHandler creates the handler logic
func makeDeleteHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeleteEndpoint, decodeDeleteRequest, encodeDeleteResponse, options...)
}

// decodeDeleteResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Delete request.
// TODO implement the decoder
func decodeDeleteRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeDeleteResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeleteResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}
func (g *grpcServer) Delete(ctx context1.Context, req *pb.DeleteRequest) (*pb.DeleteReply, error) {
	_, rep, err := g.delete.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteReply), nil
}
