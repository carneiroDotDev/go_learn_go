package main

import "fmt"

const prefix string = "Hello, "

func hello(name string, idiom string) string {
	if name == "" {
		name = "World"
	} else if idiom == "es" {
		return "Hola, " + name + "!\n"
	} else if idiom == "fr" {
		return "Bonjour, " + name + "!\n"
	}
	return prefix + name + "!\n"
}

func main() {
	businessLogic := hello("Luiz", "")
	fmt.Println(businessLogic)
}