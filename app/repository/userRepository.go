package repository

import (
	"database/sql"
	error2 "github.com/jaumeCloquellCapo/authGrpc/app/error"
	"github.com/jaumeCloquellCapo/authGrpc/app/model"
	"github.com/jaumeCloquellCapo/authGrpc/internal/storage"
)

type userRepository struct {
	db *storage.DbStore
}

//UserRepositoryInterface ...
type UserRepositoryInterface interface {
	//FindAll() ([]model.User, error)
	FindById(id int) (user *model.User, err error)
	RemoveById(id int) error
	UpdateById(id int, user model.UpdateUser) error
	FindByEmail(email string) (user *model.User, err error)
	Create(model.CreateUser) (user *model.User, err error)
}

//NewUserRepository ...
func NewUserRepository(db *storage.DbStore) UserRepositoryInterface {
	return &userRepository{
		db,
	}
}

//FindById ...
func (r *userRepository) FindById(id int) (user *model.User, err error) {
	user = &model.User{}

	var query = "SELECT id, email, name, postal_code, phone, last_name, country FROM users WHERE id = $1"
	row := r.db.QueryRow(query, id)

	if err := row.Scan(&user.ID, &user.Email, &user.Name, &user.PostalCode, &user.Phone, &user.LastName, &user.Country); err != nil {
		if err == sql.ErrNoRows {
			return nil, error2.ErrNotFound
		}

		return nil, err
	}

	return user, nil
}

func (r *userRepository) RemoveById(id int) error {

	_, err := r.db.Exec(`DELETE FROM users WHERE id = $1;`, id)
	return err
}

//UpdateById ...
func (r *userRepository) UpdateById(id int, user model.UpdateUser) error {
	result, err := r.db.Exec("UPDATE users SET name = $1, email = $2, last_name = $3, phone = $4, postal_code = $5, country = $6 where id = $7", user.Name, user.Email, user.LastName, user.Phone, user.PostalCode, user.Country, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rows != 1 {
		return error2.ErrNotFound
	}

	return nil
}

//FindByEmail
func (r *userRepository) FindByEmail(email string) (user *model.User, err error) {

	user = &model.User{}

	var query = "SELECT id, email, name, password FROM users WHERE email = $1"
	row := r.db.QueryRow(query, email)

	if err := row.Scan(&user.ID, &user.Email, &user.Name, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, error2.ErrNotFound
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Create(UserSignUp model.CreateUser) (user *model.User, err error) {
	//query := "INSERT INTO users (name, password, email, last_name, phone, postal_code, country) values  ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	createUserQuery := `INSERT INTO users (name, password, email, last_name, phone, postal_code, country) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`


	stmt, err := r.db.Prepare(createUserQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var userId int64
	err = stmt.QueryRow(UserSignUp.Name, UserSignUp.Password, UserSignUp.Email, UserSignUp.LastName, UserSignUp.Phone, UserSignUp.PostalCode, UserSignUp.Country).Scan(&userId)
	if err != nil {

		return nil, err
	}

	return &model.User{
		ID:         userId,
		Name:       UserSignUp.Name,
		Email:      UserSignUp.Email,
		LastName:   UserSignUp.LastName,
		Phone:      UserSignUp.Phone,
		Country:    UserSignUp.Country,
		PostalCode: UserSignUp.PostalCode,
	}, nil
}
