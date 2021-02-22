package handler

import (
	"context"
	"github.com/golang/mock/gomock"
	models "github.com/jaumeCloquellCapo/authGrpc/app/model"
	"github.com/jaumeCloquellCapo/authGrpc/grpc"
	"github.com/jaumeCloquellCapo/authGrpc/internal/logger"
	"github.com/jaumeCloquellCapo/authGrpc/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMicroservice_Login(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUC := mock.NewMockUserServiceCase(ctrl)
	sessUC := mock.NewMockAuthServiceCase(ctrl)
	apiLogger := logger.NewAPILogger()
	apiLogger.InitLogger()

	authServerGRPC := NewServerGRPC(sessUC,userUC, apiLogger)

	reqValue := &grpc.LoginRequest{
		Email:    "email@gmail.com",
		Password: "Password",
	}

	t.Run("Login", func(t *testing.T) {
		t.Parallel()
		credentials := models.Credentials{
			Email:    reqValue.Email,
			Password: reqValue.Password,
		}
		tokenDetails := models.TokenDetails{
			AccessToken:  "1",
			RefreshToken: "1",
			AccessUUID:   "1",
			RefreshUUID:  "1",
			AtExpires:    0,
			RtExpires:    0,
		}
		sessUC.EXPECT().Login(credentials).Return(tokenDetails, nil)


		response, err := authServerGRPC.Login(context.Background(), reqValue)
		require.NoError(t, err)
		require.NotNil(t, response)
	})
}