package main

import "fmt"

// Указатели. Структуры
func main() {
	// userId := "uuid"
	// x := "10"
	//
	// a := &userId
	// fmt.Println(a)
	//
	// a = &x
	// fmt.Println(a)

	var Point *Vertex = new(Vertex)
	Point.X = 100
	Point.Y = 200
	fmt.Println(Point.X)
	fmt.Println(Point.Y)

	// Или

	var Point1 = Vertex{1, 2}

	Move(&Point1)

	fmt.Println(Point1)
}

type Vertex struct {
	X int
	Y int
}

func Move(p *Vertex) {
	p.X = p.X + 1
	p.Y = p.Y + 1
}
