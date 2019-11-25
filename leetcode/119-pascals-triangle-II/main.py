from typing import List


class Solution:
    def getRow(self, rowIndex: int) -> List[int]:
        row = [1]
        for i in range(1, rowIndex+1):
            row = [1] + [row[x] + row[x - 1] for x in range(1, i)] + [1]
        return row


solution = Solution()
print(solution.getRow(5))
