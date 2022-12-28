
def count_direction(cubes: list[tuple[int, int, int]], dir_index: int) -> int:
    a: int = dir_index % 3
    b: int = (dir_index + 1) % 3
    c: int = (dir_index + 2) % 3

    #print(a, b, c)

    cubes.sort(key=lambda cube: cube[a])
    cubes.sort(key=lambda cube: cube[b])
    cubes.sort(key=lambda cube: cube[c])

    #for cube in cubes:
    #    print(cube)

    count = 2
    prevCube: tuple[int, int, int] | None = None
    for cube in cubes:
        if prevCube != None:
            if (prevCube[a] + 1) != cube[a] or prevCube[b] != cube[b] or prevCube[c] != cube[c]:
                #print("Cube {0} does not follow cube {1} ({2},{3},{4})".format(cube, prevCube, (prevCube[a] + 1) != cube[a], prevCube[b] != cube[b], prevCube[c] != cube[c]))
                count += 2
            #else:
                #print("Cube {0} follows cube {1}".format(cube, prevCube))
        prevCube = cube

    return count


def solve(file: str) -> int:
    with open(file) as f:
        input = f.readlines()


    cubes: list[tuple[int, int, int]] = [tuple([int(coord) for coord in c.split(",")]) for c in input]
    count = 0
    for i in range(3):
        result = count_direction(cubes, i)
        #print("Direction {0}, result {1}".format(i, result))
        count += result
    
    return count


result = solve("test-input.txt")
print("Result: {0}".format(result))
assert(result == 64)

result = solve("input.txt")
print("Result: {0}".format(result))


