package validations

import (
	"fmt"
	"strconv"
	"strings"
)

// func main()  {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	operacao := scanner.Text()
// 	Equacao(operacao)
// }

func Equacao(operacao string) int {
	resultado := 0

	switch {
	case strings.Contains(operacao, "+"):
		resultado = Somar(operacao)
	case strings.Contains(operacao, "-"):
		resultado = Subtrair(operacao)
	case strings.Contains(operacao, "*"):
		resultado = Multiplicar(operacao)
	case strings.Contains(operacao, "/"):
		resultado = Dividir(operacao)
	default:
		fmt.Println("Error: Digite um numero válido!!")
	}
	return resultado
}

func Somar(operadores string) int {

	valores := strings.Split(operadores, "+")
	for j := 0; j < len(valores); j++ {
		if strings.Contains(valores[j], "*") || strings.Contains(valores[j], "/") || strings.Contains(valores[j], "-") {
			fmt.Println(j,valores[j])
			valores[j] = strconv.Itoa(Equacao(valores[j]))
			break
		}
	}
	resultado := 0

	for i := 0; i < len(valores); i++ {
		num, err := strconv.Atoi(valores[i])
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error: Digite um numero válido!!")
		}else{
			resultado += num
		}
	}
	return resultado
}

func Subtrair(operadores string) int {

	valores := strings.Split(operadores, "-")
	for j := 0; j < len(valores); j++ {
		if strings.Contains(valores[j], "*") || strings.Contains(valores[j], "/") || strings.Contains(valores[j], "+") {
			fmt.Println(j,valores[j])
			valores[j] = strconv.Itoa(Equacao(valores[j]))
			break
		}
	}
	resultado := 0

	for i := 0; i < len(valores); i++ {
		num, err := strconv.Atoi(valores[i])
		if  i == 0{
			resultado = num
		}else if err != nil {
			fmt.Println(err)
		}else{
			resultado = resultado - num
		}
	}
	return resultado
}

func Multiplicar(operadores string) int {

	valores := strings.Split(operadores, "*")
	resultado := 1

	for i := 0; i < len(valores); i++ {
		num, err := strconv.Atoi(valores[i])
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error: Digite um numero válido!!")
		}else{
			resultado *= num
		}
	}
	return resultado
}

func Dividir(operadores string) int {

	valores := strings.Split(operadores, "/")
	resultado := 1

	for i := 0; i < len(valores); i++ {
		num, err := strconv.Atoi(valores[i])
		if  i == 0{
			resultado = num
		}else{
			if err != nil {
				fmt.Println(err)
				fmt.Println(num)
				fmt.Println("Error: Digite um numero válido!!")
			}else{
				resultado /= num
			}
		}
	}

	return resultado
}