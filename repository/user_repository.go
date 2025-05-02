package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"oneiot-server/model/entity"
)

type IUserRepository interface {
	//C
	CreateNewUser(ctx context.Context, user entity.User) (entity.User, error)
	//R
	GetUser(ctx context.Context, email string) (entity.User, error)
	GetUserByID(ctx context.Context, userID int) (entity.User, error)
	//U
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	//D
	DeleteUser(ctx context.Context, user entity.User) error

	CheckUserExist(ctx context.Context, email string) (bool, error)
	//ToDo: After this we might need the logic to add transaction or order to the database, but lemme finish this first
}

type UserRepository struct {
	db *sql.DB
}

func (u *UserRepository) CheckUserExist(ctx context.Context, email string) (bool, error) {
	var user entity.User

	query := "SELECT Email FROM users WHERE email = ?"

	err := u.db.QueryRowContext(ctx, query, email).Scan(&user.Email)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return false, errors.New("Tidak ditemukan user dengan email" + email)
	}

	return true, nil
}

// UpdateUser updating the user_pictures, returning the new updated user_pictures data
func (u *UserRepository) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	query := "UPDATE Users SET Fullname = ?, Email = ?, Password = ?, PhoneNumber = ?, Picture = ?, Address = ?, Location = ? WHERE Id = ?"

	execContext, err := u.db.ExecContext(ctx, query,
		user.FullName, user.Email, user.Password, user.PhoneNumber,
		user.Picture, user.Address, user.Location, user.Id)

	if err != nil {
		return entity.User{}, fmt.Errorf("error update user dengan ID %d: %w", user.Id, err)
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

	queryRowContext := u.db.QueryRowContext(ctx, query, email)

	var user entity.User
	err := queryRowContext.Scan(
		&user.Id, &user.FullName, &user.Email, &user.Password,
		&user.PhoneNumber, &user.Picture, &user.Address, &user.Location,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, fmt.Errorf("tidak terdapat pengguna dengan email: %s", email) // Specific error
		}
		return entity.User{}, fmt.Errorf("error scanning data user dengan email %s: %w", email, err)
	}

	return user, nil
}

func (u *UserRepository) GetUserByID(ctx context.Context, userID int) (entity.User, error) {
	query := "SELECT * FROM users WHERE Id = ? LIMIT 1"
	queryRowContext := u.db.QueryRowContext(ctx, query, userID)

	var user entity.User
	err := queryRowContext.Scan(
		&user.Id, &user.FullName, &user.Email, &user.Password,
		&user.PhoneNumber, &user.Picture, &user.Address, &user.Location,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, fmt.Errorf("tidak terdapat pengguna dengan ID: %d", userID)
		}
		return entity.User{}, fmt.Errorf("error scanning data user dengan ID %d: %w", userID, err)
	}

	return user, nil
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

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
