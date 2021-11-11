package main

import (
	"playingwitherrors/errorformatter"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestFormatErrorUser1(t *testing.T) {
	u := User{
		Name:     "!@#$!@#@!#@!$@!#",
		Email:    "email@email.com",
		Age:      20,
		Birthday: "2001-01-01",
	}
	err := validator.New().Struct(u)
	res := errorformatter.FormatError(err)
	if res == nil {
		t.Fatal("Name fild accept only alphabetical characters")
	}
}
