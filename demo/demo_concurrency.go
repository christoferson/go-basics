package demo

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func DemoConcurrency() {

	fmt.Println("*************** Demo Concurrency ***************")

	demoGoRoutine()
	fmt.Println()
	demoMutex()

}

// Demo Go Routine

func demoGoRoutine() {

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

/// Demo using Mutex

var waitMutex = sync.WaitGroup{}
var count int
var mutex sync.Mutex

func demoMutex() {

	waitMutex.Add(2)
	go increment("foo: ")
	go increment("bar: ")
	waitMutex.Wait()
	fmt.Println("Final count value ", count)

}

func increment(s string) {
	for i := 0; i < 3; i++ {
		mutex.Lock()
		x := count
		x++
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		count = x
		fmt.Println(s, i, "Count: ", count)
		mutex.Unlock()

	}
	waitMutex.Done()

}
