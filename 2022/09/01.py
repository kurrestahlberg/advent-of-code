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
    if tail.distance_to(head) <= 0:
        return tail

    return tail + (head - tail).unit()
    

with open("input.txt") as f:
    moves = [map_a_move(m.strip()) for m in f.readlines()]

    print(set([Coordinate(6, 6), Coordinate(5, 5), Coordinate(6,6)]))

    print(moves[0])
    print(Coordinate(-15, 4).unit())
    print(Coordinate(-15, 0).unit())

    head = Coordinate(0,0)
    tail = Coordinate(0,0)

    tail_positions: list[Coordinate] = [tail]

    count = 0

    for move in moves:
        head = head + move
        #print("Head goes to {0}".format(head))
        while tail.distance_to(head) > 1:
            tail = move_tail_step(tail, head)
            #print("Tail goes to {0}".format(tail))
            tail_positions.append(tail)

    print(len(tail_positions))
    print(len(set(tail_positions)))