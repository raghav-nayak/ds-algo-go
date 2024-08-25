package main

import "fmt"

func memoizedFibonacciWithMap(n int) int {
	memo_fib := make(map[int]int)
	return fibonacciWithMap(n, memo_fib)
}


func fibonacciWithMap(n int, memo_fib map[int]int) int {
	if n <= 1 {
		return n
	}

	if val, ok := memo_fib[n]; ok {
		return val
	}
	result := fibonacciWithMap(n-1, memo_fib) + fibonacciWithMap(n-2, memo_fib)
	memo_fib[n] = result
	return result
}


func memoizedFibonacciWithSlice(n int) int {
	memo_fib := make([]int, n+1)
	return fibonacciWithSlice(n, memo_fib)
}

func fibonacciWithSlice(n int, memo_fib []int) int {
	if n <= 1 {
		return n
	}
	if memo_fib[n] != 0 {
		return memo_fib[n]
	}
	result := fibonacciWithSlice(n-1, memo_fib) + fibonacciWithSlice(n-2, memo_fib)
	memo_fib[n] = result
	return result
}

func main() {
	fmt.Println(memoizedFibonacciWithMap(100))
	fmt.Println(memoizedFibonacciWithSlice(100))
}