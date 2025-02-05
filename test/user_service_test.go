package test

import (
	"context"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"oneiot-server/database"
	"oneiot-server/model/entity"
	repository2 "oneiot-server/repository"
	"oneiot-server/service"
	"testing"
)

func userServiceTestBase() *service.UserService {
	err := godotenv.Load("../.env")

	if err != nil {
		return nil
	}

	db := database.NewSqlConnection()

	repository := repository2.NewUserRepository(db)

	return service.NewUserService(repository, db)
}

// This test to add new user_pictures to the database
func TestRegisterNewUser(t *testing.T) {
	userService := userServiceTestBase()

	user, err := userService.RegisterNewUser(context.Background(), entity.User{
		FullName:    "Vincent Kenutama",
		Email:       "vincent@gmail.com",
		Password:    "password",
		PhoneNumber: "082131313",
		Picture:     "vincent.jpg",
		Address:     "Yogyakarta",
		Location:    "{}",
	})
	if err != nil {
		return
	}

	fmt.Println(user)
}

func TestRegisterExistedUser(t *testing.T) {
	userService := userServiceTestBase()

	user, err := userService.RegisterNewUser(context.Background(), entity.User{
		FullName:    "Vincent Kenutama",
		Email:       "vincent@gmail.com",
		Password:    "password",
		PhoneNumber: "082131313",
		Picture:     "vincent.jpg",
		Address:     "Yogyakarta",
		Location:    "{}",
	})

	if err != nil {
		return
	}

	fmt.Println(user)
}

func TestErrorToString(t *testing.T) {
	e := errors.New("test error")

	fmt.Println(e)
}

func TestLogin(t *testing.T) {
	userService := userServiceTestBase()

	user, err := userService.LoginUser(context.Background(), entity.User{
		FullName:    "Vincent Kenutama",
		Email:       "vincent@gmail.com",
		Password:    "password",
		PhoneNumber: "082131313",
		Picture:     "vincent.jpg",
		Address:     "Yogyakarta",
		Location:    "{}",
	})

	assert.NoError(t, err)
	fmt.Println(user)
}

func TestLoginNotExistUser(t *testing.T) {
	userService := userServiceTestBase()

	user, err := userService.LoginUser(context.Background(), entity.User{
		Email:    "mock@gmail.com",
		Password: "password",
	})

	assert.NotNil(t, err)
	fmt.Println(user, err)
}

func TestLoginWrongPassword(t *testing.T) {
	userService := userServiceTestBase()

	user, err := userService.LoginUser(context.Background(), entity.User{
		Email:    "vincent@gmail.com",
		Password: "awokawok",
	})

	assert.NotNil(t, err)
	fmt.Println(user, err)
}

func TestGetUserService(t *testing.T) {
	userService := userServiceTestBase()

	user, err := userService.GetUser(context.Background(), entity.User{
		Email: "vincent@gmail.com",
	})

	assert.Nil(t, err)
	fmt.Println(user, err)
}

func TestGetNotExistedUser(t *testing.T) {
	userService := userServiceTestBase()

	user, err := userService.GetUser(context.Background(), entity.User{
		Email: "mock@gmail.com",
	})

	assert.Error(t, err)
	fmt.Println(user, err)
}

func TestUpdateUserService(t *testing.T) {
	userService := userServiceTestBase()

	user, err := userService.UpdateUser(context.Background(), entity.User{
		FullName:    "Vincent Kenutama Update Test",
		Email:       "vincent@gmail.com",
		Password:    "password",
		PhoneNumber: "082131313",
		Picture:     "vincent.jpg",
		Address:     "Yogyakarta",
		Location:    "{}",
	})

	assert.Nil(t, err)
	fmt.Println(user, err)
}

func TestRevertUpdateUserService(t *testing.T) {
	userService := userServiceTestBase()

	user, err := userService.UpdateUser(context.Background(), entity.User{
		FullName:    "Vincent Kenutama",
		Email:       "vincent@gmail.com",
		Password:    "password",
		PhoneNumber: "082131313",
		Picture:     "vincent.jpg",
		Address:     "Yogyakarta",
		Location:    "{}",
	})

	assert.Nil(t, err)
	fmt.Println(user, err)
}

func TestWrongPasswordAtUpdatingUser(t *testing.T) {
	userService := userServiceTestBase()

	user, err := userService.UpdateUser(context.Background(), entity.User{
		FullName:    "Vincent Kenutama",
		Email:       "vincent@gmail.com",
		Password:    "wrongpassword",
		PhoneNumber: "082131313",
		Picture:     "vincent.jpg",
		Address:     "Yogyakarta",
		Location:    "{}",
	})

	assert.NotNil(t, err)
	fmt.Println(user, err)
}
