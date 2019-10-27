from typing import List


def dutch_national_partition(pivot_index: int, a: List[int]) -> None:
    pivot = a[pivot_index]

    smaller, equal, larger = 0, 0, len(a) - 1

    while equal <= larger:
        if a[equal] > pivot:
            a[equal], a[larger] = a[larger], a[equal]
            larger -= 1
        elif a[equal] == pivot:
            equal += 1
        else:
            a[smaller], a[equal] = a[equal], a[smaller]
            smaller += 1
            equal += 1


A = [4, 2, 2, 1, 0, 2, 3, -2, -1]
dutch_national_partition(5, A)
print(A)
