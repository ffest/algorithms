from typing import List


class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        cache = set()
        for i in range(len(board)):
            for j in range(len(board[0])):
                if board[i][j] == ".":
                    continue

                row = str(i)+"row"+board[i][j]
                col = str(j)+"col"+board[i][j]
                box = str(i//3)+str(j//3)+"col"+board[i][j]
                if row in cache or col in cache or box in cache:
                    return False
                cache.add(row)
                cache.add(col)
                cache.add(box)
        return True


solution = Solution()
board = [
    ["5", "3", ".", ".", "7", ".", ".", ".", "."],
    ["6", ".", ".", "1", "9", "5", ".", ".", "."],
    [".", "9", "8", ".", ".", ".", ".", "6", "."],
    ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
    ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
    ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
    [".", "6", ".", ".", ".", ".", "2", "8", "."],
    [".", ".", ".", "4", "1", "9", ".", ".", "5"],
    [".", ".", ".", ".", "8", ".", ".", "7", "9"]
]
print(solution.isValidSudoku(board))
