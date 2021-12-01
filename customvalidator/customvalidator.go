package customvalidator

import (
	"fmt"
	"log"
	"playingwitherrors/models"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
)

type StructError struct {
	Struct  string      `json:"struct"`
	Field   string      `json:"field"`
	Tag     string      `json:"tag"`
	Message string      `json:"message"`
	Value   interface{} `json:"value"`
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

func FormatError(e error, myStruct string) []StructError {
	if e == nil {
		return nil
	}
	var formatter []StructError

	for _, err := range e.(validator.ValidationErrors) {
		ex := StructError{
			Struct:  myStruct,
			Field:   err.Field(),
			Tag:     err.Tag(),
			Message: translateStruct(myStruct, err.Tag(), err.Field()),
			Value:   err.Value(),
		}
		formatter = append(formatter, ex)
	}

	return formatter
}

func translateStruct(myStruct string, myTag string, myField string) string {
	switch myStruct {
	case "models.User":
		return models.TranslateUserError(myField, myTag)
	default:
		return translateTagDefault(myTag)
	}
}

//Translate to verbose the default tags
func translateTagDefault(myTag string) string {
	switch myTag {
	case "lte":
		return "Less than or equal."
	case "gte":
		return "Greater than or equal."
	case "email":
		return "E-mail not valid."
	case "cpf":
		return "CPF not valid.."
	case "required":
		return "Fields required."
	case "len":
		return "Length not valid"
	default:
		return fmt.Sprintf("Tag <%s> not identified.", myTag)
	}
}
