from typing import List
import random


def random_sampling(k: int, a: List[int]) -> None:
    for i in range(k):
        idx = random.randint(i, len(a) - 1)
        a[idx], a[i] = a[i], a[idx]


k = 6
A = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
random_sampling(k, A)
print(A[:k+1])
