package main

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/transport/http"
	client "ironchip.net/kit/users/client/http"
)

func main() {

	userCli, err := client.New("http://localhost:8081", map[string][]http.ClientOption{})
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	user, err := userCli.Get(ctx, "10")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("user: ", user)
	users, err := userCli.List(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("users: ", users)

}
