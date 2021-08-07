package demo

import "fmt"
import "strconv"


func DemoTypes() {
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

func demoContants() {

	const X int = 100
	const Y int = 200

	z := X * Y

	fmt.Printf("Type=%T Value=%v", z, z)

}

func demoCast() {

	var i int = 10
	var f float64 = 6.44

	x := float64(i)
	y := int(f)

	fmt.Printf("Type=%T Value=%v", x, x)
	fmt.Printf("Type=%T Value=%v", y, y)

	var str1 string = "101"
	var str2 string = "10.123"

	nint, _ := strconv.ParseInt(str1, 0, 64)
	nfloat, _ := strconv.ParseFloat(str2, 64)

	fmt.Printf("Type=%T Value=%v", nint, nint)
	fmt.Printf("Type=%T Value=%v", nfloat, nfloat)

}
