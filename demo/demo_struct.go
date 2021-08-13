package demo

func DemoStruct() {

	fmt.Println("*************** Demo Struct ***************")

	demoStruct()
	
}

	
func demoStruct() {
	c := Character{name: "Conan", race: "Human", }  
	fmt.Println(c)  
	fmt.Println(c.name) 
}


type Character struct {  
	name string  
	race  string  
 }  
