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
