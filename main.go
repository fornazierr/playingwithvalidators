package main

import (
	"fmt"
	"playingwitherrors/errorformatter"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Age      int    `json:"age" validate:"required,gte=0,lte=100"`
	Birthday string `json:"birthday" validate:"required,datetime=2006-01-02"`
}

func main() {
	validate := validator.New()

	withError := User{
		Name:     "david",
		Email:    "david@email,com",
		Age:      103,
		Birthday: "1999-11-02",
	}
	err := validate.Struct(withError)
	if err != nil {
		formated := errorformatter.FormatError(err)
		fmt.Printf("Found:\n%+v\n", formated)
	}

}
