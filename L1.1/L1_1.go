package main

import "fmt"

type Human struct {
    Name string
    Age  int
}

func (h Human) Greet() {
    fmt.Println("Привет, я", h.Name, "и мне", h.Age, "лет")
}

func (h *Human) Birthday() {
    h.Age++
}

type Action struct {
    Human
    Role  string
}

func (a Action) ShowRole() {
    fmt.Println("Моя роль:", a.Role)
}