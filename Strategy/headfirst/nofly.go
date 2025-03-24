package main

import "fmt"

type Nofly struct {
}

func (n Nofly) fly() {
	fmt.Println("i cant fly")
}