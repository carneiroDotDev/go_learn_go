package main

import "fmt"

func main() {
	var age int = 37

	if age >= 18 {
		fmt.Printf("Beer's on me since you are %d years old!\n", age)
	} else {
		fmt.Printf("You are not allowed to drink!")
	}
}
