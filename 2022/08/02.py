# https://adventofcode.com/2022/day/8

def get_index_func(colCount: int):
    def idx(x: int, y: int) -> int:
        return y * colCount + x

    return idx

with open("/Users/kurre/advent-of-code/2022/08/input.txt") as file:
    rows = [l.strip() for l in file.readlines()]

    rowCount = len(rows)
    colCount = len(rows[0])

    scenic_score: list[int] = [1] * rowCount * colCount
    idx = get_index_func(colCount)

    # No need to check the edges as they always end up being 0
    for ty in range(1, rowCount - 1, 1):
        for tx in range(1, colCount - 1, 1):
            height = int(rows[ty][tx])

            for x in range(tx + 1, colCount, 1):
                val = int(rows[ty][x])
                if val >= height or x == colCount - 1:
                    scenic_score[idx(tx, ty)] *= abs(x - tx)

                    if tx == 41 and ty == 86:
                        print("X1 score: {0}, cur: {1} ({2}, {3})".format(scenic_score[idx(tx, ty)], abs(x - tx), x, tx))
                    break

            for x in range(tx - 1, -1, -1):
                val = int(rows[ty][x])
                if val >= height or x == 0:
                    scenic_score[idx(tx, ty)] *= abs(x - tx)
                    if tx == 41 and ty == 86:
                        print("X2 score: {0}, cur: {1} ({2}, {3})".format(scenic_score[idx(tx, ty)], abs(x - tx), x, tx))
                    break

            for y in range(ty + 1, rowCount, 1):
                val = int(rows[y][tx])
                if val >= height or y == rowCount - 1:
                    scenic_score[idx(tx, ty)] *= abs(y - ty)
                    if tx == 41 and ty == 86:
                        print("Y1 score: {0}, cur: {1} ({2}, {3})".format(scenic_score[idx(tx, ty)], abs(y - ty), y, ty))
                    break

            for y in range(ty - 1, -1, -1):
                val = int(rows[y][tx])
                if val >= height or y == 0:
                    scenic_score[idx(tx, ty)] *= abs(y - ty)
                    if tx == 41 and ty == 86:
                        print("Y2 score: {0}, cur: {1} ({2}, {3})".format(scenic_score[idx(tx, ty)], abs(y - ty), y, ty))
                    break

    for i in range(rowCount):
        print("".join([str(val) for val in scenic_score[i * colCount:(i+1)*colCount]]))

    print("Max: {0}".format(max(scenic_score)))

    max = (0,0)
    for y in range(rowCount):
        for x in range(colCount):
            if scenic_score[idx(x, y)] > scenic_score[idx(max[0], max[1])]:
                max = (x, y)

    print("Max is at {0} and it is {1}".format(max, scenic_score[idx(max[0], max[1])]))
        
    