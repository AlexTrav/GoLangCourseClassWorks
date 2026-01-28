package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
)

const OperationIn = "in"
const OperationOut = "out"

type Balance struct {
	mu    sync.RWMutex
	value int
	Limit int
}

func (w *Balance) Pay(sum int) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.value >= sum {
		w.value -= sum
		return nil
	}
	return InsufficientFundsError
}
func (w *Balance) Deposit(sum int) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.value+sum <= w.Limit {
		w.value += sum
		return nil
	}
	return OperationLimitError
}
func (w *Balance) GetValue() int {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.value
}
func (w *Balance) GetLimit() int {
	return w.Limit
}

type Payment struct {
	Id        int
	Operation string // in || out
	Amount    int
}

var (
	InsufficientFundsError = errors.New("insufficient funds")
	OperationLimitError    = errors.New("operation limit exceeded")
)

func processPayments(wallet *Balance, payment Payment) error {
	switch payment.Operation {
	case OperationIn:
		return wallet.Deposit(payment.Amount)
	case OperationOut:
		return wallet.Pay(payment.Amount)
	}
	return nil
}

func successWorker(wg *sync.WaitGroup, successes <-chan Payment) {
	defer wg.Done()

	for payment := range successes {
		fmt.Println(fmt.Sprintf("info: success payment %d ; amount %d", payment.Id, payment.Amount))
	}
	fmt.Println("success channel closed")
}
func failureWorker(wg *sync.WaitGroup, failures <-chan error) {
	defer wg.Done()

	for err := range failures {
		fmt.Println(fmt.Sprintf("Log: error %s", err))
		if errors.Is(err, InsufficientFundsError) {
			fmt.Println(fmt.Sprintf("critical : %s", err.Error()))
		}
		if errors.Is(err, OperationLimitError) {
			fmt.Println(fmt.Sprintf("warning: %s", err.Error()))
		}
	}
	fmt.Println("failure channel closed")
}

func worker(n int, wg *sync.WaitGroup, wallet *Balance, payments <-chan Payment, successes chan<- Payment, failures chan<- error) {
	defer wg.Done()

	for p := range payments {
		fmt.Println(fmt.Sprintf("worker %d processing payment %d", n, p.Id))
		err := processPayments(wallet, p)

		if err != nil {
			if errors.Is(err, InsufficientFundsError) {
				b := wallet.GetValue()
				fmt.Println(fmt.Sprintf("insufficient funds %d", b))
			}
			if errors.Is(err, OperationLimitError) {
				fmt.Println(fmt.Sprintf("operation limit exceeded %d", wallet.GetLimit()))
			}
			failures <- err
		} else {
			successes <- p
		}
	}
}

func generatePayments(count int, payments chan<- Payment) {
	operation := ""

	defer close(payments)

	for i := 1; i <= count; i++ {
		n := rand.Intn(100)
		if n > 50 {
			operation = OperationOut
		} else {
			operation = OperationIn
		}

		payments <- Payment{i, operation, rand.Intn(1000)}
	}
}

func main() {

	payments := make(chan Payment)
	successes := make(chan Payment)
	failures := make(chan error)

	workerCount := 20
	wallet := Balance{value: 100000, Limit: 10000000}

	var wg sync.WaitGroup
	var cg sync.WaitGroup

	cg.Add(2)
	go successWorker(&cg, successes)
	go failureWorker(&cg, failures)

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(i, &wg, &wallet, payments, successes, failures)
	}

	go generatePayments(10000, payments)

	wg.Wait()
	close(failures)
	close(successes)
	cg.Wait()

	fmt.Println("done")
}
