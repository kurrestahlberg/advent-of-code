# https://adventofcode.com/2022/day/8

def get_index_func(colCount: int):
    def idx(x: int, y:int) -> int:
        return y * colCount + x

    return idx

with open("input.txt") as file:
    rows = [l.strip() for l in file.readlines()]

    rowCount = len(rows)
    colCount = len(rows[0])

    visible: list[int] = [0] * rowCount * colCount
    idx = get_index_func(colCount)

    for y, row in enumerate(rows):
        tallest = -1
        for x in range(colCount):
            val = int(row[x])
            if val > tallest:
                visible[idx(x, y)] = 1
                tallest = val

    for y, row in enumerate(rows):
        tallest = -1
        for x in range(colCount - 1, 0, -1):
            val = int(row[x])
            if val > tallest:
                visible[idx(x, y)] = 1
                tallest = val

    for x in range(colCount):
        tallest = -1
        for y in range(rowCount):
            val = int(rows[y][x])
            if val > tallest:
                visible[idx(x, y)] = 1
                tallest = val

    for x in range(colCount):
        tallest = -1
        for y in range(rowCount - 1, 0, -1):
            val = int(rows[y][x])
            if val > tallest:
                visible[idx(x, y)] = 1
                tallest = val

    for i in range(rowCount):
        print("".join([str(val) for val in visible[i * colCount:(i+1)*colCount]]))

    print("Total: {0}".format(sum(visible)))
    