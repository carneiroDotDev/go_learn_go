package main

import "fmt"

func main(){
	num := 50
	switch {
	case num > 40:
		fmt.Print("This is true! \n")
		fallthrough
	case num < 10:
		fmt.Print("This is false, but because of fallthrough, this thing executes! \n")
	}
}