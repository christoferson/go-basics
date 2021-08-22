package demo

import "fmt"

func describeIntSlice(s []int) {
	fmt.Printf("[]int length=%d capacity=%d %v\n", len(s), cap(s), s)
}
