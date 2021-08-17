package demo

import(  
	"fmt"  
	"reflect"  
 )  

func DemoReflect() {

	fmt.Println("*************** Demo Reflect ***************")

	demoReflect()
	
}

func demoReflect() {
	age := 27.5  
	fmt.Printf("%T\n" ,age)  
	fmt.Println(reflect.TypeOf(age))  
}