package repository

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/jaumeCloquellCapo/authGrpc/app/model"
	"github.com/jaumeCloquellCapo/authGrpc/internal/storage"
	"github.com/stretchr/testify/require"
	"log"
	"reflect"
	"testing"
	"time"
)

func TestAuthRepositoryInit(t *testing.T) {
	type args struct {
		redis *storage.DbCache
	}
	tests := []struct {
		name string
		args args
		want AuthRepositoryInterface
	}{
		{
			name: "success",
			args: args{
				redis: nil,
			},
			want: &authRepository{
				redis: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthRepository(tt.args.redis); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserInit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_SetUserCtx(t *testing.T) {
	t.Parallel()

	mr, err := miniredis.Run()
	if err != nil {
		log.Fatal(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
	st := storage.DbCache{client}
	userRedisRepository := NewAuthRepository(&st)

	t.Run("SetUserCtx", func(t *testing.T) {
		user := model.User{
			ID:         0,
			Name:       "jaumke",
			LastName:   "",
			Password:   "",
			Email:      "",
			Country:    "",
			Phone:      "",
			PostalCode: "",
		}

		tk := model.TokenDetails{
			AccessToken:  "1",
			RefreshToken: "1",
			AccessUUID:   "1",
			RefreshUUID:  "1",
			AtExpires:    time.Now().Add(time.Minute * 15).Unix(),
			RtExpires:    time.Now().Add(time.Hour * 24 * 7).Unix(),
		}

		err := userRedisRepository.CreateAuth(user, tk)
		require.NoError(t, err)
	})
}
