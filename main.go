package main

import (
	"fmt"
	// "strings"
	"regexp"
	
	// "os"
)

// The function play takes two strings as arguments, equationCorrect and equationSent. It returns a
// boolean value
func main ()  {
	fmt.Println("Hello")

	equationCorrect := "30+5*2"
	equationSent := "30+5*A"
	fmt.Println(Play(equationCorrect, equationSent))
}

// It receives two strings, one with the equation sent by the user and the other with the correct
// equation, and returns a string with the result of the game
func Play (equationCorrect, equationSent string) string{
	fmt.Println(equationSent)
	fmt.Println(equationCorrect)
	equationSentLen := len(equationSent)
	fmt.Println("tamanho:", equationSentLen)
	if equationSentLen != 6 {
		return "TAMANHO ERRADO"
	}
	match, _ := regexp.MatchString("^[0-9/*+-]*$", equationSent)
	if !match {	
		return "EQUAÇÃO INVÁLIDA"
	}
	// expr , err:=eval.ParseString(equationSent,"")
	// if err!=nil{
	// 	return err
	// }
	// r , err:=expr.EvalToInterface(nil)
	// if err!=nil{
    // 	return err
	// }
	// fmt.Println(r)
	return "DEU TUDO CERTO"
}
