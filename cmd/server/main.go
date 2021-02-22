package main

import (
	"flag"
	"github.com/jaumeCloquellCapo/authGrpc/internal/dic"
	"github.com/jaumeCloquellCapo/authGrpc/server"
	"github.com/joho/godotenv"
	"log"
)

var config string

func main() {
	flag.StringVar(&config, "env", "dev.env", "help message for flagname")
	flag.Parse()

	if err := godotenv.Load(config); err != nil {
		log.Fatalf("Error loading %v", "dev.env")
	}
	container := dic.InitContainer()
	authServer := server.NewAuthServer(container)
	log.Fatal(authServer.Run())
}
