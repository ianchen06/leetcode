/*
Given an array of integers temperatures represents the daily temperatures,
return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature.
If there is no future day for which this is possible, keep answer[i] == 0 instead.

Example 1:

Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]
Example 2:

Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]
Example 3:

Input: temperatures = [30,60,90]
Output: [1,1,0]


Constraints:

1 <= temperatures.length <= 105
30 <= temperatures[i] <= 100
*/
package main

func solution(t []int) []int {
	stack := make([]int, len(t))
	top := -1 //directly add first right element to the stack
	ans := make([]int, len(t))

	for i := 0; i < len(t); i++ {
		for top > -1 && t[i] > t[stack[top]] {
			idx := stack[top]
			top--
			ans[idx] = i - idx
		}
		top++
		stack[top] = i
	}
	return ans
}
