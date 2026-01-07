package main

import (
	"fmt"
	"github.com/AlexTrav/greeting"
	"github.com/google/uuid"
)

func main() {
	fmt.Println(uuid.New().String())
	fmt.Println(greeting.Greet())
	fmt.Println(greeting.Bye())
}
