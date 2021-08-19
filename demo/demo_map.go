package demo


func DemoMap() {

	fmt.Println("*************** Demo Map ***************")

	demoMap()

	demoMapMake()
	
}

func demoMap() {

	characters := map[string]int{"Conan":28, "Natalya":37, "Griswold":20}  
	fmt.Println(characters)  
	fmt.Printf("characters['Natalya']=%v \n", characters["Natalya"])  

}

func demoMapMake() {

	m := make(map[string]int)  
	fmt.Println(m)  
	m["Conan"] = 10  
	m["Natalya"] = 20  
	m["Griswold"] = 30  
	fmt.Println(m)  
	m["Natalya"] = 24
	fmt.Println(m)  

}