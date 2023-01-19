package main

import (
	"fmt"
)

// declare function type
type FoodNameGetter func(string) string

type Food struct {
	name   string
	getter FoodNameGetter // declare function
}

func main() {
	pizza := Food{
		name: "Pizza",
		getter: func(name string) string { // declare function body
			return "This is " + name + "."
		},
	}
	pasta := Food{
		name: "Pasta",
		getter: func(name string) string { // declare function body
			return "This is Stinki Furz."
		},
	}
	fmt.Println(pizza.getter(pizza.name)) // prints "This is Pizza."
	fmt.Println(pasta.getter(pasta.name))
}
