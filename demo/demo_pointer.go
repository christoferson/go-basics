package demo

func DemoPointer() {

	fmt.Println("*************** Demo Pointer ***************")

	demoPointer()
	
}

//

func demoPointer() {

	x := 10

	modifyIntPointer(&x)

	fmt.Printf("Value of X is %v", x)

}

func modifyIntPointer(x *int) {  
	*x = 5
}

//