Problem statement:
Write a function canSum(targetSum, numbers) that takes in a targetSum and an array of numbers as arguments.
The function should return a boolean indicating whether or not it is possible to generate the targetSum using numbers from the array.
You may use an element of the array as many times as needed.
You may assume that all input numbers are nonnegative.

example:
canSum(7, [5, 3, 4, 7]) -> true
3 + 4
7

canSum(7, [2,4]) -> false


![[Pasted image 20240827005718.png]]

![[Pasted image 20240827005859.png]]



m = target sum
n = array length

with brute force
![[Pasted image 20240827011130.png]]

![[Pasted image 20240827011240.png]]


**Solution:**

```go
/*

Write a function canSum(targetSum, numbers) that takes in a targetSum and an array of numbers as arguments.
The function should return a boolean indicating whether or not it is possible to generate the targetSum using numbers from the array.
You may use an element of the array as many times as needed.
You may assume that all input numbers are nonnegative.

example:
canSum(7, [5, 3, 4, 7]) -> true
3 + 4
7

canSum(7, [2,4]) -> false
*/

package main

import "fmt"

func canSumWithoutMemo(targetSum int, numbers []int) bool {
	if targetSum == 0 {
		return true
	}

	if targetSum < 0 {
		return false
	}

	for i := range numbers {
		remainder := targetSum - numbers[i]
		
		if canSumWithoutMemo(remainder, numbers) {
			return true
		}
	}
	return false
}

func canSumWithMemo(targetSum int, numbers []int, memo map[int]bool) bool {
	if _, ok := memo[targetSum]; ok {
			return memo[targetSum]
	}

	if targetSum == 0 {
		return true
	}

	if targetSum < 0 {
		return false
	}

	for i := range numbers {
		reminder := targetSum - numbers[i]
		if canSumWithMemo(reminder, numbers, memo) {
			memo[reminder] = true
			return true
		}
	}
	memo[targetSum] = false
	return false
}

func main() {

	fmt.Println(canSumWithoutMemo(7, []int{5, 3, 4, 7})) // true
	fmt.Println(canSumWithoutMemo(7, []int{2, 4})) // false
	fmt.Println(canSumWithoutMemo(7, []int{2, 3})) // true
	fmt.Println(canSumWithoutMemo(7, []int{2, 3, 5})) // true
	fmt.Println(canSumWithoutMemo(8, []int{2})) // true
	fmt.Println(canSumWithoutMemo(300, []int{7, 14})) // false

	fmt.Println(canSumWithMemo(7, []int{5, 3, 4, 7}, map[int]bool{})) // true
	fmt.Println(canSumWithMemo(7, []int{2, 4}, map[int]bool{})) // false
	fmt.Println(canSumWithMemo(7, []int{2, 3}, map[int]bool{})) // true
	fmt.Println(canSumWithMemo(7, []int{2, 3, 5}, map[int]bool{})) // true
	fmt.Println(canSumWithMemo(8, []int{2}, map[int]bool{})) // true
	fmt.Println(canSumWithMemo(300, []int{7, 14}, map[int]bool{})) // false
}
```
