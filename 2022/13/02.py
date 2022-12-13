from __future__ import annotations
from enum import Enum
from functools import cmp_to_key

CORRECT = 1
INCORRECT = -1
UNDETERMINED = 0

def compare(left: list|int, right: list|int) -> int:
    if type(left) == int and type(right) == int:
        if left < right:
            return CORRECT
        if left > right:
            return INCORRECT
        else:
            return UNDETERMINED

    if type(left) == int:
        left = [left]
    if type(right) == int:
        right = [right]

    for i in range(min(len(left), len(right))):
        res = compare(left[i], right[i])
        if res != UNDETERMINED:
            return res

    if len(left) > len(right):
        return INCORRECT
    if len(left) == len(right):
        return UNDETERMINED
    else:
        return CORRECT

def solve(file: str) -> int:
    count = 1

    with open(file) as f:
        data = f.readlines()

        rows = [p.strip() for p in data]
        filteredrows = filter(lambda p: len(p) > 0, rows)

        packets = [eval(p) for p in filteredrows]
        packets.append([[2]])
        packets.append([[6]])

        packets.sort(key = cmp_to_key(compare), reverse=True)

        for (i, packet) in enumerate(packets):
            if compare(packet, [[2]]) == 0 or compare(packet, [[6]]) == 0:
                count *= (i + 1)
    return count


res = solve("/Users/kurre/advent-of-code/2022/13/test-input.txt")
print("Test input resulted in {0}, expected {1}".format(res, 13))
assert(res == 140)

res = solve("input.txt")
print("Actual input resulted in {0}".format(res))
