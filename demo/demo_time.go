package demo

import (
	"fmt"
	"time"
)

func DemoTime() {

	fmt.Println("*************** Demo Time ***************")
	demoCurrentTime()

	demoTimeComponents()

	demoTimeComparison()
}

func demoCurrentTime() {
	fmt.Println("--- Try Current Time ---")
	p := fmt.Println

	present := time.Now() // current time
	p(present)

}

func demoTimeComponents() {
	fmt.Println("--- Try Time Components ---")
	DOB := time.Date(2021, 07, 28, 15, 27, 42, 517, time.Local)
	fmt.Println(DOB)

	fmt.Println(DOB.Year())
	fmt.Println(DOB.Month())
	fmt.Println(DOB.Day())
	fmt.Println(DOB.Hour())
	fmt.Println(DOB.Minute())
	fmt.Println(DOB.Second())
	fmt.Println(DOB.Nanosecond())
	fmt.Println(DOB.Location())

	fmt.Println(DOB.Weekday())

}

func demoTimeComparison() {

	fmt.Println("--- Try Time Comparison ---")
	DOB := time.Date(2021, 07, 28, 15, 27, 42, 517, time.Local)
	present := time.Now()

	fmt.Println(DOB.Before(present))
	fmt.Println(DOB.After(present))
	fmt.Println(DOB.Equal(present))
}
