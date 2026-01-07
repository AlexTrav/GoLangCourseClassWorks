package counter

import (
	"fmt"
	cfg "testPackage/config"
)

func init() {
	Interval = cfg.Cfg.Counter.Interval
	fmt.Println("counter initialized")
}

func init() {
	fmt.Println("test init 1")
}

func init() {
	fmt.Println("test init 2")
}

func init() {
	fmt.Println("test init 3")
}
