package demo

import (
	"fmt"
	"sort"
)

func DemoCollection() {

	fmt.Println("*************** Demo Collection ***************")

	demoSortIntArray()

	demoSortFloatArray()
}

func demoSortIntArray() {

	fmt.Println("--- Try Integer Sort Array ---")

	intValue := []int{57, 12, 20, 5, 8, 2}
	sort.Ints(intValue)
	fmt.Println("Ints:   ", intValue)

}

func demoSortFloatArray() {

	fmt.Println("--- Try Float Sort Array ---")

	floatValue := []float64{12.7, 25.1, 15.2, 18.32}
	sort.Float64s(floatValue)
	fmt.Println("floatValue:   ", floatValue)

}
