from typing import List


def replace_and_remove(size: int, s: List[str]) -> int:
    write_idx, a_count = 0, 0
    for i in range(size):
        if s[i] != "b":
            s[write_idx] = s[i]
            write_idx += 1
        if s[i] == "a":
            a_count += 1

    current = write_idx - 1
    write_idx += a_count - 1
    total = write_idx + 1
    while current >= 0:
        if s[current] == 'a':
            s[write_idx-1:write_idx+1] = 'dd'
            write_idx -= 2
        else:
            s[write_idx] = s[current]
            write_idx -= 1
        current -= 1
    return total


s = ["a", "c", "d", "b", "b", "c", "a"]
print(replace_and_remove(4, s))
print(s)

