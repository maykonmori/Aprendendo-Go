package main

import "testing"

// It tests that the EquationLenCorrect function returns true when the equation is valid
func TestVerifyLen(t *testing.T) {
    expected := []bool{true, false}
    var test = []struct{
        EquationSent string
    }{
        {"30+5*3"},
        {"30+5*25"},
    }
    for i, v := range test {
        result := EquationLenCorrect(v.EquationSent)
        if result != expected[i] {
            t.Errorf("test %d: resultado '%t', esperado '%t'", i, result, expected[i])
        }
    } 
}

func TestValidateEquation(t *testing.T) {
    expected := []bool{true, false}
    var test = []struct{
        EquationSent string
    }{
        {"30+5*2"},
        {"30+5*A"},
    }
    for i, v := range test {
        result := ValidateEquation(v.EquationSent)
        if result != expected[i] {
            t.Errorf("test %d: resultado '%t', esperado '%t'", i, result, expected[i])
        }
    }
}

func TestComparisonEquation(t *testing.T) {
    equationSent := "32+5*2"
    equationCorrect := "32+5*2"
    expected := true
    result := ComparisonEquation(equationSent, equationCorrect)
    if result != expected {
        t.Errorf("resultado '%t', esperado '%t'", result, expected)
    }
}

func TestComparisonCharEquation(t *testing.T) {
    equationSent := "32+5*2"
    equationCorrect := "32+5*2"
    expected := "CCCCCC"
    result := ComparisonCharEquation(equationSent, equationCorrect)
    if result != expected {
        t.Errorf("resultado '%s', esperado '%s'", result, expected)
    }
}

func TestCalc(t *testing.T)  {
    equationSent := "32+5*2"
    // equationCorrect := "32+5*2"
    expected := 42
    result := calc(equationSent)
    if result != expected {
        t.Errorf("resultado '%d', esperado '%d'", result, expected)
    }
}
// func TestParseString(t *testing.T) {
//     expected := 30
//     result := ParseString("30+5*2")

//     if result != expected {
//         t.Errorf(" resultado '%d', esperado '%d'", result, expected)
//     }
// }