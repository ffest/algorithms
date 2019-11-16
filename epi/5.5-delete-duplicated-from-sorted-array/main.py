from typing import List


def delete_duplicates(a: List[int]) -> int:
    if not A:
        return 0

    write_index = 1
    for i in range(1, len(a)-1):
        if a[write_index-1] == a[i]:
            continue
        a[write_index] = a[i]
        write_index += 1

    return write_index


A = [1, 2, 3, 3, 4, 5, 5, 7, 8, 9, 9]
print(delete_duplicates(A))
