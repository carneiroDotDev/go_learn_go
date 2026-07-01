package main

import "fmt"

func Hello(name string, idiom string) string {
	if name == "" {
		name = "World"
	}

	return getPrefix(idiom) + name + "!\n"
}

func getPrefix(idiom string) (prefix string) {
	switch idiom {
	case "es":
		prefix = "Hola, "
	case "fr":
		prefix = "Bonjour, "
	default:
		prefix = "Hello, "
	}

	return
}

func main() {
	businessLogic := Hello("Luiz", "fr")
	fmt.Println(businessLogic)
}