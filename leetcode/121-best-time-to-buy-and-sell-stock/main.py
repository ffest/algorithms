from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        if len(prices) < 2:
            return 0

        max_profit = 0
        min_price = prices[0]
        for i in range(1, len(prices)):
            max_profit = max(max_profit, prices[i]-min_price)
            min_price = min(min_price, prices[i])
        return max_profit


solution = Solution()
print(solution.maxProfit([7, 1, 5, 3, 6, 4]))
