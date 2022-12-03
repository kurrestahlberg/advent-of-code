with open("input.txt") as file:
    sacks = [x.strip() for x in file.readlines()]
    prioritySum = 0
    for rucksack in sacks:
        first = set(rucksack[:int(len(rucksack) / 2)])
        second = set(rucksack[int(len(rucksack) / 2):])

        same = (first & second).pop()

        if same >= 'a':
            priority = ord(same) - ord('a') + 1
        else:
            priority = ord(same) - ord('A') + 27

        prioritySum += priority

    print(prioritySum)