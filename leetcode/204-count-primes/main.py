from typing import List


class Solution:
    def countPrimes(self, n: int) -> int:
        if n < 2:
            return 0
        primes = 0
        is_prime = [False, False] + [True] * (n - 2)
        for p in range(2, n):
            if not is_prime[p]:
                continue
            primes += 1
            for i in range(p * 2, n, p):
                is_prime[i] = False

        return primes


solution = Solution()
print(solution.countPrimes(10))
