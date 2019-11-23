from typing import List, Dict
import random


def random_subset(k: int, n: int) -> List[int]:
    changed: Dict[int, int] = {}
    for i in range(k):
        idx = random.randrange(i, n)
        idx_mapped = changed.get(idx, idx)
        i_mapped = changed.get(i, i)
        changed[i] = idx_mapped
        changed[idx] = i_mapped
    return [changed[i] for i in range(k)]


print(random_subset(6, 10))
