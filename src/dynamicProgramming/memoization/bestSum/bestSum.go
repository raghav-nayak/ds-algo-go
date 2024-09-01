/*

Write a function `bestSum(targetSum, numbers)` that takes in a targetSum and an array of numbers as arguments.
The function should return an array containing the shortest combination of numbers that add up to exactly the targetSum.
If there is a tie for the shortest combination, you may return any one of the shortest.

e.g.
bestSum(7, [5, 3, 4, 7]) -> [7]
bestSum(8, [2, 3, 5]) -> [3, 5]
bestSum(8, [1, 4, 5]) -> [4, 4]
bestSum(100, [1, 2, 5, 25]) -> [25, 25, 25, 25]
*/

package main

import "fmt"

func bestSumWithoutMemo(targetSum int, numbers []int) []int {
	if targetSum == 0 {
		return []int{}
	}

	if targetSum < 0 {
		return nil
	}

	var shortestCombination []int

	for _, num := range numbers{
		remainder := targetSum - num
		remainderCombination := bestSumWithoutMemo(remainder, numbers)
		if remainderCombination != nil {
			combination := append([]int{}, remainderCombination...)
            combination = append(combination, num)
			if (shortestCombination == nil || len(combination) < len(shortestCombination)) {
				shortestCombination = combination
			}
		}
	}

	return shortestCombination
}

func main() { 

	fmt.Println(bestSumWithoutMemo(7, []int{5, 3, 4, 7}))
	fmt.Println(bestSumWithoutMemo(8, []int{2, 3, 5}))
	fmt.Println(bestSumWithoutMemo(8, []int{1, 4, 5}))
}