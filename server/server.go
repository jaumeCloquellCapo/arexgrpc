package server

import (
	"fmt"
	"github.com/jaumeCloquellCapo/authGrpc/app/delivery"
	"github.com/jaumeCloquellCapo/authGrpc/app/service"
	grpc2 "github.com/jaumeCloquellCapo/authGrpc/grpc"
	"github.com/jaumeCloquellCapo/authGrpc/internal/dic"
	"github.com/jaumeCloquellCapo/authGrpc/internal/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sarulabs/dingo/generation/di"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// GRPC Auth Server
type Server struct {
	container di.Container
}

// Server constructor
func NewAuthServer(container di.Container) *Server {
	return &Server{container: container}
}

// Run service
func (s *Server) Run() error {

	l, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		return err
	}
	defer l.Close()
	grpcMiddleware := middleware.NewInterceptor(os.Getenv("ACCESS_SECRET"))

	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpcMiddleware.Auth),
	)
	ac := s.container.Get(dic.AuthService).(service.AuthServiceInterface)
	uc := s.container.Get(dic.UserService).(service.UserServiceInterface)

	authGRPCServer := delivery.NewAuthServerGRPC(ac, uc)
	grpc2.RegisterUserServiceServer(server, authGRPCServer)

	http.Handle("/metrics", promhttp.Handler())

	go func() {
		fmt.Print("server listen on port 8888")
		//s.logger.Infof("Server is listening on port: %v", s.cfg.Server.Port)
		if err := server.Serve(l); err != nil {
			//s.logger.Fatal(err)
			fmt.Printf(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	server.GracefulStop()
	//s.logger.Info("Server Exited Properly")

	return nil
}
