package main

import "fmt"

func main() {
	// crira um programa que completa os parenteses exemplo
	/*
		entrada =  )))()((     ((()))()(())
		entrada = )(
	*/

	fmt.Println(completarParenteses("))()(("))
	fmt.Println(completarParenteses(")("))
	fmt.Println(completarParenteses("(())((()))(("))

}

func tranformaEntradaEmArray(entrada string) []string {
	var array []string
	for i := 0; i < len(entrada); i++ {
		array = append(array, string(entrada[i]))
	}
	return array
}

func criaCopiaArray(array []string) []string {
	arrayCopia := make([]string, len(array))
	copy(arrayCopia, array)
	return arrayCopia
}

func completarParenteses(entrada string) string {
	var arrayLetraOriginal []string
	arrayLetraOriginal = tranformaEntradaEmArray(entrada)

	var arrayLetraCopiaTerceira []string
	arrayLetraCopiaTerceira = criaCopiaArray(arrayLetraOriginal)

	for i, letra := range arrayLetraOriginal {
		if letra == "(" {
			tem := false
			for _, letra2 := range arrayLetraOriginal[i:] {
				if letra2 == ")" {
					tem = true
					break
				}
			}
			if !tem {
				arrayLetraCopiaTerceira = append(arrayLetraCopiaTerceira, ")")
			}
		} else if letra == ")" {
			tem := false
			for i := len(arrayLetraOriginal[:i]) - 1; i >= 0; i-- {
				letra2 := arrayLetraOriginal[i]
				if letra2 == "(" {
					tem = true
					break
				}
			}
			if !tem {
				arrayLetraCopiaTerceira = append([]string{"("}, arrayLetraCopiaTerceira...)
			}
		}
	}

	var newString string
	for _, letra := range arrayLetraCopiaTerceira {
		newString += letra
	}
	return newString
}
