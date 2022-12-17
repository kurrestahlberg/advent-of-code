from typing import TypeAlias
import re

Sensor: TypeAlias = tuple[int, int, int, int]
Coord: TypeAlias = tuple[int, int]

def map_sensor_data(data: str) -> Sensor:
    res = re.match("Sensor at x=(-?[0-9]+), y=(-?[0-9]+): closest beacon is at x=(-?[0-9]+), y=(-?[0-9]+)", data)
    if res == None:
        print("No match for input in \"{0}\"".format(data))
        exit(666)
    
    return (int(res.group(1))), int(res.group(2)), int(res.group(3)), int(res.group(4))

def sensor_manhattan(sensor: Sensor) -> int:
    return abs(sensor[0] - sensor[2]) + abs(sensor[1] - sensor[3])

def sensor_effect(sensor: Sensor, row: int) -> int:
    distance = abs(sensor[1] - row)
    return sensor_manhattan(sensor) - distance

def solve(file: str, important_row: int) -> int:
    with open(file) as f:
        sensor_data = f.readlines()

    result: set[Coord] = set()
    beacons: set[Coord] = set()
    
    sensors = [map_sensor_data(d) for d in sensor_data]
    for sensor in sensors:
        #print("Sensor: {0}".format(sensor))
        if sensor[3] == important_row:
            beacons.add((sensor[2], important_row))
        effect = sensor_effect(sensor, important_row)
        if effect > 0:
            for x in range(sensor[0] - effect, sensor[0] + effect + 1, 1):
                result.add((x, important_row))

    #print(sorted(result - beacons))
    #print(beacons)
    return len(result - beacons)


result = solve("test-input.txt", 10)
print("Result: {0}".format(result))
assert(result == 26)

result = solve("input.txt", 2000000)
print("Result: {0}".format(result))
