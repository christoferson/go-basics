package demo

func DemoControls() {
	fmt.Println("*************** Demo Controls ***************")

	demoIfElse()
}

func demoIfElse() {
	var a int = 10
	if a%2 == 0 {
		fmt.Printf("a %v is even\n", a)
	} else {
		fmt.Printf("a %v is odd\n", a)
	}
}

func demoIfElseIf() {
	var a int = 9
	if a%2 == 0 {
		fmt.Printf("a %v is divisible by 2\n", a)
	} else if a%3 == 0 {
		fmt.Printf("a %v is divisible by 3\n", a)		
	} else {
		fmt.Printf("a %v is not divisible by 2 or 3 \n", a)
	}
}