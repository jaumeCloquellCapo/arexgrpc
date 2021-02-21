package delivery

import (
	"context"
	errorNotFound "github.com/jaumeCloquellCapo/authGrpc/app/error"
	"github.com/jaumeCloquellCapo/authGrpc/app/model"
	"github.com/jaumeCloquellCapo/authGrpc/app/service"
	"github.com/jaumeCloquellCapo/authGrpc/grpc"
	"github.com/jaumeCloquellCapo/authGrpc/internal/logger"
	"github.com/jaumeCloquellCapo/authGrpc/internal/middleware"
	"google.golang.org/grpc/status"
)

type microservice struct {
	authService service.AuthServiceInterface
	uService    service.UserServiceInterface
	logger      logger.Logger
}

// Auth service constructor
func NewUserServerGRPC(authService service.AuthServiceInterface, uService service.UserServiceInterface, logger logger.Logger) *microservice {
	return &microservice{
		authService,
		uService,
		logger,
	}
}

func (u *microservice) Register(c context.Context, r *grpc.RegisterRequest) (*grpc.TokenDetails, error) {

	user, err := u.registerReqToUserModel(r)
	if err != nil {
		return nil, status.Errorf(errorNotFound.ParseGRPCErrStatusCode(err), "registerReqToUserModel: %v", err)
	}

	_, tokenDetail, err := u.authService.SignUp(*user)
	if err != nil {
		return nil, status.Errorf(errorNotFound.ParseGRPCErrStatusCode(err), "registerReqToUserModel: %v", err)
	}

	return &grpc.TokenDetails{
		AccessToke:   tokenDetail.AccessToken,
		RefreshToken: tokenDetail.RefreshToken,
		AccessUUID:   tokenDetail.AccessUUID,
		RefreshUUID:  tokenDetail.RefreshUUID,
		AtExpires:    tokenDetail.AtExpires,
		RtExpires:    tokenDetail.RtExpires,
	}, nil
}

// Login user with email and password
func (u *microservice) Login(ctx context.Context, r *grpc.LoginRequest) (*grpc.TokenDetails, error) {

	credentials, err := u.registerReqToCredentialsModel(r)
	if err != nil {
		u.logger.Error(err)
		return nil, status.Errorf(errorNotFound.ParseGRPCErrStatusCode(err), "registerReqToUserModel: %v", err)
	}

	tokenDetail, err := u.authService.Login(*credentials)

	if err != nil {
		u.logger.Error(err)
		return nil, status.Errorf(errorNotFound.ParseGRPCErrStatusCode(err), "registerReqToUserModel: %v", err)
	}

	return &grpc.TokenDetails{
		AccessToke:   tokenDetail.AccessToken,
		RefreshToken: tokenDetail.RefreshToken,
		AccessUUID:   tokenDetail.AccessUUID,
		RefreshUUID:  tokenDetail.RefreshUUID,
		AtExpires:    tokenDetail.AtExpires,
		RtExpires:    tokenDetail.RtExpires,
	}, nil
}

// Logout user, delete current session
func (u *microservice) Logout(ctx context.Context, request *grpc.LogoutRequest) (*grpc.LogoutResponse, error) {

	tokenAuth, err := middleware.GetAccessTokenFromContext(ctx)

	if err != nil {
		return nil, status.Errorf(errorNotFound.ParseGRPCErrStatusCode(err), "Logout: %v", err)
	}
	err = u.authService.Logout(tokenAuth.AccessUUID)

	if err != nil {
		return nil, status.Errorf(errorNotFound.ParseGRPCErrStatusCode(err), "Logout: %v", err)
	}

	return &grpc.LogoutResponse{}, nil
}

func (u *microservice) registerReqToUserModel(r *grpc.RegisterRequest) (*model.CreateUser, error) {

	candidate := &model.CreateUser{
		Email:    r.GetEmail(),
		Name:     r.GetFirstName(),
		LastName: r.GetLastName(),
		Password: r.GetPassword(),
	}

	return candidate, nil
}

func (u *microservice) registerReqToCredentialsModel(r *grpc.LoginRequest) (*model.Credentials, error) {

	candidate := &model.Credentials{
		Email:    r.GetEmail(),
		Password: r.GetPassword(),
	}

	return candidate, nil
}

func (u *microservice) userModelToProto(user *model.User) *grpc.User {
	userProto := &grpc.User{
		Id:        user.ID,
		FirstName: user.Name,
		LastName:  user.LastName,
		Password:  user.Password,
		Email:     user.Email,
	}
	return userProto
}
