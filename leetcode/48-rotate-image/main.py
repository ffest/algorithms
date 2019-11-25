from typing import List


class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        for i in range(len(matrix)//2):
            matrix[i], matrix[len(matrix)-i-1] = matrix[len(matrix)-i-1], matrix[i]
        for i in range(len(matrix)):
            for j in range(i, len(matrix)):
                matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]

    def rotateAntiClockwise(self, matrix: List[List[int]]) -> None:
        for i in range(len(matrix)//2):
            for j in range(len(matrix)):
                matrix[j][i], matrix[j][len(matrix)-1-i] = matrix[j][len(matrix)-1-i], matrix[j][i]
        for i in range(len(matrix)):
            for j in range(i, len(matrix)):
                matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]




matrix = [
    [5, 1, 9, 11],
    [2, 4, 8, 10],
    [13, 3, 6, 7],
    [15, 14, 12, 16]
]
solution = Solution()
solution.rotate(matrix)
print(matrix)
