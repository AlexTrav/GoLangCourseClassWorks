package main

import (
	"fmt"
	"time"
)

func Producer(name string, ch chan<- int, limit int) {
	timer := time.NewTimer(1 * time.Second)
	defer timer.Stop()
	fmt.Println("producer " + name + " started")
	for i := 0; i < limit; i++ {
		timer.Reset(100 * time.Millisecond)
		<-timer.C
		fmt.Println("producer " + name + " writing message")
		ch <- i
	}
	fmt.Println("producer " + name + "finished")
	close(ch)
}

func FanIn(jobs1 <-chan int, jobs2 <-chan int, out chan<- int, cancel <-chan struct{}) {
	timer := time.NewTimer(12 * time.Second)
	defer timer.Stop()
	defer func() {
		fmt.Println("consumer finished")
	}()

	defer close(out)

	for {
		select {
		case v, ok := <-jobs1:
			if ok {
				println(v)
				out <- v
			}
		case v, ok := <-jobs2:
			if ok {
				println(v)
				out <- v
			}
		case <-timer.C:
			fmt.Println("timeout")
			return
		case <-cancel:
			fmt.Println("canceling")
			return
		}
	}
}

func main() {
	Job1 := make(chan int)
	Job2 := make(chan int)
	out := make(chan int, 1000)
	chancel := make(chan struct{})

	go Producer("p1", Job1, 100000000)
	go Producer("p2", Job2, 109000000)

	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("canceling consumer by timeout")
		chancel <- struct{}{}
	}()

	go func() {
		for i := range out {
			fmt.Println(i)
		}
		fmt.Println("out channel closed")
	}()

	FanIn(Job1, Job2, out, chancel)
}
