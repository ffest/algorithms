class Solution:
    def reverse(self, x: int) -> int:
        sign = -1 if x < 0 else 1
        res, x = 0, abs(x)
        while x:
            res = res * 10 + x % 10
            x //= 10
        if res > 1 << 31 - 1:
            return 0
        return sign * res


solution = Solution()
print(solution.reverse(-321))
