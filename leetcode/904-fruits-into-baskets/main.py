from typing import List
import collections


class Solution:
    def totalFruit(self, tree: List[int]) -> int:
        left, count, seen = 0, 0, collections.Counter()
        for right in range(len(tree)):
            seen[tree[right]] += 1
            while len(seen) > 2:
                seen[tree[left]] -= 1
                if not seen[tree[left]]:
                    seen.pop(tree[left])
                left += 1
                break
            count = max(count, right-left+1)
        return count


solution = Solution()
print(solution.totalFruit([3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4]))
