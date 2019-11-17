from typing import List


def apply_permutation(perm: List[int], a: List[int]) -> None:
    for i in range(len(a)):
        while perm[i] != i:
            a[perm[i]], a[i] = a[i], a[perm[i]]
            perm[perm[i]], perm[i] = perm[i], perm[perm[i]]
    return


perm = [2, 0, 1, 3]
A = [1, 2, 3, 4]
apply_permutation(perm, A)
print(A)
