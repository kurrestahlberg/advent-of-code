from __future__ import annotations
from enum import Enum

class Result(Enum):
    CORRECT = 1,
    INCORRECT = -1,
    UNDETERMINED = 0

def compare(left: list|int, right: list|int) -> Result:
    if type(left) == int and type(right) == int:
        if left < right:
            return Result.CORRECT
        if left > right:
            return Result.INCORRECT
        else:
            return Result.UNDETERMINED

    if type(left) == int:
        left = [left]
    if type(right) == int:
        right = [right]

    for i in range(min(len(left), len(right))):
        res = compare(left[i], right[i])
        if res != Result.UNDETERMINED:
            return res

    if len(left) > len(right):
        return Result.INCORRECT
    if len(left) == len(right):
        return Result.UNDETERMINED
    else:
        return Result.CORRECT

def solve(file: str) -> int:
    count = 0

    with open(file) as f:
        data = f.read()
        pairs = data.split("\n\n")

        for (i, pair) in enumerate(pairs):
            pair = pair.strip()
            packets = [eval(p) for p in pair.split("\n")]

            res = compare(packets[0], packets[1])
            if res == Result.CORRECT:
                count += (i + 1)

    return count


res = solve("test-input.txt")
print("Test input resulted in {0}, expected {1}".format(res, 13))
assert(res == 13)

res = solve("input.txt")
print("Actual input resulted in {0}".format(res))
