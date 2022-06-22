/*
76. Minimum Window Substring
https://leetcode.com/problems/minimum-window-substring/

Given two strings s and t of lengths m and n respectively, return the minimum window substring of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".

The testcases will be generated such that the answer is unique.

A substring is a contiguous sequence of characters within the string.



Example 1:

Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
Example 2:

Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.
Example 3:

Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.


Constraints:

m == s.length
n == t.length
1 <= m, n <= 105
s and t consist of uppercase and lowercase English letters.


Follow up: Could you find an algorithm that runs in O(m + n) time?
*/

package main

import (
	"math"
)

// use map to count number of chars, since we can have duplicates in t
func solution(s string, t string) string {
	// length, start, end
	ans := []int{math.MaxInt, 0, 0}

	// build need map for reference
	need := make(map[string]int)
	for _, b := range t {
		need[string(b)] += 1
	}
	// num of unique chars
	tl := len(need)

	// current
	db := make(map[string]int)
	// counter is seen unique chars
	start, end, counter := 0, 0, 0

	for end < len(s) {
		db[string(s[end])]++
		_, ok := need[string(s[end])]
		if ok && db[string(s[end])] == need[string(s[end])] {
			counter++
		}
		// if start does not pass end, the substring is valid
		for start <= end && counter == tl {
			// update the answer if we found a shorter one
			if end-start+1 < ans[0] {
				ans = []int{end - start + 1, start, end}
			}
			_, ok := need[string(s[start])]
			if ok && db[string(s[start])] == need[string(s[start])] {
				counter--
			}
			db[string(s[start])]--
			start++
		}
		end++
	}
	if ans[0] == math.MaxInt {
		return ""
	}
	return s[ans[1] : ans[2]+1]
}
