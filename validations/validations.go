package validations

import (
	"regexp"
	"strings"
)

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

func IsDigit(equationSent int) bool {
    if equationSent >= 48 && equationSent <= 57 {
        return true
    }
	return false;
}