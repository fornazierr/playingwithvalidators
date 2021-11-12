package errorformatter

import (
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
		lock.Lock()
		defer lock.Unlock()
		Validador = validator.New()
		Validador.RegisterValidation("cpf", cpfValidator)
	}

	return Validador
}

func cpfValidator(f validator.FieldLevel) bool {
	if f.Field().IsNil() {
		return false
	}

	return false
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
