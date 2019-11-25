from typing import List


class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        result = [[1 for _ in range(i+1)] for i in range(numRows)]
        for i in range(2, numRows):
            for j in range(1, i):
                result[i][j] = result[i-1][j-1] + result[i-1][j]
        return result


solution = Solution()
print(solution.generate(5))
