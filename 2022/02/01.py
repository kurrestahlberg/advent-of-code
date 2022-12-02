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


with open("input.txt") as file:
    lines = file.readlines()
    score = 0
    for line in lines:
        [them, me] = line.strip().split(" ")
        score += points[them][me]

    print(str(score))