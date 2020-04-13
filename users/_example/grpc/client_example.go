package main

import (
	"context"
	"fmt"

	grpc1 "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	client "ironchip.net/kit/users/client/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	userCli, err := client.New(connection, map[string][]grpc1.ClientOption{})
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	users, err := userCli.List(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("users: ", users)

}
