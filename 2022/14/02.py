from typing import TypeAlias

Coord: TypeAlias = tuple[int, int]
Map: TypeAlias = list[list[str]]


def solve(file: str) -> int:
    maxX = 0
    maxY = 0
    minX = 9999999999
    minY = 0

    wall_map: Map = []

    def map_coords(x: int, y: int) -> Coord:
        return (x - minX, y - minY)

    def map_edge_to_wall(start: Coord, end: Coord) -> None:
        ystart = min(start[1], end[1])
        ylen = abs(start[1] - end[1]) + 1
        xstart = min(start[0], end[0])
        xlen = abs(start[0] - end[0]) + 1

        for y in range(ystart, ystart + ylen, 1):
            for x in range(xstart, xstart + xlen, 1):
                wall_map[y][x] = '#'

    def map_rock_to_wall(rock: str) -> None:
        prev: Coord | None = None
        nodes = rock.split(" -> ")
        for node in nodes:
            coords = map_coords(*[int(c) for c in node.split(",")])
            if prev != None:
                map_edge_to_wall(prev, coords)
            prev = coords

    def add_sand():
        x = 500 - minX
        for y in range(len(wall_map)):
            if wall_map[y][x] != '.':
                if y == 0:
                    return False
                
                if wall_map[y][x - 1] == '.':
                    x -= 1
                elif wall_map[y][x + 1] == '.':
                    x += 1
                else:
                    wall_map[y - 1][x] = '*'
                    return True

        wall_map[-1][x] = '*'
        return True

    def print_map():
        print()
        for row in wall_map:
            print("".join(row))

    with open(file) as f:
        rocks = f.readlines()

    for rock in rocks:
        nodes = rock.split(" -> ")
        for node in nodes:
            (x,y) = node.split(",")
            maxX = max(maxX, int(x))
            maxY = max(maxY, int(y))
            minX = min(minX, int(x))
            minY = min(minY, int(y))

    print("{0} -> {1} ==> {2} -> {3}".format((minX, minY), (maxX, maxY), map_coords(minX, minY), map_coords(maxX, maxY)))

    minX -= maxY
    maxX += maxY
    end = map_coords(maxX + 1, maxY + 2)
    for _ in range(end[1]):
        wall_map.append(["."]*end[0])

    for rock in rocks:
        map_rock_to_wall(rock)

    print_map()

    count = 0
    while True:
        res = add_sand()
        if res == True:
            count += 1
        else:
            print_map()
            return count

    return 0


result = solve("test-input.txt")
print("Result: {0}".format(result))

result = solve("input.txt")
print("Result: {0}".format(result))
