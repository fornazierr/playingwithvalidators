package models

type User struct {
	Name         string `json:"name" validate:"required,alpha"`
	Email        string `json:"email" validate:"required,email"`
	Age          int    `json:"age" validate:"required,gte=0,lte=100"` //could be min and max too
	Birthday     string `json:"birthday" validate:"required,datetime=2006-01-02"`
	CPF          string `json:"cpf" validate:"required,len=11,cpf"`
	InternalCode string `json:"internalCode" validate:"required"`
}

func TranslateUserError(myField string, myTag string) string {
	switch myField {
	case "Name":
		switch myTag {
		case "required":
			return "<Name> field required."
		case "alpha":
			return "<Name> field accept only alphanumeric."
		default:
			return myTag
		}
	case "Email":
		switch myTag {
		case "required":
			return "<Email> field required."
		case "email":
			return "<Emai> filed is not valid."
		default:
			return myTag
		}
	case "Age":
		switch myTag {
		case "required":
			return "<Age> field required."
		case "gte":
			return "<Age> field must be greater os equal to 0 (zero)."
		case "lte":
			return "<Age> field must be less or equal to 100 (one hundred)"
		default:
			return myTag
		}
	case "Birthday":
		switch myTag {
		case "required":
			return "<Birthday> field required."
		case "datetime":
			return "<Birthday> field date format is not valid"
		default:
			return myTag
		}
	case "CPF":
		switch myTag {
		case "required":
			return "<CPF> field required."
		case "len":
			return "<CPF> field lenght must be equal to 11 (eleven)."
		case "cpf":
			return "<CPF> field is not valid."
		default:
			return myTag
		}
	case "InternalCode":
		switch myTag {
		case "required":
			return "<InternalCode> field required."
		default:
			return myTag
		}
	default:
		return myTag
	}
}
