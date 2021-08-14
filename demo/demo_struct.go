package demo

func DemoStruct() {

	fmt.Println("*************** Demo Struct ***************")

	demoStruct()

	demoStructEmbedded()
	
}


type Character struct {  
	name string  
	race  string  
}  
	
func demoStruct() {
	c := Character{name: "Conan", race: "Human", }  
	fmt.Println(c)  
	fmt.Println(c.name) 
}

type Character2 struct {  
	name string  
	race  string  
}

type NonPlayableCharacter2 struct {  
	Character2
	npcid  string  
}

func demoStructEmbedded() {
	c := NonPlayableCharacter2{npcid: "578", Character2:Character2{name: "Conan", race: "Human"}}  
	fmt.Println(c)  
	fmt.Println(c.npcid) 
}