package main

import (
	"fmt"
	"math"
)

//func swap(x, y string) (string, string) {
//	return y, x
//}

//func arr(x, y string) []string {
//	return []string{x, y}
//}
//
//func main() {
//	var z, x int = 1, 2
//
//	a, b, c := swap("a", "b")
//	fmt.Println(a, b, c)
//	fmt.Println("Hello World!")
//	fmt.Println(arr("a", "b"))
//	fmt.Println(z, x)
//}

//func split(sum int) (x, y int) {
//	x = sum * 4 / 9
//	y = sum - x
//	return
//}

//func add(a, b int) int {
//	return a + b
//}
//
//func split(sum int) (x, y int) {
//	x = sum * 4 / 10
//	y = sum - x
//	return
//}

func main() {
	//fmt.Println("Hello, Go!")
	//fmt.Println("--------------------")
	//fmt.Printf("Square root of 64 is %d\n", int(math.Sqrt(64)))
	//fmt.Println("--------------------")
	//fmt.Println(math.Pi)
	//fmt.Println("--------------------")
	//fmt.Println(add(1, 2))
	//fmt.Println("--------------------")
	//fmt.Println(swap("hello", "world"))
	//fmt.Println("--------------------")
	//fmt.Println(split(10))
	//fmt.Println("--------------------")
	//
	//var myStr = "Hello, Go!"
	//var myInt = 10
	//var myBool = false
	//
	//fmt.Println(myStr, myInt, myBool)
	//
	//fmt.Println("--------------------")
	//
	//name, age, height := "Alex", 22, 175
	//fmt.Println(name, age, height)
	//
	//fmt.Println("--------------------")
	//
	//var b bool
	//var s string
	//var i int
	//var f float64
	//var c64 complex64
	//var r rune
	//var by byte
	//
	//fmt.Printf("%T\n", b)
	//fmt.Printf("%T\n", s)
	//fmt.Printf("%T\n", i)
	//fmt.Printf("%T\n", f)
	//fmt.Printf("%T\n", c64)
	//fmt.Printf("%T\n", r)
	//fmt.Printf("%T\n", by)
	//
	//fmt.Println("--------------------")
	//
	//fmt.Println(b)
	//fmt.Println(s)
	//fmt.Println(i)
	//fmt.Println(f)
	//fmt.Println(c64)
	//fmt.Println(r)
	//fmt.Println(by)
	//
	//fmt.Println("--------------------")
	//
	//a1 := 5
	//b1 := 2
	//
	//var f1 = float64(a1)
	//var f2 = float64(b1)
	//
	//fmt.Println(f1 / f2)
	//
	//fmt.Println("--------------------")
	//
	//const (
	//	Pi   = 3.14
	//	Lang = "Go"
	//)
	//
	//fmt.Printf("The best language is %s with Pi = %f\n", Lang, Pi)
	//
	//fmt.Println("--------------------")
	//
	//const Big = 1 << 100
	//const Small = Big >> 98
	//fmt.Println(Small) // 4

	fmt.Println("------------------------------------------------------------------------------------------------------------------------")

	const Pi = math.Pi

	var r float64
	_, _ = fmt.Scan(&r)

	s := Pi * math.Pow(r, 2)

	fmt.Printf("Circle radius: %.2f → area = %.2f\n", r, s)

	fmt.Println("------------------------------------------------------------------------------------------------------------------------")

}
