package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jaumeCloquellCapo/authGrpc/app/model"
	"github.com/jaumeCloquellCapo/authGrpc/internal/storage"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserRepository_FindByEmail(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "db")
	defer sqlxDB.Close()

	userPGRepository := NewUserRepository(&storage.DbStore{DB:sqlxDB})

	columns := []string{"id", "email", "name", "password"}
	userUUID := int64(1)
	mockUser := &model.User{
		ID:         userUUID,
		Name:       "FirstName",
		LastName:   "LastName",
		Password:   "123456",
		Email:      "email@gmail.com",
		Country:    "es",
		Phone:      "es",
		PostalCode: "es",
	}

	rows := sqlmock.NewRows(columns).AddRow(
		userUUID,
		mockUser.Email,
		mockUser.Name,
		mockUser.Password,
	)

	mock.ExpectQuery("SELECT id, email, name, password FROM users WHERE email = ?").WithArgs(mockUser.Email).WillReturnRows(rows)

	foundUser, err := userPGRepository.FindByEmail(mockUser.Email)
	require.NoError(t, err)
	require.NotNil(t, foundUser)
	require.Equal(t, foundUser.Email, mockUser.Email)
}