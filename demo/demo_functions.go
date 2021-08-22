package demo

import "fmt"

func DemoFunctions() {

	fmt.Println("*************** Demo Functions ***************")

	demoFunction(35, 89)

	ret := demoFunctionWithReturn(27, 32)
	fmt.Printf("DemoFunctionWithReturn.Ret. %v", ret)

	err1, ret1 := demoFunctionWithMultiReturn(50, 23)
	fmt.Printf("demoFunctionWithMultiReturn.Ret. %v %v", err1, ret1)

	err2, ret2 := demoFunctionWithMultiReturn(50, 23)
	fmt.Printf("demoFunctionWithMultiReturn.Ret. %v %v", err2, ret2)

	demoRecursion()

}

func demoFunction(x int, y int) {
	fmt.Printf("DemoFunction. %v %v", x, y)
}

func demoFunctionWithReturn(x int, y int) int {
	fmt.Printf("DemoFunctionWithReturn. %v %v", x, y)
	return x * y
}

func demoFunctionWithMultiReturn(x int, y int) (error, int) {
	if x >= y {
		err := fmt.Errorf("X %v must be less than y %v", x, y)
		return err, 0
	}
	fmt.Printf("DemoFunctionWithMultiReturn. %v %v", x, y)
	return nil, x * y
}

func demoRecursion() {
	ret := factorial(7)
	fmt.Printf("DemoRecursion. %v", ret)
	fmt.Println()
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func demoClosure() {
	number := 12
	squareNum := func() int {
		number *= number
		return number
	}
	fmt.Printf("DemoClosure. %v", squareNum())
	fmt.Println()
}
