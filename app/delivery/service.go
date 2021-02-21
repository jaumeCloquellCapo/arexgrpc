package delivery

import (
	"context"
	errorNotFound "github.com/jaumeCloquellCapo/authGrpc/app/error"
	"github.com/jaumeCloquellCapo/authGrpc/app/model"
	"github.com/jaumeCloquellCapo/authGrpc/app/service"
	"github.com/jaumeCloquellCapo/authGrpc/grpc"
	"google.golang.org/grpc/status"
)

type microservice struct {
	authService service.AuthServiceInterface
	uService    service.UserServiceInterface
}

// Auth service constructor
func NewAuthServerGRPC(authService service.AuthServiceInterface, uService service.UserServiceInterface) *microservice {
	return &microservice{
		authService,
		uService,
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
		return nil, status.Errorf(errorNotFound.ParseGRPCErrStatusCode(err), "registerReqToUserModel: %v", err)
	}

	tokenDetail, err := u.authService.Login(*credentials)

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

// Get session id from, ctx metadata, find user by uuid and returns it
func (u *microservice) GetMe(ctx context.Context, r *grpc.GetMeRequest) (*grpc.GetMeResponse, error) {

	return nil, nil
}

// Logout user, delete current session
func (u *microservice) Logout(ctx context.Context, request *grpc.LogoutRequest) (*grpc.LogoutResponse, error) {

	return nil, nil
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
