class Solution:
    def simplifyPath(self, path: str) -> str:
        stack = []
        parts = path.split("/")
        for part in parts:
            if part == "" or part == ".":
                continue
            elif part == "..":
                if stack:
                    stack.pop()
            else:
                stack.append(part)
        return "/" + "/".join(stack)


solution = Solution()
print(solution.simplifyPath("/a//b////c/d//././/.."))
