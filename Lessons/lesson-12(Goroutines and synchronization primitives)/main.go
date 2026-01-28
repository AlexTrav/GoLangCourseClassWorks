package main

import (
	"fmt"
	"sync"
	"time"
)

type Balance struct {
	mu     sync.RWMutex
	UserId int
	value  int
}

func (w *Balance) Payout(sum int) {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.value >= sum {
		w.value -= sum
	}
}
func (w *Balance) PayIn(sum int) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.value += sum
}

func (w *Balance) GetValue() int {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.value
}

func workerPayOut(wg *sync.WaitGroup, i int, wallet *Balance, sum int) {
	// defer wg.Done()
	wallet.Payout(sum)
	fmt.Println(fmt.Sprintf("PayOut transaction id: %d, wallet balance %d", i, wallet.value))
}

func workerPayIn(wg *sync.WaitGroup, i int, wallet *Balance, sum int) {
	// defer wg.Done()
	wallet.PayIn(sum)
	fmt.Println(fmt.Sprintf("PayIn transaction id: %d, wallet balance %d", i, wallet.value))
}

func main() {
	timeStart := time.Now()
	// 1 ядро
	var wg sync.WaitGroup

	workers := 100000

	sum := 0

	wallet := Balance{sync.RWMutex{}, 1, 1000}

	for transactionId := 0; transactionId < workers; transactionId++ {
		wg.Add(1)
		sum = 15
		go func() {
			defer wg.Done()
			workerPayIn(&wg, transactionId, &wallet, sum)
			workerPayOut(&wg, transactionId, &wallet, sum)
			workerPrintBalance(&wg, &wallet)
		}()
	}
	timeFinish := time.Now()

	wg.Wait()

	fmt.Println("total wallet balance ", wallet.value)
	fmt.Println("duration ", timeFinish.Sub(timeStart))
}

func workerPrintBalance(wg *sync.WaitGroup, wallet *Balance) {
	// defer wg.Done()

	fmt.Println(fmt.Sprintf("wallet balance %d", wallet.GetValue()))
}
