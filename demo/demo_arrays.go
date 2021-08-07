package demo

func DemoArrays() {

	fmt.Println("*************** Demo Arrays ***************")

	demoArray()

	demoArrayMultidimensional()

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
