class Solution:
    def isValid(self, s: str) -> bool:
        if len(s) == 0:
            return True
        elif len(s) % 2 == 1:
            return False

        stack, closes = [], {'(': ')', '[': ']', '{': '}'}
        for c in s:
            if c in closes:
                stack.append(c)
            elif not stack or closes[stack.pop()] != c:
                return False
        return not stack


solution = Solution()
print(solution.isValid("()[]{}"))
