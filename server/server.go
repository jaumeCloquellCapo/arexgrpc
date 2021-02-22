package server

import (
	"github.com/jaumeCloquellCapo/authGrpc/app/handler"
	"github.com/jaumeCloquellCapo/authGrpc/app/service"
	grpc2 "github.com/jaumeCloquellCapo/authGrpc/grpc"
	"github.com/jaumeCloquellCapo/authGrpc/internal/dic"
	"github.com/jaumeCloquellCapo/authGrpc/internal/interceptors"
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
	logger.InitLogger()

	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		logger.Fatal(err)
		return err
	}
	defer l.Close()

	im := interceptors.NewInterceptorManager(logger)
	grpc.UnaryInterceptor(im.Logger)

	server := grpc.NewServer(grpc.UnaryInterceptor(im.Logger))
	ac := s.container.Get(dic.AuthService).(service.AuthServiceInterface)
	uc := s.container.Get(dic.UserService).(service.UserServiceInterface)

	authGRPCServer := handler.NewServerGRPC(ac, uc, logger)
	grpc2.RegisterUserServiceServer(server, authGRPCServer)

	go func() {
		logger.Info("server listen on port 8888")
		if err := server.Serve(l); err != nil {
			logger.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	server.GracefulStop()

	return nil
}
