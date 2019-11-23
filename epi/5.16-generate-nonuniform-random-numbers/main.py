from typing import List
import bisect
import itertools
import random


def nonuniform_random_number_generation(values: List[int], probabilities: List[float]) -> int:
    accumulated = list(itertools.accumulate(probabilities))
    return values[bisect.bisect(accumulated, random.random())]


values = [4, 6, 7, 8]
probabilities = [0.2, 0.3, 0.4, 0.1]
print(nonuniform_random_number_generation(values, probabilities))
