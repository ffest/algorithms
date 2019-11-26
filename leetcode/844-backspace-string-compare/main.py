class Solution:
    def backspaceCompare(self, S: str, T: str) -> bool:
        i, j = len(S) - 1, len(T) - 1
        backS = backT = 0
        while True:
            while i >= 0 and (backS > 0 or S[i] == '#'):
                backS += 1 if S[i] == '#' else -1
                i -= 1
            while j >= 0 and (backT > 0 or T[j] == '#'):
                backT += 1 if T[j] == '#' else -1
                j -= 1

            if i >= 0 and j >= 0 and S[i] == T[j]:
                i = i - 1
                j = j - 1
            else:
                return i == j == -1


solution = Solution()
S = "y#fo##f"
T = "y#f#o##f"
print(solution.backspaceCompare(S, T))
