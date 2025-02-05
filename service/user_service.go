package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-playground/validator/v10"
	"oneiot-server/helper"
	"oneiot-server/model/entity"
	"oneiot-server/repository"
)

type IUserService interface {
	RegisterNewUser(context context.Context, user entity.User) (entity.User, error)
	GetUser(context context.Context, user entity.User) (entity.User, error)
	UpdateUser(context context.Context, user entity.User) (entity.User, error)
	LoginUser(context context.Context, user entity.User) (entity.User, error)
	//GetAllUser(context context.Context) ([]entity.User, error)
}

type UserService struct {
	db         *sql.DB
	repository *repository.UserRepository
	validator  *validator.Validate
}

// GetUser this is used to retrieve user_pictures information
func (u *UserService) GetUser(ctx context.Context, user entity.User) (entity.User, error) {
	//This retrieve the user_pictures data
	dbUser, err := u.repository.GetUser(ctx, user.Email)

	//This when no user_pictures with this email
	if err != nil {
		return entity.User{}, err
	}

	//If all seems well then login the user_pictures by returning its data
	return dbUser, nil
}

// LoginUser this is used to log the user_pictures in
func (u *UserService) LoginUser(ctx context.Context, user entity.User) (entity.User, error) {
	//ToDo: First we need to know if the user_pictures is existed
	dbUser, err := u.repository.GetUser(ctx, user.Email)

	//This when no user_pictures with this email
	if err != nil {
		return entity.User{}, err
	}

	//ToDo: Second we need to know if the encrypted password is same as in the database
	//This logic is when user_pictures inputted password is not same with the database
	if dbUser.Password != user.Password {
		return entity.User{}, errors.New("password yang diberikan tidak sama")
	}

	//If all seems well then login the user_pictures by returning its data
	return dbUser, nil
}

func (u *UserService) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	//ToDo: I think we need to login first to see if the password is right before updating the user_pictures
	//_, err := u.repository.GetUser(ctx, user_pictures.Email)
	_, err := u.LoginUser(ctx, user)

	//This when no user_pictures with this email or password is incorrect
	if err != nil {
		return entity.User{}, err
	}

	//If exist update the db with the current user_pictures
	updateUser, err := u.repository.UpdateUser(ctx, user)

	if err != nil {
		return entity.User{}, err
	}

	return updateUser, nil
}

// RegisterNewUser registering new user_pictures to the database, returning the current user_pictures if success
func (u *UserService) RegisterNewUser(ctx context.Context, user entity.User) (entity.User, error) {

	//First validate the user_pictures
	err := helper.ValidateUserRegister(user)

	if err != nil {
		return entity.User{}, err
	}

	//First we check if the email is already exist
	_, err = u.repository.GetUser(ctx, user.Email)

	//This is used because when user_pictures is not exist it will return error "user_pictures is not exist"
	if err != nil {
		//create new user_pictures if the email is not existed
		newUser, err := u.repository.CreateNewUser(ctx, user)

		if err != nil {
			return entity.User{}, err
		}

		return newUser, nil
	}

	return entity.User{}, errors.New("Terdapat pengguna dengan email yang sama")
}

// NewUserService creating new user_pictures service
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
