### Problem statement
Write a function `canConstruct(target, wordBank)` that accepts a target string and an array of strings.
The function should return a boolean indicating whether or not the `target` can be constructed by concatenating elements of the `wordBank` array.
You may reuse elements of `wordBank` as many times as needed.

example:
`canConstruct("abcdef", ["ab", "abc", "cd", "def", "abcd"])` -> true

`canConstruct("skateboard", ["bo", "rd", "ate", "t", "ska", "sk", "boar"])` -> false
ska + t + ?
sk + ate + boar + ?
sk + ate _ bo + ?

`canConstruct("", ["cat", "dog", "mouse"])` -> true
empty string generation takes 0 elements from the array


![[Pasted image 20240905230927.png]]
don't take anything from middle



![[Pasted image 20240905231028.png]]


![[Pasted image 20240905231121.png]]


![[Pasted image 20240905231143.png]]



![[Pasted image 20240905231151.png]]



![[Pasted image 20240905231159.png]]




another example

![[Pasted image 20240905231347.png]]


