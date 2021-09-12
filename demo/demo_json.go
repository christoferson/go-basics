package demo

import (
	"encoding/json"
	"fmt"
)

func DemoJSON() {

	fmt.Println("*************** Demo JSON ***************")

	demoMarshallBasic()

	demoMarshallStruct()

}

func demoMarshallBasic() {

	fmt.Println("--- Try Marshall Basic ---")

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

type mstruct1 struct {
	Position int
	Planet   []string
}
type mstruct2 struct {
	Position int      `json:"coordinate"`
	Planet   []string `json:"planets"`
}

func demoMarshallStruct() {

	fmt.Println("--- Try Marshall Struct ---")

	res1A := &mstruct1{Position: 1, Planet: []string{"mercury", "venus", "earth"}}
	res1B, _ := json.Marshal(res1A)
	fmt.Println(string(res1B))

	res2D := &mstruct2{Position: 1, Planet: []string{"mercury", "venus", "earth"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

}
