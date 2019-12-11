from typing import List


class Solution:
    def is_valid(self, s: str) -> bool:
        if len(s) > 1 and s[0] == '0' or int(s) > 255:
            return False
        return True

    def restoreIpAddresses(self, s: str) -> List[str]:
        ips = []
        parts = [""] * 4
        for i in range(1, min(4, len(s))):
            parts[0] = s[:i]
            if not self.is_valid(parts[0]):
                continue
            for j in range(1, min(4, len(s) - i)):
                parts[1] = s[i:i+j]
                if not self.is_valid(parts[1]):
                    continue
                for k in range(1, min(4, len(s) - i - j)):
                    parts[2] = s[i + j:i + j + k]
                    if not self.is_valid(parts[2]):
                        continue
                    parts[3] = s[i + j + k:]
                    if not self.is_valid(parts[3]):
                        continue
                    ips.append(".".join(parts))
        return ips


solution = Solution()
print(solution.restoreIpAddresses("25525511135"))
