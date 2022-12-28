direction_map = [
    (1, 0),
    (0, 1),
    (-1, 0),
    (0, -1)
]

def movement(pos: tuple[int, int], dir: int) -> tuple[int, int]:
    mov = direction_map[dir]
    return (pos[0] + mov[0], pos[1] + mov[1])

def get_tile(map: list[str], pos: tuple[int, int]):
    if pos[1] < 0 or pos[1] >= len(map) or pos[0] < 0 or pos[0] >= len(map[pos[1]]):
        return ' '

    tile = map[pos[1]][pos[0]]
    return tile

def wrap(map: list[str], pos: tuple[int, int], dir: int) -> tuple[int, int]:
    new_pos: tuple[int, int] = pos
    
    match dir:
        case 0: new_pos = (0, pos[1])
        case 1: new_pos = (pos[0], 0)
        case 2: new_pos = (len(map[pos[1]])-1, pos[1])
        case 3: new_pos = (pos[0], len(map)-1)
        case _: assert dir >= 0 and dir < 4

    while get_tile(map, new_pos) == ' ':
        new_pos = movement(new_pos, dir)

    if get_tile(map, new_pos) == '#':
        return pos

    return new_pos

def move(map: list[str], pos: tuple[int, int], dir: int, amount: int) -> tuple[int, int]:
    assert get_tile(map, pos) == '.'

    for _ in range(amount):
        new_pos = movement(pos, dir)
        match get_tile(map, new_pos):
            case ' ':
                pos = wrap(map, pos, dir)
            case '#':
                return pos
            case '.':
                pos = new_pos
            case _: pass

    assert get_tile(map, pos) == '.'
    return pos

def solve(file: str) -> int:
    map: list[str] = []
    instructions = ""
    with open(file) as f:
        map_done = False
        for row in f.readlines():
            if row.strip() == "":
                map_done = True
                continue

            if map_done:
                instructions = row.strip()
            else:
                map.append(row.strip("\n"))

    value = 0
    dir = 0
    pos = wrap(map, (0, 0), 0)
    print("Starting from", pos)
    count = 0
    for instruction in instructions:
        count += 1
        if count > 10000:
            break
        if instruction.isdigit():
            value = value * 10 + int(instruction)
        else:
            pos = move(map, pos, dir, value)
            print(count, "Move to", pos, dir, value, instruction)
            if instruction == 'R':
                dir = (dir + 1) % 4
            elif instruction == 'L':
                dir = (dir - 1) % 4
            else:
                print("WTF?", instruction)

            value = 0

    pos = move(map, pos, dir, value)
    print("Final move to", pos, dir, value)

    return (pos[0] + 1) * 4 + (pos[1] + 1) * 1000 + dir


result = solve("/Users/kurre/projects/advent-of-code/2022/22/test-input.txt")
print("Result:", result)

assert result == 6032

result = solve("input.txt")
print("Result:", result)
