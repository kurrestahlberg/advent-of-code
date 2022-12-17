from typing import TypeAlias
import re

Sensor: TypeAlias = tuple[int, int, int, int, int]
Coord: TypeAlias = tuple[int, int]

def manhattan(x1: int, y1: int, x2: int, y2: int) -> int:
    return abs(x1 - x2) + abs(y1 - y2)

def map_sensor_data(data: str) -> Sensor:
    res = re.match("Sensor at x=(-?[0-9]+), y=(-?[0-9]+): closest beacon is at x=(-?[0-9]+), y=(-?[0-9]+)", data)
    if res == None:
        print("No match for input in \"{0}\"".format(data))
        exit(666)
    
    return (int(res.group(1)), int(res.group(2)), int(res.group(3)), int(res.group(4)), manhattan(int(res.group(1)), int(res.group(2)), int(res.group(3)), int(res.group(4))))

def check_point(sensors: list[Sensor], y: int, x: int, size: int):
    for sensor in sensors:
        m = manhattan(x, y, sensor[0], sensor[1])
        if m <= sensor[4]:
            #if y == 11:
            #print("Manhattan {0},{1} {2},{3} {4} <= {5}".format(x, y, sensor[0], sensor[1], m, sensor[4]))
            if x < sensor[0]:
                #print("Moving cursor: {0} {1} {2}".format(abs(sensor[0] - x), sensor[4], abs(sensor[1] - y)))
                return x + (abs(sensor[0] - x) + (sensor[4] - abs(sensor[1] - y))) + 1

            #print("Moving cursor 2: {0} {1} {2}".format(sensor[4], abs(sensor[1] - y), abs(sensor[0] - x)))
            return x + (sensor[4] - abs(sensor[1] - y)) - abs(sensor[0] - x) + 1
    return x

def check_row(sensors: list[Sensor], row: int, size: int) -> int:
    x: int = 0
    while True:
        #print("Checking point ({0},{1})".format(x, row))
        result = check_point(sensors, row, x, size)
        #print("Result of check is {0}".format(result))
        if result == x:
            return x
        if result > size:
            return -1
        x = result

def solve(file: str, size: int) -> int:
    with open(file) as f:
        sensor_data = f.readlines()

    sensors = [map_sensor_data(d) for d in sensor_data]

    for y in range(size):
        #if y % 100000 == 0:
        #    print("Checking row {0}".format(y))
        x = check_row(sensors, y, size)
        if x >= 0:
            print("Result found in ({0},{1})".format(x, y))
            return x * 4000000 + y

    return -1

#result = solve("test-input.txt", 20)
#print("Result: {0}".format(result))
#assert(result == 56000011)

result = solve("input.txt", 4000000)
print("Result: {0}".format(result))
