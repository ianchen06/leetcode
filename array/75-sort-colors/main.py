"""
75. Sort Colors
https://leetcode.com/problems/sort-colors/

Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function.

Example 1:

Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
Example 2:

Input: nums = [2,0,1]
Output: [0,1,2]

Constraints:

n == nums.length
1 <= n <= 300
nums[i] is either 0, 1, or 2.

Follow up: Could you come up with a one-pass algorithm using only constant extra space?
"""
import unittest
from typing import List


class Solution:
    """
    https://en.wikipedia.org/wiki/Dutch_national_flag_problem
    """

    def sortColors(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        r, w, b = 0, 0, len(nums) - 1
        while w <= b:
            if nums[w] == 0:
                nums[w], nums[r] = nums[r], nums[w]
                w += 1
                r += 1
            elif nums[w] == 1:
                w += 1
            else:
                nums[w], nums[b] = nums[b], nums[w]
                b -= 1


class TestSolution(unittest.TestCase):
    def test_sortColors(self):
        test_cases = [[[2, 0, 2, 1, 1, 0], [0, 0, 1, 1, 2, 2]], [[2, 0, 1], [0, 1, 2]]]
        for test_case, expected in test_cases:
            s = Solution()
            s.sortColors(test_case)
            self.assertEqual(test_case, expected)


if __name__ == "__main__":
    unittest.main()
