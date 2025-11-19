package main // Package lesson_2 Типы данных
import (
	"fmt"
	"math"
)

func main() {
	// Integer
	// fmt.Println("Integer:")
	// var a1 int8 = 127
	// var a2 int16 = 12588
	// var a3 int32 = 156464765
	// var a4 int64 = -4444798464654165
	// fmt.Println(a1)
	// fmt.Println(a2)
	// fmt.Println(a3)
	// fmt.Println(a4)

	// Unsigned Integer - без отриц чисел
	// fmt.Println("Unsigned Integer:")
	// var b1 uint8 = 127
	// var b2 uint16 = 12588
	// var b3 uint32 = 156464765
	// var b4 uint64 = 4444798464654165
	// fmt.Println(b1)
	// fmt.Println(b2)
	// fmt.Println(b3)
	// fmt.Println(b4)

	// fmt.Println("Integer + Unsigned Integer:", int32(a1)+int32(b1))

	// Float - Default Float64
	// var c1 float32 = 127.8464123123
	// var c2 float64 = 12588.8648812312321312
	// fmt.Println(c1)
	// fmt.Println(c2)

	// Complex
	// var d1 complex64 = complex(1, 2)
	// var d2 complex128 = complex(1, 2)
	// g := 0.867 + 0.5i // complex128
	// fmt.Println(d1)
	// fmt.Println(d2)
	// fmt.Println(g)

	// Boolean
	// var isActive bool = true
	// var isCompleted bool = false
	// fmt.Println(isActive)
	// fmt.Println(isCompleted)

	// String
	// var str1 string = "Hello,"
	// var str2 string = "World!"
	// fmt.Println(str1, str2)
	// fmt.Println(string(str1[0]))
	// fmt.Println(string(str2[0]))

	// Byte
	// var bt1 byte = 0
	// var bt2 byte = 1
	// var bt3 byte = 256 // от 0 до 255 - ОШИБКА
	// fmt.Println(bt1)
	// fmt.Println(bt2)
	// fmt.Println(bt3)

	// Rune
	// var rn rune = 'Ү'
	// fmt.Println(string(rn))

	// Default variables
	// var i int
	// var f float32
	// var b bool
	// var s string
	// var r rune
	// var bt byte

	// fmt.Println(i, f, b, s, r, bt)

	// Convert
	// var x int = 100
	// var y float64 = float64(x) + 3.14
	// var z int = int(y)
	// fmt.Println(z)

	// numStr := "123"
	// num, _ := strconv.Atoi(numStr)
	// str := strconv.Itoa(num)
	// fmt.Println(num)
	// fmt.Println(str)
	//
	// flStr := "3.14"
	// floatNum, _ := strconv.ParseFloat(flStr, 64)
	// fmt.Println(floatNum)

	// const pi float64 = 3.1415926
	// fmt.Println(pi)

	// Циклы
	// for со счётчиком
	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	// for с условием
	// sum := 1
	// for sum < 1000 {
	// 	sum += sum
	// }
	// fmt.Println(sum)

	// или

	// sum := 1
	// for sum < 1000 {
	// 	sum += sum
	// }
	// fmt.Println(sum)

	// For без условий (бесконечный)
	// sum := 2
	// for {
	// 	sum += sum // удваиваем каждый раз
	// 	fmt.Println(sum)
	// }

	// Условия - IF - ELSE - ELSEIF
	// fmt.Println(sqrt(2), sqrt(-4))
	//
	// fmt.Println(
	// 	pow1(3, 2, 10),
	// 	pow1(3, 3, 20),
	// )
	//
	// fmt.Println(
	// 	pow2(3, 2, 10),
	// 	pow2(3, 3, 20),
	// )

	// SWITCH
	result := true
	switch result {
	case true:
		fmt.Println("yes")
	case false:
		fmt.Println("no")
	case nil:
		fmt.Println("nil")
	default:
		fmt.Println("maybe")
	}
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow1(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { // Присваивание внутри условия
		return v
	}
	return lim
}

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}
