package main

import (
	"context"
	grpc2 "github.com/jaumeCloquellCapo/authGrpc/grpc"
	"google.golang.org/grpc"
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

	c := grpc2.NewUserServiceClient(conn)
	response, err := c.Register(ctx, &grpc2.RegisterRequest{
		Email:     "jausssmes@d.com",
		FirstName: "123",
		LastName:  "321",
		Password:  "123456789",
	})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response)


}