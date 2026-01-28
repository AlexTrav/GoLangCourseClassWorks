package main

import (
	"fmt"
	"time"
)

func producer(out chan<- int) {
	timer := time.NewTimer(1 * time.Second)
	defer timer.Stop()

	for i := 0; i < 10; i++ {
		timer.Reset(1 * time.Second)
		<-timer.C
		out <- i
	}
	close(out)

}

func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int)
	go producer(ch)
	Consumer(ch)
}
