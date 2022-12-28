from __future__ import annotations
from typing import TypeAlias
import re

#from math import sqrt

State: TypeAlias = tuple[int, int, int, int]
Recipe: TypeAlias = tuple[int, int, tuple[int, int], tuple[int, int]]

def can_build(needs: tuple[int, int], haves: tuple[int, int]) -> bool:
    if haves[0] >= needs[0] and haves[1] >= needs[1]:
        return True

    return False

class Node:
    def __init__(self, parent: Node | None, recipe: Recipe, stuff: State, robots: State, time_left: int) -> None:
        self.parent = parent
        self.stuff = stuff
        self.robots = robots
        self.time_left = time_left
        self.recipe = recipe
        self.heuristic = self.calc_heuristic()
        #if self.heuristic < self.stuff[3]:
        #    print("WTF?!", self.heuristic)
        #    print(self.recipe, self.stuff)


    def collect_stuff(self, stuff: State) -> State:
        return tuple(i + j for i, j in zip(stuff, self.robots))

    def add_robot(self, robot_type: int) -> State:
        adder = tuple(1 if i == robot_type else 0 for i in range(4))
        return tuple(i + j for i, j in zip(adder, self.robots))

    def try_building(self, robot_type: int) -> Node | None:
        r = self.recipe[robot_type]

        if type(r) is tuple:
            if not can_build(r, (self.stuff[0], self.stuff[robot_type - 1])):
                return None

            stuff = tuple(self.stuff[i] - (r[i] if i == 0 else 0) for i in range(4))
            stuff = tuple(stuff[i] - (r[1] if i == (robot_type - 1) else 0) for i in range(4))
            #assert stuff[0] >= 0 and stuff[1] >= 0 and stuff[2] >= 0 and stuff[3] >= 0
        elif type(r) is int:
            if self.stuff[0] < r:
                return None

            stuff = tuple(self.stuff[i] - (r if i == 0 else 0) for i in range(4))
            #assert stuff[0] >= 0 and stuff[1] >= 0 and stuff[2] >= 0 and stuff[3] >= 0
        else:
            return None
        
        node = Node(self, self.recipe, self.collect_stuff(stuff), self.add_robot(robot_type), self.time_left - 1)
        return node


    def run_step(self) -> list[Node]:
        result: list[Node] = []

        if self.time_left == 0:
            return result
        
        if self.time_left > 1:
            for i in range(4):
                if i < 1 and self.stuff[i + 1] > 0:
                    continue
                node = self.try_building(i)
                #if self.robots == (1, 5, 0, 0):
                #    print("Robot type: {0} -> {1}".format(i, node))
                if node is not None:
                    result.append(node)

        node = Node(self, self.recipe, self.collect_stuff(self.stuff), self.robots, self.time_left - 1)
        #if self.stuff == (4, 15, 0, 0) and self.robots == (1, 3, 0, 0):
        #    print("No build -> {0}".format(node))
        result.append(node)

        return result

    def get_geodes(self) -> int:
        return self.stuff[3]

    def sort_index(self) -> int:
        return 1000 * (self.stuff[3] + (self.robots[3] * self.time_left)) + \
                100 * (self.stuff[2] + (self.robots[2] * self.time_left)) + \
                 10 * (self.stuff[1] + (self.robots[1] * self.time_left)) + \
                  1 * (self.stuff[0] + (self.robots[0] * self.time_left))

    def calc_heuristic(self) -> int:
        if self.time_left == 0:
            return self.stuff[3]

        if self.robots[0] > max(self.recipe[0], self.recipe[1], self.recipe[2][0], self.recipe[3][0]):
            return 0
        
        if self.robots[1] > self.recipe[2][1]:
            return 0

        if self.robots[2] > self.recipe[3][1]:
            return 0

        #max0 = self.stuff[0] + sum([int(i / self.recipe[0]) + self.robots[0] for i in range(self.time_left)])
        #time_left = max(self.time_left - max(0, self.recipe[1] - self.stuff[0]), 0)
        #max1 = self.stuff[1] + sum([min(i, max0 / self.recipe[1]) + self.robots[1] for i in range(time_left)])

        #b = self.stuff[1] >> 1
        #need = max(0.5, self.recipe[2][1] - self.stuff[1])
        #time_needed = max(1, round(sqrt(b*b + 2*need) - b))
        #max2 = self.stuff[2] + sum([min(round(i / time_needed), max1 / self.recipe[2][1]) + self.robots[2] for i in range(time_left)])

        #b = self.stuff[2] >> 1
        #need = max(0.5, self.recipe[3][1] - self.stuff[2])
        #time_needed_2 = max(1, round(sqrt(b*b + 2*need) - b))
        #max3 = self.stuff[3] + sum([min(round(i / time_needed_2), max2 / self.recipe[3][1]) + self.robots[3] for i in range(time_left)])


        #if self.stuff[3] > 0:
        #    return self.stuff[3] + round((self.time_left * self.robots[3]) / self.recipe[3][1])
        #elif self.stuff[2] > 0:
            
        #elif self.stuff[1] > 0:
        #else:

        max3 = 1

        #if self.parent is not None and self.parent.stuff == (4, 15, 0, 0) and self.parent.robots == (1, 3, 0, 0):
        #    print(max0, max1, max2, max3, self.time_left, time_left, time_needed, time_needed_2, need)
        #    print(self)

        # exit(1)

        return int(max3)
    
    def __str__(self) -> str:
        return "{0}: {1} {2} {3} -> heuristic: {4}".format(32 - self.time_left, self.recipe, self.stuff, self.robots, self.heuristic if hasattr(self, 'heuristic') else -1)

