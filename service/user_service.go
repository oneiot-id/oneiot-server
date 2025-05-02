package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"oneiot-server/helper"
	"oneiot-server/model/entity"
	"oneiot-server/repository"
	"strings"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	RegisterNewUser(context context.Context, user entity.User) (entity.User, error)
	GetUser(context context.Context, user entity.User) (entity.User, error)
	GetUserByID(ctx context.Context, userID int) (entity.User, error)
	UpdateUser(context context.Context, user entity.User) (entity.User, error)
	LoginUser(context context.Context, user entity.User) (entity.User, error)
	CheckUserExistence(context context.Context, user entity.User) (bool, error)
	//GetAllUser(context context.Context) ([]entity.User, error)
}

type UserService struct {
	db         *sql.DB
	repository *repository.UserRepository
	validator  *validator.Validate
}

func (u *UserService) CheckUserExistence(context context.Context, user entity.User) (bool, error) {
	_, err := u.repository.CheckUserExist(context, user.Email)

	if err != nil {
		return false, err
	}

	return true, nil
}

// GetUser this is used to retrieve user information
func (u *UserService) GetUser(ctx context.Context, user entity.User) (entity.User, error) {
	if user.Email == "" || user.Password == "" {
		return entity.User{}, errors.New("email and password are required")
	}

	//This retrieve the user data
	dbUser, err := u.repository.GetUser(ctx, user.Email)

	//This when no user with this email
	if err != nil {
		return entity.User{}, err
	}

	passwordIsSame := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))

	if passwordIsSame != nil {
		return entity.User{}, errors.New("password yang diberikan tidak sama")
	}

	//If all seems well then login the user by returning its data
	return dbUser, nil
}

func (u *UserService) GetUserByID(ctx context.Context, userID int) (entity.User, error) {
	dbUser, err := u.repository.GetUserByID(ctx, userID)
	if err != nil {
		return entity.User{}, err
	}
	return dbUser, nil
}

// LoginUser this is used to log the user in
func (u *UserService) LoginUser(ctx context.Context, user entity.User) (entity.User, error) {
	//ToDo: First we need to know if the user is existed
	dbUser, err := u.GetUser(ctx, user)

	//This when no user with this email
	if err != nil {
		return entity.User{}, err
	}

	return dbUser, nil
}

func (u *UserService) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	//ToDo: I think we need to login first to see if the password is right before updating the user
	//_, err := u.repository.GetUser(ctx, user.Email)

	dbUser, err := u.repository.GetUserByID(ctx, user.Id)

	//This when no user with this email or password is incorrect
	if err != nil {
		return entity.User{}, err
	}

	if user.Password == "" {
		user.Password = dbUser.Password // Keep existing password if not provided
	}
	// If the email is being changed, check if the *new* email is already taken by *another* user.
	if user.Email != "" && user.Email != dbUser.Email {
		// Email is being changed, check if the new email is already taken by *another* user.
		conflictingUser, checkErr := u.repository.GetUser(ctx, user.Email)
		if checkErr == nil && conflictingUser.Id != user.Id {
			// Found another user with the target email
			return entity.User{}, errors.New("email address is already in use by another account")
		} else if checkErr != nil && !strings.Contains(checkErr.Error(), "no user found") {
			// Handle potential DB error during email check, but ignore "not found" error
			fmt.Printf("Database error checking email availability for %s: %v\n", user.Email, checkErr)
			return entity.User{}, errors.New("internal server error checking email availability")
		}
		// If checkErr indicates "no user found", it's safe to proceed with the email change.
	} else if user.Email == "" {
		user.Email = dbUser.Email // Keep existing email if not provided
	}

	// 5. Call repository to update
	updatedUser, err := u.repository.UpdateUser(ctx, user)
	if err != nil {
		fmt.Printf("Error updating user ID %d in repository: %v\n", user.Id, err)
		return entity.User{}, errors.New("failed to update user information") // Generic error
	}

	return updatedUser, nil
}

// RegisterNewUser registering new user to the database, returning the current user if success
func (u *UserService) RegisterNewUser(ctx context.Context, user entity.User) (entity.User, error) {

	//First validate the user
	err := helper.ValidateUserRegister(user)

	if err != nil {
		return entity.User{}, err
	}

	//First we check if the email is already exist
	_, err = u.repository.GetUser(ctx, user.Email)

	//This is used because when user is not exist it will return error "user is not exist"
	if err != nil {
		//create new user if the email is not existed

		//Hash the password with bcrypt
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(hashedPassword)

		if err != nil {
			return entity.User{}, err
		}

		newUser, err := u.repository.CreateNewUser(ctx, user)

		if err != nil {
			return entity.User{}, err
		}

		return newUser, nil
	}

	return entity.User{}, errors.New("Terdapat pengguna dengan email yang sama")
}

// NewUserService creating new user service
func NewUserService(userRepository *repository.UserRepository, db *sql.DB) *UserService {

	//this will use later if we use transactional method
	//repo := repository.NewUserRepository(db)

	return &UserService{
		repository: userRepository,

		//This is not used for now because we don't use the sql.Tx method, we'll try to use this later when we need the transactional
		db:        db,
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}
}
