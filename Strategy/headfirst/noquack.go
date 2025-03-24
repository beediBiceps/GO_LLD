package main

import "fmt"

type Noquack struct {
}

func (y Noquack) quack() {
	fmt.Println("no quack lmao")
}