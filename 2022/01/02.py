# https://adventofcode.com/2022/day/1

with open("input.txt") as file:
    input = file.read()

    totals: list[int] = []
    count = 0

    inputByElf = input.split("\n\n")

    for elf in inputByElf:
        items = elf.split("\n")
        totals.append(0)
        for item in items:
            totals[count] += int(item)
        count += 1

    totals.sort(reverse = True)

    print(totals[0])
    print(totals[1])
    print(totals[2])

    print(totals[0] + totals[1] + totals[2])