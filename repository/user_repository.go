package repository

import (
	"context"
	"database/sql"
	"errors"
	"oneiot-server/model/entity"
)

type IUserRepository interface {
	//C
	CreateNewUser(ctx context.Context, user entity.User) (entity.User, error)
	//R
	GetUser(ctx context.Context, email string) (entity.User, error)
	//U
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	//D
	DeleteUser(ctx context.Context, user entity.User) error

	//ToDo: After this we might need the logic to add transaction or order to the database, but lemme finish this first
}

type UserRepository struct {
	db *sql.DB
}

// UpdateUser updating the user_pictures, returning the new updated user_pictures data
func (u *UserRepository) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	query := "UPDATE Users SET Fullname = ?, Email = ?, Password = ?, PhoneNumber = ?, Picture = ?, Address = ?, Location = ? WHERE Email = ?"

	execContext, err := u.db.ExecContext(ctx, query, user.FullName, user.Email, user.Password, user.PhoneNumber, user.Picture, user.Address, user.Location, user.Email)

	if err != nil {
		return entity.User{}, errors.New("error while updating user_pictures")
	}

	_, err = execContext.RowsAffected()

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

// GetUser this will get the user_pictures by email
func (u *UserRepository) GetUser(ctx context.Context, email string) (entity.User, error) {
	query := "SELECT * FROM users WHERE email = ? LIMIT 1"

	queryContext, err := u.db.QueryContext(ctx, query, email)

	defer queryContext.Close()

	if err != nil {
		return entity.User{}, errors.New("Error saat mendapatkan user_pictures")
	}

	if !queryContext.Next() {
		return entity.User{}, errors.New("Tidak terdapat pengguna yang menggunakan email ini")
	}

	var user entity.User
	err = queryContext.Scan(
		&user.Id, &user.FullName, &user.Email, &user.Password,
		&user.PhoneNumber, &user.Picture, &user.Address, &user.Location,
	)

	if err != nil {
		return entity.User{}, errors.New("Error saat scanning pengguna")
	}

	return user, nil
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// DeleteUser deleting the user_pictures by querying the id
func (u *UserRepository) DeleteUser(ctx context.Context, user entity.User) error {
	query := "DELETE FROM Users WHERE id=?"

	_, err := u.db.ExecContext(ctx, query, user.Id)

	if err != nil {
		return err
	}
	return nil
}

// CreateNewUser creating new user_pictures based on the requested data
func (u *UserRepository) CreateNewUser(ctx context.Context, user entity.User) (entity.User, error) {
	query := "INSERT INTO Users(Fullname, Email, Password, PhoneNumber, Picture, Address, Location) VALUES(?, ?, ?, ?, ?, ?, ?);"

	execContext, err := u.db.ExecContext(ctx, query, user.FullName, user.Email, user.Password, user.PhoneNumber, user.Picture, user.Address, user.Location)

	if err != nil {
		return entity.User{}, errors.New("error while creating new user_pictures")
	}

	id, err := execContext.LastInsertId()
	if err != nil {
		return entity.User{}, errors.New("error retrieving last inserted ID")
	}

	user.Id = int(id)

	return user, nil
}
