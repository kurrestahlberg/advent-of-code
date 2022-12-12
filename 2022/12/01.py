from __future__ import annotations
from typing import TypeAlias

Coord: TypeAlias = tuple[int, int]
Map: TypeAlias = list[str]

class Node:
    pos: Coord
    parent: Node | None
    path_len: int
    valid: bool
    tile: str

    def __init__(self, pos: Coord, tile_map: Map, finish: Coord, parent: Node | None):
        self.pos = pos

        if parent is None:
            self.path_len = 0
            self.parent = None
        else: 
            self.set_parent(parent)

        self.valid = self.pos[0] >= 0 and self.pos[0] < len(tile_map[0]) and self.pos[1] >= 0 and self.pos[1] < len(tile_map)
        if self.valid:
            self.tile = get_tile(tile_map, pos)
            if self.tile == "S":
                self.tile = "a"
            elif self.tile == "E":
                self.tile = "z"
        else:
            self.tile ="a"

        self.distance_left = distance(pos, finish)

    def distance_to(self, other: Coord) -> int:
        return distance(self.pos, other)

    def set_parent(self, parent: Node) -> None:
        self.parent = parent
        self.path_len = parent.path_len + 1

    def can_move_to(self, other: Node) -> bool:
        #print(self.pos, other.pos, other.valid, ord(self.tile), ord(other.tile))

        if other.valid == False:
            return False

        if ord(self.tile) >= ord(other.tile) - 1:
            return True

        return False

    def get_score(self) -> int:
        height = ord("z") - ord(self.tile)
        return int(self.path_len * self.path_len + self.distance_left * self.distance_left + height * height)

    def __str__(self) -> str:
        return "{0} -> {1} ({2}/{3}/{4})".format(self.pos, self.tile, self.get_score(), self.path_len, self.distance_left)

directions = [
    (0, 1),
    (0, -1),
    (1, 0),
    (-1, 0)
]

def find_tile(tile_map: list[str], tile: str) -> Coord:
    for (y, row) in enumerate(tile_map):
        for x in range(len(row)):
            if row[x] == tile:
                print("Looking for {0}, found {1} at {2}".format(tile, row[x], (x, y)))
                return (x, y)

    return (-1, -1)

def get_tile(tile_map: list[str], pos: Coord) -> str:
    return tile_map[pos[1]][pos[0]]

def distance(here: Coord, there: Coord) -> int:
    return abs(here[0]-there[0]) + abs(here[1]-there[1])

def move(here: Coord, step: Coord) -> Coord:
    return (here[0] + step[0], here[1] + step[1])

def add_to_priority_queue(queue: list[Node], new: Node) -> None:
    for item in queue:
        if item.pos == new.pos:
            if item.get_score() <= new.get_score():
                return
            
            queue.remove(item)

    for (i, item) in enumerate(queue):
        if item.get_score() > new.get_score():
            queue.insert(i, new)
            return

    queue.append(new)

with open("input.txt") as f:
    elevation_map = [row.strip() for row in f.readlines()]

    finish: Coord = find_tile(elevation_map, "E")
    start: Node = Node(find_tile(elevation_map, "S"), elevation_map, finish, None)

    #print("Start is {0}".format(start))

    possibilities = [start]
    visited: dict[Coord, Node] = {}

    while len(possibilities) > 0:
        current = possibilities.pop(0)
        visited[current.pos] = current
        distance_left = current.distance_to(finish)

        #print("{0}, distance left {1}, score: {2}".format(current, distance_left, current.get_score()))
        if distance_left == 0:
            print(start, finish, len(elevation_map[0]), len(elevation_map), current.path_len)

            
            route: list[list[str]] = []
            for _ in range(len(elevation_map)):
                route.append(["."]*len(elevation_map[0]))

            while current.parent:
                print(current)
                current = current.parent
                route[current.pos[1]][current.pos[0]] = "*"
            print(current)
            route[current.pos[1]][current.pos[0]] = "*"

            for row in route:
                print("".join(row))
            break

        for direction in directions:
            new_pos = move(current.pos, direction)
            node = visited.get(new_pos)
            if node == None:
                node = Node(new_pos, elevation_map, finish, current)
            elif node.path_len <= current.path_len + 1:
                continue

            if current.can_move_to(node) == False:
                #print("Can't move from {0}/{1} to {2}/{3}".format(current.pos, current.tile, node.pos, node.tile))
                continue

            #print("Okay to move from {0}/{1} to {2}/{3}".format(current.pos, current.tile, node.pos, node.tile))

            node.set_parent(current)
            add_to_priority_queue(possibilities, node)
