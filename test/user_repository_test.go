package test

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"oneiot-server/database"
	"oneiot-server/model/entity"
	"oneiot-server/repository"
	"testing"
)

func base() (repository.IUserRepository, error) {
	err := godotenv.Load("../.env")

	if err != nil {
		return nil, nil
	}

	db := database.NewSqlConnection()

	return repository.NewUserRepository(db), nil
}

func TestInsertingNewUser(t *testing.T) {

	userRepository, _ := base()

	user, err := userRepository.CreateNewUser(context.Background(), entity.User{
		FullName:    "John Doe",
		Email:       "john.doe@gmail.com",
		Password:    "inipasswordaaa",
		PhoneNumber: "08961930",
		Picture:     "picture.jpg",
		Address:     "sana sini",
		Location:    "Jonggol",
	})

	if err != nil {
		return
	}

	fmt.Println(user)
}

func TestGetUser(t *testing.T) {
	uRepo, _ := base()

	user, err := uRepo.GetUser(context.Background(), "testing@gmail.com")
	if err != nil {
		return
	}

	fmt.Println(user)

}

func TestGetNoExistingUser(t *testing.T) {
	uRepo, _ := base()

	user, err := uRepo.GetUser(context.Background(), "awokawok@gmail.com")
	if err != nil {
		return
	}

	fmt.Println(user)
}

func TestUpdateUser(t *testing.T) {
	uRepo, _ := base()

	user, err := uRepo.UpdateUser(context.Background(), entity.User{
		FullName:    "Testing Go",
		Email:       "testing@gmail.com",
		Password:    "go.lang",
		PhoneNumber: "0891231",
		Picture:     "go.jpeg",
		Address:     "Jonggol Timur",
		Location:    "{}",
	})

	if err != nil {
		return
	}

	fmt.Println(user)
}

func TestRevertUpdateUser(t *testing.T) {
	uRepo, _ := base()

	user, err := uRepo.UpdateUser(context.Background(), entity.User{
		FullName:    "testing",
		Email:       "testing@gmail.com",
		Password:    "testingpassword",
		PhoneNumber: "081234",
		Picture:     "testingpic.jpg",
		Address:     "testing jawa tengah",
		Location:    "testing{}",
	})

	if err != nil {
		return
	}

	fmt.Println(user)
}
