
def get_range(range: str):
    return [int(value) for value in range.split("-")]

with open("input.txt") as file:
    lines = [l.strip() for l in file.readlines()]
    overlap_count = 0
    for line in lines:
        [elf1, elf2] = [get_range(elf) for elf in line.split(",")]
        if elf1[0] >= elf2[0] and elf1[1] <= elf2[1]:
            overlap_count += 1
        elif elf1[0] <= elf2[0] and elf1[1] >= elf2[1]:
            overlap_count += 1

    print(overlap_count)