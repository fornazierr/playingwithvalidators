package models

type User struct {
	Name         string `json:"name" validate:"required,alpha"`
	Email        string `json:"email" validate:"required,email"`
	Age          int    `json:"age" validate:"required,gte=0,lte=100"` //could be min and max too
	Birthday     string `json:"birthday" validate:"required,datetime=2006-01-02"`
	CPF          string `json:"cpf" validate:"required,min=11,max=11,cpf"`
	InternalCode string `json:"internalCode" validate:"required"`
}
