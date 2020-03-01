package main

import (
	"flag"
	"fmt"
	"log"
)

//题目：假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

//注意：给定 n 是一个正整数。

//分析
//定义：f( n)表示n个台阶的爬台阶的方法数
//选择：上一步爬了一步还是爬了两步
//初始化：f(0)=1,f(1)=1
//递推表达：f(n)=f(n-1)+f(n-2),n>=2
//注：该问题类似于求解斐波那契数列（Fibonacci sequence）。斐波那契数列从0开始：0,1,1,2,3,5...，此问题解为1,1,2,3,5,....
func main() {
	var n int
	flag.IntVar(&n, "n", 2, "stairs step")
	flag.Parse()
	fmt.Println("Input val:", n)
	c := climbStairs(n)
	fmt.Printf("一共%d种方法\n", c)
}

func climbStairs(n int) int {
	if n < 0 {
		log.Fatal(fmt.Errorf("param err."))
	}

	a := make([]int, n+1)
	a[0] = 1
	a[1] = 1
	if n == 1 || n == 0 {
		return a[n]
	}
	for i := 2; i <= n; i++ {
		a[i] = a[i-1] + a[i-2]
	}
	return a[n]
}
