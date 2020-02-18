from typing import List, Iterator
import collections


def examine_building_with_sunset(sequence: Iterator[int]) -> List[int]:
    buildings = collections.namedtuple('building', ('id', 'height'))
    candidates: List[buildings] = []
    for idx, height in enumerate(sequence):
        while candidates and height >= candidates[-1].height:
            candidates.pop()
        candidates.append(buildings(idx, height))

    return [c.id for c in reversed(candidates)]


mytuple = (1, 3, 2, 3, 4, 3, 2, 1)
sequence = iter(mytuple)
print(examine_building_with_sunset(sequence))
