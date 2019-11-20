from typing import List, Iterator
import random
import itertools


def random_sampling(stream: Iterator[int], k: int) -> List[int]:
    sample = list(itertools.islice(stream, k))
    seen = k
    for x in stream:
        seen += 1
        idx = random.randrange(seen)
        if idx < k:
            sample[idx] = x

    return sample


k = 6
A = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
random_sampling(A, k)
print(random_sampling(A, k))
