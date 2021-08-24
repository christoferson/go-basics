package demo

import "fmt"

func DemoRecover() {

	fmt.Println("*************** Demo Recover ***************")

	demoRecover()

}

func demoRecover() {

	fmt.Printf("Return Value: %v \n", demoRecoverDivide(1, 0))
	fmt.Printf("Return Value: %v \n", demoRecoverDivide(30, 10))

}

func demoRecoverDivide(num1, num2 int) int {
	defer func() {
		e := recover()
		if e != nil {
			fmt.Printf("Recover from Exception. Msg=%v \n", e)
		}
	}()
	quotient := num1 / num2
	return quotient
}
