package demo

import (
	"encoding/json"
	"fmt"
)

func DemoJSON() {

	fmt.Println("*************** Demo JSON ***************")

	demoMarshallBasic()

}

func demoMarshallBasic() {

	bolType, _ := json.Marshal(false) //boolean Value
	fmt.Println(string(bolType))
	intType, _ := json.Marshal(10) // integer value
	fmt.Println(string(intType))
	fltType, _ := json.Marshal(3.14) //float value
	fmt.Println(string(fltType))
	strType, _ := json.Marshal("mstring") // string value
	fmt.Println(string(strType))
	slcA := []string{"apple", "orange", "banana"} //slice value
	slcB, _ := json.Marshal(slcA)
	fmt.Println(string(slcB))
	mapA := map[string]int{"sword": 15, "dagger": 12} //map value
	mapB, _ := json.Marshal(mapA)
	fmt.Println(string(mapB))

}
