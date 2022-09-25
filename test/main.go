package main

import (
	"fmt"
	"eval"
)

func main() {
	src:="int8(1*(1+2))"
	expr,err:=ParseString(src,"")
	if err!=nil{
		return err
	}
	r,err:=expr.EvalToInterface(nil)
	if err!=nil{
		return err
	}
	fmt.Printf("%v %T", r, r)	// "3 int8"
}