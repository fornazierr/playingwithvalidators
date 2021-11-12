package customvalidator

import (
	"fmt"
	"strconv"
)

func calculaDigitoCPF(cpf string) string {
	var mod int
	if len(cpf) == 9 {
		mod = 10
	} else {
		mod = 11
	}

	soma := 0
	for i := 0; i < len(cpf); i++ {
		digito, _ := strconv.Atoi(string(cpf[i]))
		soma += digito * mod
		mod--
	}

	resto := soma % 11
	if resto < 2 {
		return "0"
	} else {
		return fmt.Sprint(11 - resto)
	}
}

func utilGeraDigitoCNPJ(cnpj string) string {
	var mod int
	if len(cnpj) == 12 {
		mod = 5
	} else {
		mod = 6
	}

	soma := 0
	for i := 0; i < len(cnpj); i++ {
		digito, _ := strconv.Atoi(string(cnpj[i]))
		soma += digito * mod
		mod--
		if mod < 2 {
			mod = 9
		}
	}

	resto := soma % 11
	if resto < 2 {
		return "0"
	} else {
		return fmt.Sprint(11 - resto)
	}
}
