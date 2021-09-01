package demo

import (
	"fmt"
	"time"
)

func DemoTime() {

	fmt.Println("*************** Demo Time ***************")
	demoCurrentTime()
}

func demoCurrentTime() {
	fmt.Println("--- Try Current Time ---")
	p := fmt.Println

	present := time.Now() // current time
	p(present)

}
