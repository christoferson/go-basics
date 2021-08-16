package demo


func DemoInterface() {

	fmt.Println("*************** Demo Interface ***************")

	demoInterfaceMethod()
	
}

//

type ICharacter interface {  
	Attack()  
}

type Character31 struct {  
	name  string  
	race  string  
}

func (c Character31) Attack()  {  
	fmt.Printf("C31 %v attacks.", c.name)   
}

func icAttack(ic ICharacter) {
	fmt.Println("C31 Calling ICharacter.Attack")
	ic.Attack()  
}

func demoInterfaceMethod() {
	c := Character31{name: "Conan", race: "Human"}  
	c.Attack()
	fmt.Println()
	var i ICharacter = c
	i.Attack()
	fmt.Println()
	icAttack(c)
	fmt.Println()
}

//