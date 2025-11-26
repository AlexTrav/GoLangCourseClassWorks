package main

import "fmt"

type Greeter interface {
	Greet() string
}

type User struct {
	Name string
}

func (u User) Greet() string {
	return fmt.Sprintf("Hello, %s", u.Name)
}

type Robot struct {
	Model string
}

func (r Robot) Greet() string {
	return "Hello, I am a " + r.Model
}

func SayHello(g Greeter) string {
	return g.Greet()
}

type Worker interface {
	Work() string
}

type WorkerImpl struct {
	WorkDone bool
}

func (w WorkerImpl) Work() string {
	w.WorkDone = true
	return "Doing work"
}

func NewWorker() Worker {
	var w *WorkerImpl = nil
	w = &WorkerImpl{}
	return w
}

func PrintGreeter(g Greeter) {
	u, ok := g.(User)
	if ok {
		fmt.Println(u.Greet())
	}
}

func PrintType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("Это int")
	case string:
		fmt.Println("Это string")
	case bool:
		fmt.Println("Это bool")
	case Greeter:
		fmt.Println("Это Greeter")
	case Worker:
		fmt.Println("Это worker")
	default:
		fmt.Println("Неизвестный тип")
	}
}

type V interface {
	Val()
}

type P interface {
	Ptr()
}

type MyType struct {
	Name string
}

func (MyType) Val() {}

func (*MyType) Ptr() {}

type Saver interface {
	Save() error
	IsSaved() bool
}

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Saved    bool
}

func (c Config) Save() error {
	fmt.Println(fmt.Sprintf("%s:%s@%s:%d", c.User, c.Password, c.Host, c.Port))
	c.Saved = true
	return nil
}

func (c Config) IsSaved() bool {
	return c.Saved
}

func main() {

	andrey := User{"Andrey"}
	fmt.Println(SayHello(andrey))

	pilesos := Robot{"Pilesos001"}
	fmt.Println(SayHello(pilesos))

	var SomeGreeter Greeter

	SomeGreeter = User{"Alex"}
	fmt.Println(SayHello(SomeGreeter))

	SomeGreeter = Robot{"Pilesos002"}
	fmt.Println(SayHello(SomeGreeter))

	var w Worker

	w = NewWorker()
	_ = w

	if w != nil {
		fmt.Println(w.Work())
	}

	var someValue interface{} = "some str"

	v, ok := someValue.(int)
	if ok {
		fmt.Println("Это integer", v)
	}

	s, ok := someValue.(string)
	if ok {
		fmt.Println("Это string", s)
	}

	// var u *User = nil
	var u = User{"Alex"}
	PrintGreeter(u)

	PrintType(u)
	PrintType(123)

	var _ V = &MyType{}
	var _ V = MyType{}

	var _ P = &MyType{} // Исправлено c var _ P = MyType
	var _ P = &MyType{}

	var dbConfig Saver = &Config{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		Password: "root123456",
	}

	err := dbConfig.Save()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(dbConfig.IsSaved())
}
