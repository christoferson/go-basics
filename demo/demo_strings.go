package demo

import "reflect"
import "strings"

func DemoStrings() {

	fmt.Println("*************** Demo Strings ***************")

	demoStrings()

	demoStringsLength()

	demoStringsLibrary()

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


func demoStringsLibrary() {

	var x string = "sample string"  
	fmt.Println(strings.ToUpper(x))
	fmt.Println(strings.ToLower(x))

	fmt.Println(strings.HasSuffix(x, "string"))  

	var arr = []string{"a","b","c","d", "e"}  
	fmt.Println(strings.Join(arr,"*")) // a*b*c*d*e

	var str = "abc " 
	fmt.Println(strings.Repeat(str, 5))  

	str2:= "Loren ipsum"  
	fmt.Println(strings.Contains(str2, "ips"))  

}