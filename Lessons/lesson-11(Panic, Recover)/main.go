package main

import (
	"errors"
	"fmt"
	"runtime/debug"
)

var panicError = errors.New("panic error")

func rec() {
	if r := recover(); r != nil {
		fmt.Println("recover")
		if e := recover(); e != nil {
			err := r.(error)
			if errors.Is(err, panicError) {
				fmt.Println(panicError, r)
			}
		} else {
			fmt.Println(string(debug.Stack()))
		}

		fmt.Println("close db ...")
	}
}

func main() {
	fmt.Println("start")

	defer rec()

	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")

	err := safeCall(f, "some string")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("end")
}

func safeCall(fn func(a string) error, a string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	return fn(a)
}

func f(a string) error {
	defer fmt.Println("defer f ", a)
	panic(struct{ Code int }{Code: 1000})
	return nil
}
