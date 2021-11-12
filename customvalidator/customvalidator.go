package customvalidator

import (
	"log"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
)

type StructError struct {
	Field string
	Tag   string
}

var lock = &sync.Mutex{}
var Validador *validator.Validate

func InitValidador() *validator.Validate {
	if Validador == nil {
		log.Println("Validador nulo, iniciando")
		lock.Lock()
		defer lock.Unlock()
		Validador = validator.New()
		Validador.RegisterValidation("cpf", cpfValidator)
		Validador.RegisterValidation("cnpj", cnpjValidator)
	}

	return Validador
}

//Realiza a validação de CPF
func cpfValidator(f validator.FieldLevel) bool {
	val := f.Field().String()
	newCpf := val[:9]

	dv1 := calculaDigitoCPF(newCpf)
	newCpf += dv1

	dv2 := calculaDigitoCPF(newCpf)
	newCpf += dv2

	if strings.Compare(newCpf, val) == 0 {
		return true
	} else {
		return false
	}
}

func cnpjValidator(f validator.FieldLevel) bool {
	return true
}

func FormatError(e error) []StructError {
	if e == nil {
		return nil
	}
	var formatter []StructError

	for _, err := range e.(validator.ValidationErrors) {
		ex := StructError{
			Field: err.Field(),
			Tag:   translateTag(err.Tag()),
		}
		formatter = append(formatter, ex)
	}

	return formatter
}

func translateTag(tag string) string {
	switch tag {
	case "lte":
		return "Less than"
	case "gte":
		return "Greater than"
	case "email":
		return "E-mail"
	default:
		return tag
	}
}
