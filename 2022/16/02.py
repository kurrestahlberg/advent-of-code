from __future__ import annotations

#from typing import TypeAlias
import re

#Valve: TypeAlias = tuple[int, int, int, int]
#Coord: TypeAlias = tuple[int, int]

class Valve:
    def __init__(self, data: str) -> None:
        m = re.match("Valve ([A-Z]{2}) has flow rate=([0-9]+); tunnels? leads? to valves? (([A-Z]{2},? ?)+)", data)
        if m == None:
            print("Failed to parse \"{0}\"".format(data))
            exit(666)
        self.id = m.group(1)
        self.flow = int(m.group(2))
        self.path_ids: list[str] = m.group(3).split(", ")
        self.paths: list[Valve] = []

    def __str__(self) -> str:
        return "Valve {0} with flow {1} and tunnels to {2}".format(self.id, self.flow, [p.id for p in self.paths])

    def resolve_paths(self, valves: dict[str, Valve]) -> None:
        for path_id in self.path_ids:
            self.paths.append(valves[path_id])

def find_max_left(valve_list: list[Valve], time_left: int, open_valves: list[str]) -> int:
    valve_list.sort(key=lambda l: l.flow if l.id not in open_valves else 0, reverse=True)
    total = 0
    for (i, valve) in enumerate(valve_list):
        time_left -= (2 if i%2 == 0 else 0)
        if time_left <= 0:
            return total
        total += valve.flow * time_left
    return total

#def traverse(valve_list: list[Valve], valves: list[Valve], flow: int, time_left: int, best: int, open_valves: list[str], path_to_here: list[list[str]]) -> int:
def traverse(valve_list: list[Valve], valves: list[Valve], flow: int, time_left: int, best: int, open_valves: list[str], path_to_here: list[list[str]]) -> int:
    path_to_here.append([valves[0].id, valves[1].id])
    if time_left <= 0:
        return max(flow, best)

    max_left = find_max_left(valve_list, time_left, open_valves)
    if max_left == 0 or max_left + flow < best:
        return max(flow, best)

    valves[0].paths.sort(key=lambda l: l.flow if l.id not in open_valves else 0, reverse=True)

    if valves[0].id not in open_valves and valves[0].flow > 0:
        open_valves_copy = open_valves.copy()
        open_valves_copy.append(valves[0].id)
        open_valves_copy.sort()
        new_flow = flow + ((time_left - 1) * valves[0].flow)

        if valves[1].id not in open_valves_copy and valves[1].flow > 0:
            open_valves_copy_2 = open_valves_copy.copy()
            open_valves_copy_2.append(valves[1].id)
            open_valves_copy_2.sort()
            new_flow_2 = new_flow + ((time_left - 1) * valves[1].flow)

            res = traverse(valve_list, valves, new_flow_2, time_left - 1, best, open_valves_copy_2, path_to_here.copy())
            if res > best:
                best = res
        for path in valves[1].paths:
            if path.id == valves[0].id:
                continue
            valves_copy = valves.copy()
            valves_copy[1] = path
            #if best > 0:
                #print("Checking {0}->{1}, time left: {2}, flow: {3}, best: {4}, max_left: {5}".format(valve.id, path.id, time_left, flow, best, max_left))
            res = traverse(valve_list, valves_copy, new_flow, time_left - 1, best, open_valves_copy.copy(), path_to_here.copy())
            if res > best:
                best = res
                print("{4}: {0} {3}s left ({1} & {2}) {5}".format(best, new_flow, max_left, time_left, path_to_here[-1], open_valves))
                print(path_to_here)
    for path in valves[0].paths:
        if path.id == valves[1].id:
            continue

        valves_copy = valves.copy()
        valves_copy[0] = path

        if valves[1].id not in open_valves and valves[1].flow > 0:
            open_valves_copy_2 = open_valves.copy()
            open_valves_copy_2.append(valves[1].id)
            open_valves_copy_2.sort()
            new_flow_2 = flow + ((time_left - 1) * valves[1].flow)

            res = traverse(valve_list, valves_copy, new_flow_2, time_left - 1, best, open_valves_copy_2, path_to_here.copy())
            if res > best:
                best = res
        for path_2 in valves[1].paths:
            if path_2.id == valves_copy[0].id:
                continue
            valves_copy_2 = valves_copy.copy()
            valves_copy_2[1] = path_2
            #if best > 0:
                #print("Checking {0}->{1}, time left: {2}, flow: {3}, best: {4}, max_left: {5}".format(valve.id, path.id, time_left, flow, best, max_left))
            res = traverse(valve_list, valves_copy_2, flow, time_left - 1, best, open_valves.copy(), path_to_here.copy())
            if res > best:
                best = res
                print("{4}: {0} {3}s left ({1} & {2}) {5}".format(best, flow, max_left, time_left, path_to_here[-1], open_valves))
                print(path_to_here)

    return best

def solve(file: str) -> int:
    with open(file) as f:
        valvedata = f.readlines()

    valve_list = [Valve(v) for v in valvedata]
    valves: dict[str, Valve] = {}
    for valve in valve_list:
        valves[valve.id] = valve

    for valve in valve_list:
        valve.resolve_paths(valves)
        print(valve)

    best = traverse(valve_list, [valves['AA'], valves['AA']], 0, 26, 0, [], [])
    
    return best

res = solve("/Users/kurre/projects/advent-of-code/2022/16/test-input.txt")
print("Result = {0}".format(res))
assert(res == 1707)

#res = solve("/Users/kurre/projects/advent-of-code/2022/16/input.txt")
#print("Result = {0}".format(res))
