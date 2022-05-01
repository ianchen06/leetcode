"""
647. Palindromic Substrings
https://leetcode.com/problems/palindromic-substrings/

Given a string s, return the number of palindromic substrings in it.

A string is a palindrome when it reads the same backward as forward.

A substring is a contiguous sequence of characters within the string.

Example 1:

Input: s = "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".
Example 2:

Input: s = "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
"""
import unittest


class Solution:
    """
    iterate through string,
    use each char as center,
    two pointers, then expand right and left from center til out-of-bounds
    """

    def countSubstrings(self, s: str) -> int:
        res = 0
        n = len(s)
        for a in range(n):
            i, j = a, a
            while 0 <= i < n and 0 <= j < n and s[i] == s[j]:
                res += 1
                i -= 1
                j += 1
            i, j = a, a + 1
            while 0 <= i < n and 0 <= j < n and s[i] == s[j]:
                res += 1
                i -= 1
                j += 1
        return res


class TestSolution(unittest.TestCase):
    def test_countSubstrings(self):
        test_cases = [["abc", 3], ["aaa", 6], ["abcbb", 7]]
        for test_case, expected in test_cases:
            s = Solution()
            res = s.countSubstrings(test_case)
            self.assertEqual(res, expected)


if __name__ == "__main__":
    unittest.main()
