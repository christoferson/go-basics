package demo

import (
	"fmt"
	"os"
)

func DemoPanic() {

	fmt.Println("*************** Demo Panic ***************")

	demoPanicText()

	demoPanicError()

}

func demoPanicText() {
	defer func() {
		e := recover()
		if e != nil {
			fmt.Printf("Recover from Panic. Msg=%v \n", e)
		}
	}()
	panic("Error Situation")
}

func demoPanicError() {

	defer func() {
		e := recover()
		if e != nil {
			fmt.Printf("Recover from Panic. Msg=%v \n", e)
		}
	}()

	_, err := os.Open("/tmp/file")
	if err != nil {
		panic(err)
	}
}
