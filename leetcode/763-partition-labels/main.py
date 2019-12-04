from typing import List, Dict

class Solution:
    def partitionLabels(self, s: str) -> List[int]:
        ends: Dict[str, int] = {}
        for i in range(len(s)):
            ends[s[i]] = i

        partitions = []
        prev, current = 0, 0
        for i in range(len(s)):
            current = max(current, ends[s[i]])
            if i == current:
                partitions.append(current-prev+1)
                prev = current+1

        return partitions


solution = Solution()
s = "ababcbacadefegdehijhklij"
print(solution.partitionLabels(s))