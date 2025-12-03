package main

import "fmt"

type IReader interface {
	Read() string
}

type IWriter interface {
	Write(data string)
}

type IReaderWriter interface {
	IReader
	IWriter
}

type MemoryBuffer struct {
	data string
}

func (m *MemoryBuffer) Read() string {
	return m.data
}

func (m *MemoryBuffer) Write(data string) {
	m.data = data
}

func main() {
	var iReader IReader = &MemoryBuffer{}
	_ = iReader
	fmt.Println("iReader:", iReader.Read())

	var iWriter IWriter = &MemoryBuffer{}
	_ = iWriter
	iWriter.Write("Test data from iWriter")

	var iReaderWriter IReaderWriter = &MemoryBuffer{}
	_ = iReaderWriter
	iReaderWriter.Write("Test data from iReaderWriter")
	fmt.Println("iReaderWriter:", iReaderWriter.Read())
}

/* Вывод из консоли:
iReader:
iReaderWriter: Test data from iReaderWriter
*/
