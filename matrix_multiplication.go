package main

import "fmt"

func main() {
	a := [][]int{
		{1, 2},
	}
	b := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(len(a[0]), len(b))
}
