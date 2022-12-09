from __future__ import annotations

class Coordinate:
    def __init__(self, x, y) -> None:
        self.x = x
        self.y = y

    def __add__(self, other: Coordinate):
        return Coordinate(self.x + other.x, self.y + other.y)

    def __mul__(self, multiplier: int) -> Coordinate:
        return Coordinate(self.x * multiplier, self.y * multiplier)

    def __str__(self) -> str:
        return str((self.x, self.y))

    def __eq__(self, o: Coordinate) -> bool:
        return self.x == o.x and self.y == o.y

    def __hash__(self) -> int:
        return hash(("Coordinate", self.x, self.y))

    def distance_to(self, other: Coordinate) -> int:
        return max(abs(self.x - other.x), abs(self.y - other.y))

    def __sub__(self, other: Coordinate) -> Coordinate:
        return Coordinate(self.x - other.x, self.y - other.y)

    def unit(self) -> Coordinate:
        return Coordinate(round(self.x / max(1, abs(self.x))), round(self.y / max(1, abs(self.y))))

direction_map = {
    'U': Coordinate(0, -1),
    'D': Coordinate(0, 1),
    'L': Coordinate(-1, 0),
    'R': Coordinate(1, 0)
}

def map_a_move(input: str) -> Coordinate:
    [dir, distance] = input.split()
    output = direction_map[dir] * int(distance)
    # print("From {0} we get {1}/{2} and the result is {3}".format(input, dir, distance, output))
    return output

def move_head(pos: Coordinate, move: Coordinate) -> Coordinate:
    return pos + move

def move_tail_step(tail: Coordinate, head: Coordinate) -> Coordinate:
    return tail + (head - tail).unit()
    

with open("input.txt") as f:
    moves = [map_a_move(m.strip()) for m in f.readlines()]

    print(set([Coordinate(6, 6), Coordinate(5, 5), Coordinate(6,6)]))

    print(moves[0])
    print(Coordinate(-15, 4).unit())
    print(Coordinate(-15, 0).unit())

    rope = [Coordinate(0, 0)] * 10
    tail = len(rope) - 1

    tail_positions: list[Coordinate] = [rope[tail]]

    count = 0

    for move in moves:
        target = rope[0] + move
        #print("Target is {0}, it is {1} steps away".format(target, rope[0].distance_to(target)))
        while rope[0].distance_to(target) > 0:
            rope[0] = move_tail_step(rope[0], target)
            #print("Head goes to {0}".format(rope[0]))

            for i in range(1, len(rope), 1):
                if rope[i].distance_to(rope[i - 1]) > 1:
                    rope[i] = move_tail_step(rope[i], rope[i - 1])
                    if i == tail:
                        #print("Knot {1} goes to {0}".format(rope[i], i))
                        tail_positions.append(rope[i])

        #count += 1
        #if count > 3:
        #    break

    print(len(tail_positions))
    print(len(set(tail_positions)))
    #print([str(pos) for pos in tail_positions])