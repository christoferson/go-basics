package demo

import "reflect"

func DemoStrings() {

	fmt.Println("*************** Demo Strings ***************")

	demoStrings()

	demoStringsLength()

}


func demoStrings() {

	var x string = "sample string"  
	fmt.Println(x)  
	fmt.Println(reflect.TypeOf(x))  

}

func demoStringsLength() {

	var x string = "sample string"  
	fmt.Println(len(x))

}
