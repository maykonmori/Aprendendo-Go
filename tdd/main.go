package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// equationCorrect := "32+5*2"
	equationSent := "32+2*5"
    // Calling the function Soma and passing the variable equationSent as a parameter.
	fmt.Println(Calc(equationSent))
    // fmt.Println(EquationLenCorrect(equationSent))
}

func EquationLenCorrect(equationSent string) bool{
    return len(equationSent) == 6
}

func ValidateEquation(equationSent string) bool {
    valid, _ := regexp.MatchString("^[0-9/*+-]*$", equationSent)
    return valid
}

func ComparisonEquation(equationSent, equationCorrect string) bool  {
    return strings.Compare(equationSent,equationCorrect) == 0
}

func ComparisonCharEquation(equationSent, equationCorrect string) string  {
    result := ""
    for i := 0; i < 6; i++ {
        if equationSent[i] == equationCorrect[i]{
            result += "C"
        }else{
            result += VerifyCharPosition(equationSent, equationCorrect, i)
        }
    }
    return result
}

func VerifyCharPosition(equationSent, equationCorrect string, position int) string {
    for i := 0; i < len(equationCorrect); i++ {
        if equationSent[position] == equationCorrect[i]{
            return "T"
        }
    }
    return "X"
}

func Calc(equationSent string) int {

    for i := 0; i < len(equationSent); i++ {
        if equationSent[i] == '*' {
            fmt.Println(equationSent[i])
            i--
            fmt.Println(equationSent[i])
            for j := i; j < 6; j-- {
                fmt.Println(equationSent[j])
                if !IsDigit(int(equationSent[j])){
                    fmt.Println(equationSent[j])
                    break
                }   
            }
        }
    }
    
    return 42
}

func IsDigit(equationSent int) bool {
    if equationSent >= 48 && equationSent <= 57 {
        return true
    }
	return false;
}
// // //parsear a string e separar os numeros , do operadores
// func ParseString(equationSent string) []byte {
//     var slice []byte
//     for i := 0; i < 6; i++ {
//         if regexp.MatchString("^[0-9]*$", equationSent)
//             slice = append(slice, equationSent[i])
//     }
//     return slice
// }

