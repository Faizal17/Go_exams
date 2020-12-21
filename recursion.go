package main

import "fmt"

func main() {
	// fmt.Println("")
	var no int
	fmt.Print("Enter any number to find it's factorial : ")
	fmt.Scanln(&no)
	if no < 0 {
		fmt.Println("Please enter valid input")
	} else {
		ans := fact(no)
		fmt.Println("Factorial of ", no, " is", ans)
	}
	// fmt.Println(no)
}

func fact(no int) int {
	defer fmt.Println("Factorial of ", no, "is completed")
	if no == 0 || no == 1 {
		return 1
	}
	return no * fact(no-1)
}
