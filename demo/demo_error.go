package main

import (
	"errors"
)

func DemoErrors() {

	fmt.Println("*************** Demo Errors ***************")

	DemoErrorNew()
}

func DemoErrorNew() {
	var e = errors.New("Missing Value")
	fmt.Println(e)
}