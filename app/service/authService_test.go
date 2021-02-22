package service

import (
	"github.com/golang/mock/gomock"
	"github.com/jaumeCloquellCapo/authGrpc/app/model"
	"github.com/jaumeCloquellCapo/authGrpc/app/repository"
	"github.com/jaumeCloquellCapo/authGrpc/internal/helpers"
	"github.com/jaumeCloquellCapo/authGrpc/mock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestAuthRepositoryInit(t *testing.T) {
	type args struct {
		authRepository repository.AuthRepositoryInterface
		userRepository repository.UserRepositoryInterface
	}
	tests := []struct {
		name string
		args args
		want AuthServiceInterface
	}{
		{
			name: "success",
			args: args{
				authRepository: nil,
				userRepository: nil,
			},
			want: &authService{
				authRepository: nil,
				userRepository: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthService(tt.args.authRepository, tt.args.userRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserInit() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
func TestAuthService_SignUp(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userR := mock.NewMockUserPGRepository(ctrl)
	authR := mock.NewMockAuthRepository(ctrl)
	authServerGRPC := NewAuthService(authR, userR)


	var pass = "123"
	bytePassword := []byte(pass)

	password, _ := helpers.HashAndSalt(bytePassword)

	reqValue := model.CreateUser{
		Name:       "q",
		LastName:   "q",
		Password:   password,
		Email:      "d123@ll.com",
		Country:    "d",
		Phone:      "d",
		PostalCode: "d",
	}


	t.Run("SignUpd", func(t *testing.T) {
		t.Parallel()
		//		userID := uint64(5)
		user := &model.CreateUser{
			Name:       reqValue.Name,
			LastName:   reqValue.LastName,
			Email:      reqValue.Email,
			Country:    reqValue.Country,
			Phone:       reqValue.Phone,
			PostalCode:  reqValue.PostalCode,
			Password:   reqValue.Password,
		}

		var err error

		var userResponse *model.User

		var token model.TokenDetails

		userR.EXPECT().Create(user).Return(userResponse, err)

		authR.EXPECT().CreateToken(userResponse).Return(token, err)

		authR.EXPECT().CreateAuth(userResponse, token)

		model, token, err := authServerGRPC.SignUp(reqValue)
		require.NoError(t, err)
		require.NotNil(t, model)
	})
}
*/

func TestAuthService_Login(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userR := mock.NewMockUserPGRepository(ctrl)
	authR := mock.NewMockAuthRepository(ctrl)
	authServerGRPC := NewAuthService(authR, userR)

	reqValue := model.Credentials{
		Email:    "email@gmail.com",
		Password: "kjshfkjdshfklhfkjhfkjwhfkwfhkwjhfkjwhfkjwehfkjwhfkjwfhkjwhfkjhhglkejgoej",
	}

	bytePassword := []byte(reqValue.Password)

	hashPassword, _ := helpers.HashAndSalt(bytePassword)

	t.Run("Login", func(t *testing.T) {
		t.Parallel()
		userID := int64(1)
		user := model.User{
			ID:         userID,
			Name:       "a",
			LastName:   "a",
			Password:   hashPassword,
			Email:      reqValue.Email,
			Country:    "a",
			Phone:      "a",
			PostalCode: "a",
		}

		userRes := &model.User{
			ID:         userID,
			Name:       user.Name,
			LastName:   user.LastName,
			Password:   user.Password,
			Email:      user.Email,
			Country:    user.Country,
			Phone:      user.Phone,
			PostalCode: user.PostalCode,
		}
		token := model.TokenDetails{
			AccessToken:  "123",
			RefreshToken: "123",
			AccessUUID:   "123",
			RefreshUUID:  "123",
			AtExpires:    0,
			RtExpires:    0,
		}
		var err error

		//userR.EXPECT().Create(user).Return(userRes, err)

		userR.EXPECT().FindByEmail(user.Email).Return(userRes, err)

		authR.EXPECT().CreateToken(user).Return(token, err)

		authR.EXPECT().CreateAuth(user, token)

		response, err := authServerGRPC.Login(reqValue)

		require.NoError(t, err)
		require.NotNil(t, response)

		require.Equal(t, token, response)
	})
}
