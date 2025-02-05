package test

import (
	"github.com/go-playground/validator/v10"
	"testing"
)

func TestJsonMarshaller(t *testing.T) {
	//user_pictures := entity.User{
	//	FullName: "vincent kenutama",
	//	Email:    "vincent@gmail.com",
	//}
	//
	//output, err := helper.MarshalThis(&user_pictures)
	//
	//assert.NoError(t, err)
	//fmt.Println(output)
}

type ValidationStructTest struct {
	FullName string `validate:"required"`
}

func TestValidator(t *testing.T) {
	test := ValidationStructTest{
		FullName: "",
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(&test)

	if err != nil {
		return
	}

}