"""
y = x * (((x - 1) / 2) + 1) + x*k
y = x * (((x - 1) / 2) + k + 1)
y = (x2 - x) / 2 + xk + x
y = 0.5x2 - 0.5x + xk + x

55 = 0.5 * 100 - 5 + 0 + 10 = 55

0.5x2 + 0.5*k*x - y = 0

-b +- sqrt(b2 + 2y)

-5 +- sqrt(25 + 110)




"""

def decode_recipe(data: str) -> Recipe:
    res = re.match("Blueprint ([0-9]+): Each ore robot costs ([0-9]+) ore. Each clay robot costs ([0-9]+) ore. Each obsidian robot costs ([0-9]+) ore and ([0-9]+) clay. Each geode robot costs ([0-9]+) ore and ([0-9]+) obsidian.", data)
    if res is None:
        print("Failed to parse \"{0}\"".format(data))
        exit(666)
    
    values = [int(val) for val in res.groups()[1:]]
    return (values[0], values[1], (values[2], values[3]), (values[4], values[5]))

def add_to_priority_queue(queue: list[Node], new: Node) -> None:
    for (i, item) in enumerate(queue):
        if item.heuristic < new.heuristic:
            queue.insert(i, new)
            return

    queue.append(new)

def solve(file: str) -> int:
    with open(file) as f:
        recipe_data = f.readlines()
    
    recipes = [decode_recipe(rec) for rec in recipe_data]
    total = 1
    for (i, recipe) in enumerate(recipes):
        if i > 2:
            break
        root = Node(None, recipe, (0, 0, 0, 0), (1, 0, 0, 0), 32)
        nodes = [root]
        max_geodes = 0
        print("Starting with recipe {0}".format(recipe))
        counter = 0
        best = root
        visited: set[tuple[int, int, int, int, int, int, int, int]] = set()
        while len(nodes) > 0:
            counter += 1
            #nodes.sort(key = lambda n: n.heuristic, reverse=True)
            node = nodes.pop(0)
            visited.add(node.robots + node.stuff)
            #for n in nodes:
            #    if n.heuristic > node.heuristic:
            #        node = n
            #nodes.remove(node)
            if max_geodes < node.get_geodes() or counter % 100000 == 0:
                print("Run ({5}) - max: {0}, nodes: {1}, time: {2}, robots: {3}, stuff: {4}, heuristic: {6}, geodes: {7}, visited: {8}"\
                    .format(max_geodes, len(nodes), node.time_left, node.robots, \
                        node.stuff, counter, node.heuristic, node.get_geodes(), len(visited)))
            #if node.heuristic < max_geodes or node.heuristic == 0:
            if node.time_left < 4 and (node.get_geodes() + node.robots[3] * node.time_left * 2) < max_geodes:
                continue
            if node.get_geodes() > max_geodes:
                best = node
            max_geodes = max(max_geodes, node.get_geodes())
            new_nodes = node.run_step()
            new_nodes.extend(nodes)
            nodes = new_nodes
            nodes = list(filter(lambda n: n.heuristic > 0 and (n.robots + n.stuff) not in visited, nodes))
            #for n in new_nodes:
                #if n.heuristic > max_geodes:
                    #nodes.append(n)
                    #nodes.insert(0, n)
                    #add_to_priority_queue(nodes, n)

        print("Recipe done after {0} rounds".format(counter))
        total *= max_geodes

        while True:
            print(best)
            if best.parent == None:
                break
            best = best.parent

    return total


#result = solve("/Users/kurre/projects/advent-of-code/2022/19/test-input.txt")
#print("Result: ", result)

#assert result == (62 * 56)

result = solve("input.txt")
print("Result: ", result)
