with open("input.txt") as file:
    sacks = [x.strip() for x in file.readlines()]
    prioritySum = 0
    for i in range(int(len(sacks)/3)):
        group = [sack.strip() for sack in sacks[i*3:(i*3)+3]]
        #print(group)
        group = [set(sack) for sack in group]
        #print(group[0] & group[1])
        #print(group[1] & group[2])
        #print(group[0] & group[2])
        same = (group[0] & group[1] & group[2]).pop()
        #print("{0}: {1}".format(i, same))

        if same >= 'a':
            priority = ord(same) - ord('a') + 1
        else:
            priority = ord(same) - ord('A') + 27

        prioritySum += priority

        #print("First: {0}, second: {1}, same: {2}, prio: {3}".format(first, second, same, priority))

    print(prioritySum)