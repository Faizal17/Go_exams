package main

import "fmt"

func main() {
	var no int
	fmt.Println("Enter number of disc to shift from A to C: ")
	fmt.Scanln(&no)
	hanoiTower(no, "A", "B", "C")
}

func hanoiTower(n int, from string, intermediate string, to string) {
	// fmt.Println(n, from, intermediate, to)
	if n > 0 {
		hanoiTower(n-1, from, to, intermediate)
		fmt.Println("Move disk from ", from, "intermediate rod", to)
		hanoiTower(n-1, intermediate, from, to)
	}
}
