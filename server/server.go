package server

import (
	"fmt"
	"github.com/jaumeCloquellCapo/authGrpc/app/delivery"
	"github.com/jaumeCloquellCapo/authGrpc/app/service"
	grpc2 "github.com/jaumeCloquellCapo/authGrpc/grpc"
	"github.com/jaumeCloquellCapo/authGrpc/internal/dic"
	"github.com/jaumeCloquellCapo/authGrpc/internal/logger"
	"github.com/sarulabs/dingo/generation/di"
	"google.golang.org/grpc"
	"net"
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
	logger := logger.NewAPILogger()
	l, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		logger.Fatal(err)
		return err
	}
	defer l.Close()
	//grpcMiddleware := middleware.NewInterceptor(os.Getenv("ACCESS_SECRET"))
	//im := interceptors.NewInterceptorManager(logger)


	server := grpc.NewServer(
		//grpc.UnaryInterceptor(im.Logger),
		//grpc.UnaryInterceptor(grpcMiddleware.Auth),
	)
	ac := s.container.Get(dic.AuthService).(service.AuthServiceInterface)
	uc := s.container.Get(dic.UserService).(service.UserServiceInterface)

	authGRPCServer := delivery.NewUserServerGRPC(ac, uc, logger)
	grpc2.RegisterUserServiceServer(server, authGRPCServer)

	go func() {
		fmt.Print("server listen on port 8888")
		if err := server.Serve(l); err != nil {
			logger.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	server.GracefulStop()
	//s.logger.Info("Server Exited Properly")

	return nil
}
