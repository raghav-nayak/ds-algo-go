##### Problem statement:
Say that you are a traveler on a 2D grid. You begin in the top-left corner and your goal is to travel to the bottom-right corner. You may only move down or right.
In how many ways can you travel to the goal on a grid with dimensions m * n?
write a function gridTraveler(m, n) that calculates this


##### Complexity
brute force
- time complexity: O(2^m+n)
- space complexity: O(m+n)

memoization
- time complexity: O(mn)
- space complexity: O(m+n)

gridTraveler(m,n) = gridTraveler(m,n)

![[Pasted image 20240826013317.png]]


##### solution
```go
package main

import (
	"fmt"
	"time"
)


func gridTravelerWithoutMemo(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 && n == 1 {
		return 1
	}
	return gridTravelerWithoutMemo(m-1,n)  + gridTravelerWithoutMemo(m, n-1)
}

func gridTravelerWithMemo(m int, n int) int {
	memo_grid_value := make(map[[2]int]int)
	return gridTravelerWithMap(m, n, memo_grid_value)
}

func gridTravelerWithMap(m int, n int, memo_grid_value map[[2]int]int) int {
	if m == 0 || n == 0 {
		return 0
	}

	if m == 1 && n == 1 {
		return 1
	}

	// both [m][n] and [n][m] same
	// gridTravelerWithMap(m,n) = gridTravelerWithMap(n,m)
	if val, ok := memo_grid_value[[2]int{m, n}]; ok {
		return val
	}

	if val, ok := memo_grid_value[[2]int{n, m}]; ok {
		return val
	}

	calculated_val := gridTravelerWithMap(m-1, n, memo_grid_value) + gridTravelerWithMap(m, n-1, memo_grid_value)
	memo_grid_value[[2]int{m,n}]  = calculated_val
	return calculated_val
}

func main() {
	start := time.Now()
	fmt.Println("Without memoization (1,1) = ", gridTravelerWithoutMemo(1, 1), " time taken = ", time.Since(start))
	start = time.Now()
	fmt.Println("With memoization (1,1) = ", gridTravelerWithMemo(1, 1), " time taken = ", time.Since(start))

	fmt.Println()
	start = time.Now()
	fmt.Println("Without memoization (2,3) = ", gridTravelerWithoutMemo(2, 3), " time taken = ", time.Since(start))
	start = time.Now()
	fmt.Println("With memoization (2,3) = ", gridTravelerWithMemo(2, 3), " time taken = ", time.Since(start))
	
	fmt.Println()
	start = time.Now()
	fmt.Println("Without memoization (3,2) = ", gridTravelerWithoutMemo(3, 2), " time taken = ", time.Since(start))
	start = time.Now()
	fmt.Println("With memoization (3,2) = ", gridTravelerWithMemo(3, 2), " time taken = ", time.Since(start))

	fmt.Println()
	start = time.Now()
	fmt.Println("Without memoization (3,3) = ", gridTravelerWithoutMemo(3, 3), " time taken = ", time.Since(start))
	start = time.Now()
	fmt.Println("With memoization (3,3) = ", gridTravelerWithMemo(3, 3), " time taken = ", time.Since(start))

	fmt.Println()
	start = time.Now()
	fmt.Println("Without memoization (18,18) = ", gridTravelerWithoutMemo(18, 18), " time taken = ", time.Since(start))
	start = time.Now()
	fmt.Println("With memoization (18,18) = ", gridTravelerWithMemo(18, 18), " time taken = ", time.Since(start))
}
```

output
```shell
$ go run gridTraveller.go
Without memoization (1,1) =  1  time taken =  125ns
With memoization (1,1) =  1  time taken =  166ns

Without memoization (2,3) =  3  time taken =  167ns
With memoization (2,3) =  3  time taken =  1.625µs

Without memoization (3,2) =  3  time taken =  125ns
With memoization (3,2) =  3  time taken =  417ns

Without memoization (3,3) =  6  time taken =  166ns
With memoization (3,3) =  6  time taken =  875ns

Without memoization (18,18) =  2333606220  time taken =  21.514572167s
With memoization (18,18) =  2333606220  time taken =  34.167µs
```
