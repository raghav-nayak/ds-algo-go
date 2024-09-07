
[article link](https://hackernoon.com/14-patterns-to-ace-any-coding-interview-question-c5bb3357f6ed)


#### Fundamentals
To find the greatest amount of success when practicing, it is highly recommended to know the methods and runtimes of the following data structures and their operations:

- Arrays
- Maps
- Linked Lists
- Queues
- Heaps
- Stacks
- Trees
- Graphs

In addition, you should have a good grasp on common algorithms such as:

- Breadth-first search
- Depth-first search
- Binary search
- Recursion



## Patterns
### 1. Sliding window
- Specific window size of a given array / linked list / string
#### questions
1. Max sum subarray of K (E)
2. Longest subarray with k distinct char(M)
3. String anagram (H)

### 2. Two pointers or iterators
- drawbacks of single operator
	- inefficient
	- asymptotic analysis

- useful in sorted array / linked list
- comparing array elements with other elements
- set of elements from array -> pair / triplet / subarray
#### questions
1. squaring a sorted array (E)
2. triplet that runs to zero (M)
3. comparing strings that contains backspaces (M)

### 3. Fast and flow pointers
- also known as Hare and Tortoise algorithm
- uses two pointers that move through the array / sequence / linked list at different speeds.
- useful when dealing with cyclic(loop) linked lists or arrays -> two pointers will bound to meet
- used when you need to know the position of a certain element or the overall length of the linked list
- useful when linked list is a palindrome number

Not used in 
- in single linked list where you can't move in a backwards direction
#### questions
1. Linked list cycle (E)
2. Palindrome Linked List (M)
3. Cycle in a circular array (H)

### 4. Merge intervals
- to deal with overlapping intervals or merge intervals if they overlap
- 6 case ( a and b )
- used to produce a list with only mutually exclusive intervals
#### questions
1. Intervals intersection (M)
2. Maximum CPU load (H)

### 5. Cyclic sort
- deal with problems involving arrays containing numbers in a given range
- iterates over the array one number at a time, and if the current number you are iterating is not a t the correct index, you swap it with the number at its correct index
- results in O(n^2) complexity

- useful with problems involving a sorted array with numbers in a given range
- problems to find the missing / duplicate / smallest number in an sorted  / rotated array
#### questions
1. Find the missing number (E)
2. Find the smallest missing positive number (M)

### 6. In-place reversal of linked list

- using the existing node objects and without using extra memory
- lock-step manner
#### questions
1. Reverse a sub-list (M)
2. Reverse a K-element sub-list (M)

### 7. Tree BFS
- Breadth-first tree
- uses a queue to keep track of all the nodes of a level before go to next level
#### questions
1. Binary tree level order traversal (E)
2. Zigzag traversal (M)

### 8. Tree DFS
- Depth-first search
- recursion or stack for iterative approach to keep track of parent nodes while traversing
- 3 types
	- pre-order (P-LC-RC)
	- in-order (LC-P-RC)
	- post-order (LC-RC-P)
#### questions
1. sum of path numbers (M)
2. all paths of a sum (M)
3. searching for something where the node is closer to a leaf

### 9. Two heaps
- Min and max heaps
- useful in problems featuring a BST
#### questions
1. useful in situations like priority queue, scheduling
2. find the smallest / largest / median elements of a set
3. find the median of a number stream (M)

### 10. Subsets
- permutations and combinations of a given set of elements
- this pattern describes an efficient BFS approach
#### questions:
1. subsets with duplicates (E)
2. string permutations by changing case (M)

### Modified binary search
### Top K elements
### K-way merge
### Topological sort