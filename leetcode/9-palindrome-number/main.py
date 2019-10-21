import math


class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x <= 0:
            return x == 0

        digits = math.floor(math.log10(x)) + 1
        mask = 10 ** (digits - 1)
        for i in range(digits // 2):
            if x % 10 != x // mask:
                return False
            x %= mask
            x //= 10
            mask //= 100

        return True


solution = Solution()
print(solution.isPalindrome(12322))
