package main

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	grpc2 "github.com/jaumeCloquellCapo/authGrpc/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

func main() {
	ctx := context.Background()
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	email := gofakeit.Email()
	pass := "122"

	c := grpc2.NewUserServiceClient(conn)
	response, err := c.Register(ctx, &grpc2.RegisterRequest{
		Email:     email,
		FirstName: "123",
		LastName:  "32s1",
		Password:  pass,
	})
	if err != nil {
		log.Fatalf("Error when calling Register: %s", err)
	}

	response, err = c.Login(ctx, &grpc2.LoginRequest{
		Email:    email,
		Password: pass,
	})
	if err != nil {
		log.Fatalf("Error when calling Login: %s", err)
	}

	md := metadata.Pairs("authorization", response.AccessToke)
	ctx = metadata.NewOutgoingContext(context.Background(), md)

	response2, err2 := c.Logout(ctx, &grpc2.LogoutRequest{})
	if err2 != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	fmt.Print(response2)

}
