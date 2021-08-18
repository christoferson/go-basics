package demo

import(  
	"fmt"  
	"reflect"  
 )  

func DemoReflect() {

	fmt.Println("*************** Demo Reflect ***************")

	demoReflect()

	demoRune()
	
}

func demoReflect() {
	age := 27.5  
	fmt.Printf("%T\n" ,age)  
	fmt.Println(reflect.TypeOf(age))  
}

func demoRune() {
	rune := 'C'    
	fmt.Printf("Type=%T Code=%d \n", rune, rune)    
	fmt.Println(reflect.TypeOf(rune))
}