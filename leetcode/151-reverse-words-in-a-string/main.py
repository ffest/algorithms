class Solution:
    def reverseWords(self, s: str) -> str:
        return ' '.join(reversed(s.split()))


solution = Solution()
print(solution.reverseWords("a good   example"))
