package main

import (
	"playingwitherrors/customvalidator"
	"testing"
)

func TestFormatErrorUser1(t *testing.T) {
	u := User{
		Name:     "!@#$!@#@!#@!$@!#",
		Email:    "email@email.com",
		Age:      20,
		Birthday: "2001-01-01",
		CPF:      "111444777322",
	}
	err := customvalidator.InitValidador().Struct(u)
	if err == nil {
		t.Fatalf("Some fields with error: %s", err.Error())
	}
}
