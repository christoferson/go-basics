package demo

func DemoControls() {
	fmt.Println("*************** Demo Controls ***************")

	var a int = 10
	if a%2 == 0 {
		fmt.Printf("a %v is even\n", a)
	} else {
		fmt.Printf("a %v is odd\n", a)
	}
}