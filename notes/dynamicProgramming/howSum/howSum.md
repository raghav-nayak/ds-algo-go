##### Problem statement
Write a function howSum(targetSum, numbers) that takes in a targetSum and an array of numbers as arguments.
The function should return an array containing any combination of elements that add up to exactly the targetSum. If there is no combination that adds up to the targetSum, then return null.
If there are multiple combinations possible, you may return any single one.

example:
howSum(7, [5, 3, 4, 7]) 
[7]
[3,4]

howSum(8, [2, 3, 5])
[2, 2, 2, 2]
[3,5]

howSum(7, [2, 4])
null

howSum(0, [1,2,3])
[]



![[Pasted image 20240829004333.png]]


![[Pasted image 20240829004357.png]]

![[Pasted image 20240829004449.png]]


![[Pasted image 20240829004606.png]]


#### complexity
time: O(n^m * m)
space: O(m)



#### Solution:
```go
package main

import "fmt"

/*
func howSum(targetSum int, numbers []int) []int {
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}
	for _, num := range numbers {
		remainder := targetSum - num;
		remainderResult := howSum(remainder, numbers)
		if remainderResult != nil {
			remainderResult = append(remainderResult, num)
			return remainderResult
		}
	}
	return nil
}
*/

func howSumWithoutMemo(targetSum int, numbers []int) []int {
    if targetSum == 0 {
        return []int{}
    }
    if targetSum < 0 {
        return nil
    }
    for _, num := range numbers {
        remainder := targetSum - num
        if remainderResult := howSumWithoutMemo(remainder, numbers); remainderResult != nil {
            remainderResult = append(remainderResult, num)
            return remainderResult
        }
    }
    return nil
}

func howSumMemo(targetSum int, numbers []int, memo map[int][]int) []int {
    if _, ok := memo[targetSum]; ok {
        return memo[targetSum]
    }
    if targetSum == 0 {
        return []int{}
    }
    if targetSum < 0 {
        return nil
    }
    for _, num := range numbers {
        remainder := targetSum - num
        if remainderResult := howSumMemo(remainder, numbers, memo); remainderResult != nil {
            remainderResult = append(remainderResult, num)
            memo[targetSum] = remainderResult
            return remainderResult
        }
    }
    memo[targetSum] = []int{} // Set to an empty slice to indicate no combination
    return nil
}

func main() {
	fmt.Println(howSumWithoutMemo(7, []int{5, 3, 4, 7}))
	fmt.Println(howSumWithoutMemo(8, []int{2, 3, 5}))
	fmt.Println(howSumWithoutMemo(7, []int{2, 4}))
	fmt.Println(howSumWithoutMemo(8, []int{2, 3, 5}))
	fmt.Println(howSumWithoutMemo(300, []int{7, 14}))

	fmt.Println(howSumMemo(7, []int{5, 3, 4, 7}, map[int][]int{}))
	fmt.Println(howSumMemo(8, []int{2, 3, 5}, map[int][]int{}))
	fmt.Println(howSumMemo(7, []int{2, 4}, map[int][]int{}))
	fmt.Println(howSumMemo(8, []int{2, 3, 5}, map[int][]int{}))
	fmt.Println(howSumMemo(300, []int{7, 14}, map[int][]int{}))
}
```
