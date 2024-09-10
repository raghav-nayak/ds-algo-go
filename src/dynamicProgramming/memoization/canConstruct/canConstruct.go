/*
Write a function `canConstruct(target, wordBank)` that accepts a target string and an array of strings.
The function should return a boolean indicating whether or not the `target` can be constructed by concatenating elements of the `wordBank` array.
You may reuse elements of `wordBank` as many times as needed.

example:
canConstruct("abcdef", ["ab", "abc", "cd", "def", "abcd"]) -> true

canConstruct("skateboard", ["bo", "rd", "ate", "t", "ska", "sk", "boar"]) -> false
ska + t + ?
sk + ate + boar + ?
sk + ate _ bo + ?


canConstruct("", ["cat", "dog", "mouse"]) -> true
empty string generation takes 0 elements from the array

*/

package main

import (
	"fmt"
	"strings"
)

func canConstructWithoutMemo(target string, wordBank []string) bool {
	if target == "" {
		return true
	}

	for word := range wordBank {
		index := strings.Index(target, wordBank[word])
		if index == 0 {
			suffix := target[len(wordBank[word]):]
			if canConstructWithoutMemo(suffix, wordBank) {
				return true
			}	
		}
	}

	return false
}

func canConstructWithMemo(target string, wordBank []string, memo map[string]bool) bool {
	if _, ok := memo[target]; ok {
			return memo[target]
	}
	if target == "" {
		return true
	}

	for word := range wordBank {
		index := strings.Index(target, wordBank[word])
		if index == 0 {
			suffix := target[len(wordBank[word]):]
			if canConstructWithMemo(suffix, wordBank, memo) {
				memo[target] = true
				return true
			}	
		}
	}

	memo[target] = false
	return false
}


func main() {
	fmt.Println(canConstructWithoutMemo("abcdef", []string{"ab", "abc", "cd", "def", "abcd"})) // true
	fmt.Println(canConstructWithoutMemo("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"})) // false
	fmt.Println(canConstructWithoutMemo("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"})) // true
	// fmt.Println(canConstructWithoutMemo("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"})) // false


	fmt.Println(canConstructWithMemo("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}, map[string]bool{})) // true
	fmt.Println(canConstructWithMemo("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}, map[string]bool{})) // false
	fmt.Println(canConstructWithMemo("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}, map[string]bool{})) // true
	fmt.Println(canConstructWithMemo("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"}, map[string]bool{})) // false
}