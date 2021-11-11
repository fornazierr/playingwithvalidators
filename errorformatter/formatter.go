package errorformatter

import "github.com/go-playground/validator/v10"

type StructError struct {
	Field string
	Tag   string
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
