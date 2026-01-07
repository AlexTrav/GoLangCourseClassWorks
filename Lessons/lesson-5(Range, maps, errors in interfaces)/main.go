package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

type IRepository interface {
	Find(id int) (*User, *UserNotFoundError)
}

type Repository struct {
	Data []*User
}

type UserNotFoundError struct {
	cause error
	Code  int
	msg   string
}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("%s: %s", e.msg, e.cause.Error())
}

func (e *UserNotFoundError) Unwrap() error {
	return e.cause
}

func (r *Repository) Find(id int) (*User, *UserNotFoundError) {
	if id < 0 || id >= len(r.Data) {
		return nil, &UserNotFoundError{errors.New("not found"), 404, fmt.Sprintf("no such user with id: %d", id)}
	} else {
		return r.Data[id], nil
	}
}

func (r *Repository) FindAll() []*User {
	return r.Data
}

func main() {
	a := []string{"one", "two"}
	u := []User{{"one", 1}, {"two", 2}}

	for i, v := range a {
		println(i, v)
	}

	for i, _ := range u {
		if u[i].Name == "two" {
			u[i].Age = 30
		}
	}

	for _, v := range u {
		fmt.Printf("пользователь %s возраст %d \n", v.Name, v.Age)
	}

	for i := 0; i < len(u); i++ {
		fmt.Printf("for user %s \n", u[i].Name)
	}

	userRepo := Repository{Data: make([]*User, 0)}

	for i := 0; i < 10; i++ {
		userRepo.Data = append(userRepo.Data, &User{fmt.Sprintf("user %d", i), i + 30})
	}

	user, err := userRepo.Find(11)
	if err != nil {
		if err.Code == 404 {
			fmt.Println("fire event | metric")
		}
		fmt.Printf("Пользователь не найден: %s\n", err.Error())
	} else {
		fmt.Printf("найден пользователь %s, возраст %d \n", user.Name, user.Age)
	}

	fmt.Println("список пользователей ------")
	for i, u := range userRepo.FindAll() {
		fmt.Printf("user %d: name %s age: %d \n", i, u.Name, u.Age)
	}

}
