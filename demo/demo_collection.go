package demo

import (
	"fmt"
	"sort"
)

func DemoCollection() {

	fmt.Println("*************** Demo Collection ***************")

	demoSortIntArray()
}

func demoSortIntArray() {

	fmt.Println("--- Try Integer Sort Array ---")

	intValue := []int{57, 12, 20, 5, 8, 2}
	sort.Ints(intValue)
	fmt.Println("Ints:   ", intValue)
}
