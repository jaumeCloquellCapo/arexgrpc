package delivery

import (
	"context"
	"github.com/jaumeCloquellCapo/authGrpc/app/model"
	"github.com/jaumeCloquellCapo/authGrpc/grpc"
	"github.com/jaumeCloquellCapo/authGrpc/internal/logger"
	"github.com/jaumeCloquellCapo/authGrpc/mock"
	"testing"
)

func TestMicroservice_Register(t *testing.T) {
	t.Parallel()
	us := &mock.MockUserService{}
	au := &mock.MockAuthService{}
	logger := logger.NewAPILogger()
	userUC := NewUserServerGRPC(au,us, logger)

	mockUser := &model.User{
		ID:         0,
		Name:       "",
		LastName:   "LastName",
		Password:   "123456",
		Email:      "email@gmail.com",
		Country:    "",
		Phone:      "",
		PostalCode: "",
	}

	ctx := context.Background()

	userUC.Register(ctx, &grpc.RegisterRequest{
		Email:     mockUser.Email,
		FirstName: mockUser.Name,
		LastName:  mockUser.LastName,
		Password:  mockUser.Password,
	})
}

