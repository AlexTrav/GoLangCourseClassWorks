package main

import "fmt"

type User struct {
	Name string
}

func (u *User) Greet() string {
	return fmt.Sprintf("Hello %s", u.Name)
}

type Dog struct {
	Name string
}

func (d *Dog) Greet() string {
	return fmt.Sprintf("Woof! Hello %s", d.Name)
}

type Admin struct {
	User
	Dog
	banned     bool
	banMessage string
}

func (a *Admin) GetBanMessage() string {
	return a.banMessage
}

func (a *Admin) Ban(message string) {
	a.banned = true
	a.banMessage = fmt.Sprintf("User %s has been banned for %s", a.User.Name, message)
}

func (a *Admin) Greet() string {
	if a.banned {
		return a.banMessage
	}
	return a.User.Greet()
}

type Greeter interface {
	Greet() string
}

func NewAdmin(user User) *Admin {
	return &Admin{user, Dog{"Bobik"}, false, ""}
}

func main() {
	user1 := new(User)
	user1.Name = "Alexey"
	fmt.Println(user1.Greet())

	admin1 := NewAdmin(User{"John"})
	fmt.Println(admin1.Greet())
	admin1.Ban("БААААН!!")
	fmt.Println(admin1.GetBanMessage())
}

/* Вывод из консоли:
Hello Alexey
Hello John
User John has been banned for БААААН!!
*/
