from typing import List


def can_reach_end(a: List[int]) -> bool:
    longest, last = 0, len(a) - 1
    for i in range(len(a) - 1):
        longest = max(longest, i + a[i])
        if longest >= last:
            break
    return longest >= last


def min_steps_to_end(a: List[int]) -> int:
    dp = [1 << 31 - 1] * len(a)
    dp[0] = 0
    for i in range(len(a) - 1):
        for j in range(a[i]+1):
            if i+j >= len(a):
                break
            dp[i+j] = min(dp[i+j], dp[i]+1)
    return dp[len(a)-1]


A = [4, 2, 2, 1, 1, 2, 1, 1, 0]
print(min_steps_to_end(A))
