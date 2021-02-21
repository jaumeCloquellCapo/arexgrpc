package repository

import (
	"github.com/jaumeCloquellCapo/authGrpc/internal/storage"
	"reflect"
	"testing"
)

func TestUserRepositoryInit(t *testing.T) {
	type args struct {
		db *storage.DbStore
	}
	tests := []struct {
		name string
		args args
		want UserRepositoryInterface
	}{
		{
			name: "success",
			args: args{
				db: nil,
			},
			want: &userRepository{
				db: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserInit() = %v, want %v", got, tt.want)
			}
		})
	}
}
/*
func TestUserRepository_Create(t *testing.T) {
	t.Parallel()

	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	sqlxDB := sqlx.NewDb(db, "db")

	fmt.Print(sqlxDB.DB.Ping())

	userPGRepository := NewUserRepository(&storage.DbStore{DB: sqlxDB})

	mockUser := model.CreateUser{
		Name:       "FirstName",
		LastName:   "LastName",
		Email:      "email@gmail.com",
		Country:    "es",
		Phone:      "6254551",
		PostalCode: "07440",
		Password:   "123456",
	}

	createdUser, err := userPGRepository.Create(mockUser)
	fmt.Print(err.Error())
	require.NoError(t, err)
	require.NotNil(t, createdUser)
}
*/