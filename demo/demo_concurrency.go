package demo

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func DemoConcurrency() {

	fmt.Println("*************** Demo Concurrency ***************")

	demoGoRoutine()
	fmt.Println()
	demoMutex()
	fmt.Println()
	demoAtomic()
	fmt.Println()
	demoChannel()
	fmt.Println()
	demoWorkerPool()

}

// Demo Go Routine

func demoGoRoutine() {
	fmt.Println("--- Try Go Routine ---")
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
	fmt.Println("--- Try Mutex ---")
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

// Demo Atomic

var waitAtomic sync.WaitGroup
var countAtomic int64

func demoAtomic() {
	fmt.Println("--- Try Atomic Variable ---")
	waitAtomic.Add(2)
	go incrementAtomic("foo: ")
	go incrementAtomic("bar: ")
	waitAtomic.Wait()
	fmt.Println("last count value ", countAtomic)
}

func incrementAtomic(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Duration((rand.Intn(3))) * time.Millisecond)
		atomic.AddInt64(&countAtomic, 1)
		fmt.Println(s, i, "Count ->", countAtomic)
	}
	waitAtomic.Done()
}

// Demo Channel

func demoChannel() {
	fmt.Println("--- Try Channel ---")
	done := make(chan bool, 1)
	go demoChannelWork(done)
	<-done
}

func demoChannelWork(done chan bool) {
	fmt.Println("Start Channel Work...")
	for i := 0; i < 2; i++ {
		time.Sleep(time.Second)
		fmt.Printf("Do Channel Work : %v \n", i)
	}
	fmt.Println("End Channel Work")
	done <- true
}

// Try Worker Pool

func demoWorkerPoolTask(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func demoWorkerPool() {
	fmt.Println("--- Try Worker Pool ---")
	job := make(chan int, 10)
	result := make(chan int, 10)
	for w := 1; w <= 2; w++ {
		go demoWorkerPoolTask(w, job, result)
	}
	for j := 1; j <= 5; j++ {
		job <- j
	}
	fmt.Println("Finished Job Submission.")
	close(job)
	for a := 1; a <= 5; a++ {
		<-result
	}
	fmt.Println("Finished Jobs.")
}
