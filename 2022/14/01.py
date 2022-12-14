def solve(file: str) -> int:
    maxX = 0
    maxY = 0
    minX = 9999999999
    minY = 9999999999

    def map_coords(x: int, y: int) -> tuple[int, int]:
        return (x - minX, y - minY)

    wall_map: list[list[str]] = []

    def map_rock_to_wall(rock: str) -> None:
        prev: tuple[int, int] | None = None
        edges = rock.split(" -> ")
        for edge in edges:
            coords = map_coords(*[int(c) for c in edge.split(",")])
            if prev == None:
                prev = coords
                continue





    with open(file) as f:
        rocks = f.readlines()

        for rock in rocks:
            edges = rock.split(" -> ")
            for edge in edges:
                (x,y) = edge.split(",")
                maxX = max(maxX, int(x))
                maxY = max(maxY, int(y))
                minX = min(minX, int(x))
                minY = min(minY, int(y))

        print("{0} -> {1} ==> {2} -> {3}".format((minX, minY), (maxX, maxY), map_coords(minX, minY), map_coords(maxX, maxY)))

        end = map_coords(maxX, maxY)
        for _ in range(end[1]):
            wall_map.append(["."]*end[0])

        for row in wall_map:
            print("".join(row))

        

    return 0


solve("test-input.txt")

solve("input.txt")