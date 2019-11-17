from typing import List


class Solution:
    def nextPermutation(self, nums: List[int]) -> None:
        if len(nums) < 2:
            return

        idx = len(nums) - 1
        while idx > 0:
            if nums[idx - 1] < nums[idx]:
                break
            idx -= 1
        # Need to reverse sort in any case
        self.reverseSort(nums, idx, len(nums)-1)
        for i in range(idx, len(nums)):
            if nums[i] > nums[idx-1]:
                nums[idx-1], nums[i] = nums[i], nums[idx-1]
                return

    def reverseSort(self, nums: List[int], start: int, end: int) -> None:
        while start < end:
            nums[start], nums[end] = nums[end], nums[start]
            start += 1
            end -= 1


solution = Solution()
nums = [1, 1, 5]
solution.nextPermutation(nums)
print(nums)
