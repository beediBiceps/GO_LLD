package main

import "fmt"

func main(){

    pizza:=VeggieMania{}

    pizzaWithCheese:=&CheeseTopping{
        pizza: &pizza,
    }

    pizzaWithCheeseAndTomato:=&TomatoTopping{
        pizza: pizzaWithCheese,
    }

    fmt.Printf("Price of veggieMania with cheese and tomato topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}