package demo

func DemoControls() {
	fmt.Println("*************** Demo Controls ***************")

	demoIfElse()
	demoElseIf()
	demoCase()
	demoForLoop()
	demoForRange()
}

func demoIfElse() {
	var a int = 10
	if a%2 == 0 {
		fmt.Printf("a %v is even\n", a)
	} else {
		fmt.Printf("a %v is odd\n", a)
	}
}

func demoElseIf() {
	var a int = 9
	if a%2 == 0 {
		fmt.Printf("a %v is divisible by 2\n", a)
	} else if a%3 == 0 {
		fmt.Printf("a %v is divisible by 3\n", a)		
	} else {
		fmt.Printf("a %v is not divisible by 2 or 3 \n", a)
	}
}

func demoCase() {
	var a int = 10
	switch a {
	case 10:
		fmt.Print("the value is 10")
	case 20:
		fmt.Print("the value is 20")
	default:
		fmt.Print("Value is not 10,20")
	}

	switch a {
	case 10:
		fmt.Print("the value is 10")
		fallthrough
	case 20:
		fmt.Print("the value is 20")
		fallthrough
	default:
		fmt.Print("Value is not 10,20")
	}
}

func demoForLoop() {
	for a := 0; a < 11; a++ {
		fmt.Print(a, "-")
	}

	sum := 1  
	for sum < 32 {  
	   sum += sum  
	   fmt.Print(sum, "-")  
	}  
}

func demoForRange() {

	nums := []int{2, 3, 4}  
	sum := 0  
	for _, value := range nums {
	   sum += value  
	}  
	fmt.Println("sum:", sum)  

	kvs := map[string]string{"1":"mango","2":"apple","3":"banana"}  
   for k, v := range kvs {  
      fmt.Printf("%s -> %s\n", k, v)  
   } 

}