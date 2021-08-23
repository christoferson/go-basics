package demo

import "fmt"

func DemoPointer() {

	fmt.Println("*************** Demo Pointer ***************")

	demoPointer()

	fmt.Println()

	demoPointerRepoint()

	fmt.Println()

}

//

func demoPointer() {

	x := 10

	modifyIntPointer(&x)

	fmt.Printf("Value of X is %v \n", x)

}

func modifyIntPointer(x *int) {
	*x = 5
}

//

func demoPointerRepoint() {

	x := 10

	fmt.Printf("Before: Value of X is %v \n", x)

	repointIntPointer(&x)

	fmt.Printf("After: Value of X is %v \n", x)

}

func repointIntPointer(x *int) {
	*x = 5
}

//
