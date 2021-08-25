package demo

import (
	"fmt"
)

func DemoDefer() {

	fmt.Println("*************** Demo Defer ***************")

	demoDefer()

}

func demoDefer() {
	defer fmt.Println("Defereeed Clause")
	fmt.Println("Foo")
}
