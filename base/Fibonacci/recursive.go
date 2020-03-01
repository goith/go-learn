package main

import "fmt"

func main() {
	fmt.Println(F(6))
}

func F(i int) (f int) {
	if i <= 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return F(i-1) + F(i-2)
}
