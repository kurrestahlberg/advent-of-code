with open("input.txt") as file:
    sacks = [x.strip() for x in file.readlines()]
    prioritySum = 0
    for i in range(0, len(sacks), 3):
        group = [set(sack.strip()) for sack in sacks[i:i+3]]
        same = (group[0] & group[1] & group[2]).pop()

        if same >= 'a':
            priority = ord(same) - ord('a') + 1
        else:
            priority = ord(same) - ord('A') + 27

        prioritySum += priority

    print(prioritySum)