package demo

func DemoArrays() {

	fmt.Println("*************** Demo Arrays ***************")

	demoArray()

	demoArrayMultidimensional()

	demoSlice()

	demoMultipleSlice()

	demoSliceLiteral()

}


func demoArray() {

	var x [5]int
	var i, j int
	for i = 0; i < 5; i++ {
		x[i] = i + 10
	}
	for j = 0; j < 5; j++ {
		fmt.Printf("Item[%d] = %d\n", j, x[j])
	}

}

func demoArrayMultidimensional() {

	var a = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	var i, j int

	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			fmt.Print(a[i][j])
		}
		fmt.Println()
	}

}


func demoSlice() {

	odd := [6]int{1, 3, 5, 7, 9, 11}
	var s []int = odd[1:4]
	fmt.Println(s)
	odd[1] = 13
	fmt.Println(s)
	s[0] = 15
	fmt.Println(s)
	fmt.Println(odd)
}

func demoMultipleSlice() {

	names := [4]string{"Apple", "Orange", "Banana",  "Kiwi"} 
	 fmt.Println(names)
	 slice1 := names[0:2] 
	 slice2 := names[1:3] 
	 fmt.Println(slice1, slice2) 
	 slice2[0] = "Juice"
	 fmt.Println(slice1, slice2) 
	 fmt.Println(names)

}

func demoSliceLiteral() {

	s := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
		{3, true},
		{4, true},
		{5, false},
		{6, true},
	}
	fmt.Println(s)

}
