package main

import "fmt"

func main() {

	//fmt.Println(F(0))
	//fmt.Println(F(1))
	//fmt.Println(F(2))
	fmt.Println(F(6))
}

func F(i int) (f int) {
	if i <= 0 {
		return 0
	}
	mem := make([]int, i+1)
	for k := 0; k < i+1; k++ {
		mem[k] = -1
	}
	fmt.Println(mem)

	return dp(i, mem)
}

func dp(i int, mem []int) int {
	if mem[i] != -1 {
		fmt.Println("节省次数", mem, "len", len(mem), "i:", i)
		return mem[i]
	}
	if i <= 2 {
		mem[i] = 1
	} else {
		mem[i] = dp(i-1, mem) + dp(i-2, mem)
	}

	fmt.Println(mem, "len", len(mem), "i:", i)
	return mem[i]
}
