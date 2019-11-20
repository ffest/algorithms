from typing import List
import random


def random_sampling(k: int, a: List[int]) -> None:
    for i in range(k):
        idx = random.randint(i, len(a) - 1)
        a[idx], a[i] = a[i], a[idx]


def compute_random_permutation(n: int) -> List[int]:
    permutation = list(range(n))
    # Same same
    # random_sampling(n, permutation)
    for i in range(n):
        idx = random.randrange(n)
        permutation[idx], permutation[i] = permutation[i], permutation[idx]

    return permutation


k = 10
print(compute_random_permutation(k))
