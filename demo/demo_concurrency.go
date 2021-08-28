package demo

import (
	"fmt"
	"sync"
	"time"
)

func DemoConcurrency() {

	fmt.Println("*************** Demo Concurrency ***************")

	demoConcurrency()

}

func demoConcurrency() {

	wg.Add(2)
	go routine1()
	go routine2()
	wg.Wait()

}

var wg = sync.WaitGroup{}

func routine1() {
	for i := 0; i < 5; i++ {
		fmt.Println("routine1,  ->", i)
		time.Sleep(time.Duration(5 * time.Millisecond))
	}
	wg.Done()
}

func routine2() {
	for i := 0; i < 5; i++ {
		fmt.Println("routine2,  ->", i)
		time.Sleep(time.Duration(5 * time.Millisecond))
	}
	wg.Done()
}
