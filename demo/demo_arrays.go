package demo

func DemoArrays() {

	fmt.Println("*************** Demo Arrays ***************")

	demoArray()

	demoArrayMultidimensional()

	demoSlice()

	demoMultipleSlice()

	demoSliceLiteral()

	demoSliceBounds()

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
		{101, false},
		{201, false},
		{301, true},
		{401, false},
		{501, true},
		{601, false},
	}
	fmt.Println(s)

}

func demoSliceBounds() {

	slice1 := []int{1,3,5,7,9,11,13}  
	slice2 := slice1[2:4]
	fmt.Println(slice2)  
	slice3 := slice1[:3]  
	fmt.Println(slice3)  
	slice4 := slice1[2:]  
	fmt.Println(slice4)  
	fmt.Println(slice1)  

}


func demoSliceLengthCapacity() {

	slice1 := []int{1,3,5,7,9,11,13}
	describeIntSlice(slice1)

	slice2 := slice1[:0]  
	describeIntSlice(slice2)

	slice3 := slice1[:4]
	describeIntSlice(slice3)  

	slice4 := slice1[2:]  
	describeIntSlice(slice4)  

}

func demoSliceMake() {

	slice1 := make([]int, 12)  
	describeIntSlice(slice1)

	slice2 := make([]int, 0, 15)
	describeIntSlice(slice2)

	slice3 := slice1[:5]  
	describeIntSlice(slice3) 
	
	slice4 := slice2[2:5]  
	describeIntSlice(slice4)

}