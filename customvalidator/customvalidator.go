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

func Instance() *validator.Validate {
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

//Realiza a validação de CNPJ
func cnpjValidator(f validator.FieldLevel) bool {
	val := f.Field().String()

	newCnpj := val[:12]

	dv1 := utilGeraDigitoCNPJ(newCnpj)
	newCnpj += dv1

	dv2 := utilGeraDigitoCNPJ(newCnpj)
	newCnpj += dv2

	if strings.Compare(newCnpj, val) == 0 {
		return true
	} else {
		return false
	}
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

//Translate to verbose
func translateTag(tag string) string {
	switch tag {
	case "lte":
		return "Less than or equal"
	case "gte":
		return "Greater than or equal"
	case "email":
		return "E-mail"
	case "cpf":
		return "CPF inválido"
	default:
		return tag
	}
}
