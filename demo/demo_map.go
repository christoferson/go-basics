package demo


func DemoMap() {

	fmt.Println("*************** Demo Map ***************")

	demoMap()


	
}

func demoMap() {

	characters := map[string]int{"Conan":28, "Natalya":37, "Griswold":20}  
	fmt.Println(characters)  
	fmt.Printf("characters['Natalya']=%v \n", characters["Natalya"])  

}

