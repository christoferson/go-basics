package demo

import (
	"errors"
	"fmt"
)

func DemoErrors() {

	fmt.Println("*************** Demo Errors ***************")

	DemoErrorNew()
	DemoErrorF()
}

func DemoErrorNew() {
	var e = errors.New("Missing Value")
	fmt.Println(e)
}

func DemoErrorF() {
	var f = 829.2
	var e = fmt.Errorf("math: square root of negative number %g", f)
	fmt.Println(e)
}
