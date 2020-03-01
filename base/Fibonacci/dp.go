package main

import "fmt"

func main() {

	//fmt.Println(F(0))
	//fmt.Println(F(1))
	//fmt.Println(F(2))
	fmt.Println("F", F(6))
	fmt.Println("F2", F2(6))
	fmt.Println("F3", F3(6))
}

func F(i int) (f int) {
	if i <= 0 {
		return 0
	}
	mem := make([]int, i+1)
	for k := 0; k < i+1; k++ {
		mem[k] = -1
	}
	//fmt.Println(mem)

	return dp(i, mem)
}

func dp(i int, mem []int) int {
	if mem[i] != -1 {
		//fmt.Println("节省次数", mem, "len", len(mem), "i:", i)
		return mem[i]
	}
	if i <= 2 {
		mem[i] = 1
	} else {
		mem[i] = dp(i-1, mem) + dp(i-2, mem)
	}

	//fmt.Println(mem, "len", len(mem), "i:", i)
	return mem[i]
}

func F2(n int) (f int) {
	if n <= 0 {
		return 0
	}
	a := make([]int, n+1)
	a[0] = 0
	a[1] = 1

	if n <= 1 {
		return a[n]
	}

	for i := 2; i <= n; i++ {
		a[i] = a[i-1] + a[i-2]
	}

	return a[n]

}

func F3(n int) (f int) {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	a_2 := 0
	a_1 := 1
	a_i := 1

	for i := 2; i <= n; i++ {
		a_i = a_1 + a_2
		a_2 = a_1
		a_1 = a_i
	}

	return a_i

}
