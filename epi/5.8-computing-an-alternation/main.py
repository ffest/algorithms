from typing import List


def rearrange(a: List[int]) -> None:
    for i in range(len(a)):
        a[i:i+2] = sorted(a[i:i+2], reverse=bool(i % 2))
    return


A = [1, 2, 3, 3, 4, 5, 5, 7, 8, 9, 9]
rearrange(A)
print(A)
