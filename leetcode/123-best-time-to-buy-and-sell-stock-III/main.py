from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        first_sell, second_sell = 0, 0
        first_buy, second_buy = -float('inf'), -float('inf')

        for price in prices:
            second_sell = max(second_sell, second_buy+price)
            second_buy = max(second_buy, first_sell-price)
            first_sell = max(first_sell, first_buy+price)
            first_buy = max(first_buy, -price)

        return second_sell


solution = Solution()
print(solution.maxProfit([3, 3, 5, 0, 0, 3, 1, 4]))
