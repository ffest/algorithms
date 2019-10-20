class Solution:
    def myPow(self, x: float, n: int) -> float:
        result = 1.0
        if n < 0:
            x, n = 1.0 / x, -n
        while n:
            if n & 1:
                result *= x
            x, n = x * x, n >> 1
        return result


solution = Solution()
print(solution.myPow(2, -2))
