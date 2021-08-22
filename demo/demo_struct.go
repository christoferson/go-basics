package demo

import "fmt"

func DemoStruct() {

	fmt.Println("*************** Demo Struct ***************")

	demoStruct()

	demoStructEmbedded()

	demoStructMethod()

}

//

type Character struct {
	name string
	race string
}

func demoStruct() {
	c := Character{name: "Conan", race: "Human"}
	fmt.Println(c)
	fmt.Println(c.name)
}

//

type Character2 struct {
	name string
	race string
}

type NonPlayableCharacter2 struct {
	Character2
	npcid string
}

func demoStructEmbedded() {
	c := NonPlayableCharacter2{npcid: "578", Character2: Character2{name: "Conan", race: "Human"}}
	fmt.Println(c)
	fmt.Println(c.npcid)
}

//
type Character3 struct {
	name string
	race string
}

func (c Character3) greet() {
	fmt.Printf("C3 %v says hello", c.name)
}

func demoStructMethod() {
	c := Character3{name: "Conan", race: "Human"}
	c.greet()
	fmt.Printf("Name: %v", c.name)
	fmt.Println()
}

//

type Character4 struct {
	name string
	race string
}

func (c *Character4) greet() {
	fmt.Printf("C4 %v says hello", c.name)
}

func demoStructMethodPointer() {
	c := &Character4{name: "Conan", race: "Human"}
	c.greet()
	fmt.Printf("Name: %v", c.name)
	fmt.Println()
}
