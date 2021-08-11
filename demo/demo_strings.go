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

	str2 := "Loren ipsum"  
	fmt.Println(strings.Contains(str2, "ips"))  // true

	str3 := "Loren ipsum" 
	fmt.Println(strings.Index(str3, "ips"))  // 6

	str4 := "Loren ipsume"  
	fmt.Println(strings.Count(str4, "e"))  // 2

	str5 := "Loren ipsumeDe"  
	fmt.Println(strings.Replace(str5, "e", "Z", 2))	 // LorZn ipsumZDe

	fmt.Println(strings.Compare("123", "456"))  
	fmt.Println(strings.Compare("abc", "abc"))  
	fmt.Println(strings.Compare("abc", "ABC"))  

	fmt.Println("'" + strings.TrimSpace(" \t\n Abns djsie siiens Ajdd  \n\t\r\n") + "'") 


	fmt.Println(strings.ContainsAny("Apple", "A"))  
	fmt.Println(strings.ContainsAny("Algebra", "A & g"))  
	fmt.Println(strings.ContainsAny("Random", ""))  
	fmt.Println(strings.ContainsAny("", ""))  

}