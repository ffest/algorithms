from typing import List


class Solution:
    def canCross(self, stones: List[int]) -> bool:
        if stones[1] != 1:
            return False
        cache = {stone: set() for stone in stones}
        cache[0].add(1)
        for i in range(len(stones)-1):
            for step in cache[stones[i]]:
                reach = step+stones[i]
                if reach == stones[-1]:
                    return True
                if reach in cache:
                    cache[reach].add(step)
                    cache[reach].add(step+1)
                    if step > 1:
                        cache[reach].add(step-1)
        return False


solution = Solution()
print(solution.canCross([0, 1, 3, 5, 6, 8, 12, 17]))
