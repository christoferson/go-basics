package demo

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func DemoTime() {

	fmt.Println("*************** Demo Time ***************")
	demoCurrentTime()

	demoTimeComponents()

	demoTimeComparison()

	demoTimeArithmetic()

	demoTimeFormat()

	demoTimeEpoch()

	demoTimeTicker()

	demoTimeTickerWithGoRoutine()
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

func demoTimeArithmetic() {
	fmt.Println("--- Try Time Arithmetic ---")
	DOB := time.Date(2021, 07, 28, 15, 27, 42, 517, time.Local)
	present := time.Now()

	diff := present.Sub(DOB)
	fmt.Println(diff)
	fmt.Println(diff.Hours())
	fmt.Println(diff.Minutes())
	fmt.Println(diff.Seconds())
	fmt.Println(diff.Nanoseconds())
	fmt.Println(DOB.Add(diff))
	fmt.Println(DOB.Add(-diff))

}

func demoTimeFormat() {
	fmt.Println("--- Try Time Format ---")
	present := time.Now()

	fmt.Println(present.Format("2006-01-02"))
	fmt.Println(present.Format("Monday January 02, 2006"))
	fmt.Println(present.Format("Mon 02 January 2006"))

	fmt.Println(present.Format("15:04:05"))
	fmt.Println(present.Format("03:04:05 PM"))

}

func demoTimeEpoch() {

	fmt.Println("--- Try Time Epoch ---")

	p := fmt.Println
	current_time := time.Now()
	secs := current_time.Unix()
	nanos := current_time.UnixNano()

	fmt.Println(current_time)

	millis := nanos / 1000000
	p(secs)
	p(millis)
	p(nanos)
	p(time.Unix(secs, 0))
	p(time.Unix(0, nanos))

}

func demoTimeTicker() {

	fmt.Println("--- Try Time Ticker ---")

	tickerValue := time.NewTicker(time.Millisecond * 100)
	go func() {
		for t := range tickerValue.C {
			fmt.Println("Tick at", t)
		}
	}()
	time.Sleep(time.Millisecond * 500)
	tickerValue.Stop()
	fmt.Println("Ticker stopped")

}

// Try Time Ticker 2

var waitTicker sync.WaitGroup

func demoTimeTickerJob() {
	i := 1
	for range time.Tick(2 * time.Second) {
		log.Printf("\tA: Let's have fun: %v %v ", "\t B : Okay!", i)
		i = i + 1
		if i > 3 {
			break
		}
	}
	waitTicker.Done()
	log.Printf("\tComplete")
}

func demoTimeTickerWithGoRoutine() {
	fmt.Println("--- Try Time Ticker on Go Routine ---")
	waitTicker.Add(1)
	go demoTimeTickerJob()
	//select {} //The select statement lets a goroutine wait on multiple communication operations.
	waitTicker.Wait()
}
