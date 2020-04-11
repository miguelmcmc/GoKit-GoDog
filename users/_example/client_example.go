package main

import (
	"context"
	"fmt"

	client "ironchip.net/kit/users/client/http"
)

func main() {

	userCli, err := client.New("localhost:8081", nil)
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
