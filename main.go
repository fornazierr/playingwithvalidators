package main

import (
	"fmt"
	"log"
	"playingwitherrors/customvalidator"
	"playingwitherrors/models"
)

func main() {
	log.Println("Starting playingwithvalidators")
	validate := customvalidator.Instance()

	//////
	////// Checking Struct
	//////
	log.Println("Check User struct")
	withError := models.User{
		Name:     "david",
		Email:    "david@email,com",
		Age:      103,
		Birthday: "1999-11-02",
		CPF:      "111444777322",
	}
	err := validate.Struct(withError)
	if err != nil {
		formated := customvalidator.FormatError(err)
		fmt.Printf("Found:\n%+v\n", formated)
	}

	//////
	////// Checking variables
	//////

	log.Println("Check CPF variable")
	cpf := "12345899445"
	err = customvalidator.Instance().Var(cpf, "required,min=11,max=11,cpf")
	if err != nil {
		formated := customvalidator.FormatError(err)
		log.Printf("CPF [%s] error: %+v\n", cpf, formated)
	} else {
		log.Printf("CPF [%s] without error.\n", cpf)
	}

	log.Println("Check CNPJ variable")
	cnpj := "11222333000181"
	err = customvalidator.Instance().Var(cnpj, "required,min=14,max=14,cnpj")
	if err != nil {
		formated := customvalidator.FormatError(err)
		log.Printf("CNPJ [%s] error: %+v\n", cnpj, formated)
	} else {
		log.Printf("CNPJ [%s] without error.\n", cpf)
	}
}
