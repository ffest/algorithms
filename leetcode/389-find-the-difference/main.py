class Solution:
    def findTheDifference(self, s: str, t: str) -> str:
        diff = 0
        for i in range(len(s)):
            diff ^= ord(s[i])
            diff ^= ord(t[i])
        diff ^= ord(t[len(t)-1])
        return chr(diff)


solution = Solution()
print(solution.findTheDifference("abcd", "abcde"))
