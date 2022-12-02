# https://adventofcode.com/2022/day/2

points = {
    "A": {
        "X": 4,
        "Y": 8,
        "Z": 3
    },
    "B": {
        "X": 1,
        "Y": 5,
        "Z": 9
    },
    "C": {
        "X": 7,
        "Y": 2,
        "Z": 6
    },
}

resultmap = {
    "A": {
        "X": "Z",
        "Y": "X",
        "Z": "Y"
    },
    "B": {
        "X": "X",
        "Y": "Y",
        "Z": "Z"
    },
    "C": {
        "X": "Y",
        "Y": "Z",
        "Z": "X"
    },
}


with open("input.txt") as file:
    lines = file.readlines()
    score = 0
    for line in lines:
        [them, me] = line.strip().split(" ")
        score += points[them][resultmap[them][me]]

    print(str(score))