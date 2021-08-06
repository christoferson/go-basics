package demo

func demoTypes() {
	fmt.Println("*************** Demo Types ***************")

	var i int = 35
	fmt.Printf("Type=%T Value=%v", i, i)
	var f float64 = 76.87
	fmt.Printf("Type=%T Value=%v", f, f)
	var b bool = true
	fmt.Printf("Type=%T Value=%v", b, b)
	var s string = "ipsum"
	fmt.Printf("Type=%T Value=%v", s, s)
}