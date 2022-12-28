from typing import TypeAlias

Cube: TypeAlias = tuple[int, int, int]

def find_adjacent(cube: Cube, cubes: list[Cube], space: tuple[int, int, int]) -> list[Cube]:
    to_check: list[Cube] = [cube]
    checked: set[Cube] = set()
    existing = set(cubes)

    tested: set[tuple[str, Cube]] = set()

    count = 0

    while len(to_check) > 0:
        count += 1
        current = to_check.pop(0)
        if current in checked:
            continue
        checked.add(current)
        print("Current", current)
        for axis in range(3):
            for dir in [-1, 1]:
                new = list(current)
                new[axis] += dir
                if new[axis] < 0 or new[axis] > space[axis]:
                    tested.add(("Outside", tuple(new)))
                    print(("Outside", tuple(new)))
                    continue
                new = tuple(new)
                if new in checked:
                    tested.add(("Checked", new))
                    print(("Checked", new))
                    continue
                if new in existing:
                    tested.add(("Existing", new))
                    print(("Existing", new))
                    continue
                to_check.append(new)
                tested.add(("Valid", new))
                print(("Valid", new))

        print("Check", count, len(checked))


    print("Tested:", tested)
    return list(checked)

def check_cubes(cubes: list[Cube]) -> list[Cube]:
    a: int = 0
    b: int = 1
    c: int = 2

    cubes.sort(key=lambda cube: cube[a])
    cubes.sort(key=lambda cube: cube[b])
    cubes.sort(key=lambda cube: cube[c])

    space = (0, 0, 0)
    for cube in cubes:
        if cube[0] > space[0]:
            space = (cube[0], space[1], space[2])
        if cube[1] > space[1]:
            space = (space[0], cube[1], space[2])
        if cube[2] > space[2]:
            space = (space[0], space[1], cube[2])

    prevCube: Cube | None = None
    checked: set[Cube] = set()
    cubes_copy = cubes.copy()
    for cube in cubes_copy:
        if prevCube != None:
            # is at edge?
            if prevCube[b] != cube[b] or prevCube[c] != cube[c]:
                #print("Edge: ", prevCube, cube)
                prevCube = cube
                continue
            if (prevCube[a] + 1) != cube[a]:
                #checked.add(cube)
                to_check = (cube[a] - 1, cube[b], cube[c])
                if to_check in checked:
                    prevCube = cube
                    continue

                print("Checking", to_check)
                chain = find_adjacent(to_check, cubes, space)
                for link in chain:
                    checked.add(link)

                trapped = True
                for link in chain:
                    if link[0] == 0 or link[1] == 0 or link[2] == 0 \
                        or link[0] == space[0] or link[1] == space[1] or link[2] == space[2]:
                        trapped = False
                        break
                
                print("Trapped" if trapped else "Open", chain)

                if trapped:
                    cubes.extend(chain)

                #print("Cube {0} does not follow cube {1} ({2},{3},{4})".format(cube, prevCube, (prevCube[a] + 1) != cube[a], prevCube[b] != cube[b], prevCube[c] != cube[c]))
            #else:
                #print("Cube {0} follows cube {1}".format(cube, prevCube))
        prevCube = cube

    return cubes


def count_direction(cubes: list[Cube], dir_index: int) -> int:
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
    prevCube: Cube | None = None
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


    cubes: list[Cube] = [tuple([int(coord) for coord in c.split(",")]) for c in input]

    print("Before:", len(cubes))
    cubes = check_cubes(cubes)
    print("After:", len(cubes))
    count = 0
    for i in range(3):
        result = count_direction(cubes, i)
        #print("Direction {0}, result {1}".format(i, result))
        count += result
    
    return count


result = solve("test-input.txt")
print("Result: {0}".format(result))
assert(result == 58)

result = solve("input.txt")
print("Result: {0}".format(result))


