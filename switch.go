package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var str string
	fmt.Print("Enter a binary string: ")
	fmt.Scanln(&str)
	switch {
	case ThreeZeroOdd(str):
		fmt.Println("'000 - ODD'")
	case LessThanThreeZeroOdd(str):
		fmt.Println("'>3 0s - ODD'")
	case LessThanFiveOneEven(str):
		fmt.Println("'>5 1s - EVEN'")
	}
}

// ThreeZeroOdd for first case
func ThreeZeroOdd(str string) bool {
	if binary(str)%2 == 0 {
		return false
	}
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '0' {
			count++
		} else {
			count = 0
		}
		if count > 2 {
			return true
		}
	}
	return false
}

// LessThanThreeZeroOdd for second case
func LessThanThreeZeroOdd(str string) bool {
	if binary(str)%2 == 0 {
		return false
	}
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '0' {
			count++
		} else {
			count = 0
		}
		if count > 2 {
			return false
		}
	}
	return true
}

//LessThanFiveOneEven for third case
func LessThanFiveOneEven(str string) bool {
	if binary(str)%2 == 1 {
		return false
	}
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '1' {
			count++
		} else {
			count = 0
		}
		// fmt.Println(count)
		if count > 5 {
			return false
		}
	}
	return true
}

func binary(str string) int {
	var ans float64 = 0
	for i := 0; i < len(str); i++ {
		var ch string = string(str[i])
		var no, err = strconv.Atoi(ch)
		_ = err
		ans += float64(no) * math.Pow(2, float64(len(str)-i-1))
	}
	return int(ans)
}
