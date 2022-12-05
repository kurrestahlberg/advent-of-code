import math
import re

def add_stack_layer(stacks: list[list[str]], line: str):
    stackItems = [line[i: i + 4] for i in range(0, len(line), 4)]
    if stackItems[0][1].isdigit():
        return False
    else:
        for i in range(len(stackItems)):
            crate = stackItems[i][1]
            if crate.isalpha():
                stacks[i].insert(0, crate)
    
    return True

def move_crates(stacks: list[list[str]], line: str):
    [quantity, source, destination] = re.match("move ([0-9]+) from ([0-9]+) to ([0-9]+)", line).groups()

    for i in range(int(quantity)):
        crate = stacks[int(source) - 1].pop()
        stacks[int(destination) - 1].append(crate)


with open("input.txt") as f:
    lines = f.readlines()
    stacking = True
    stacks: list[list[str]] = []
    for i in range(math.ceil(len(lines[0]) / 4)):
        stacks.append([])
    for line in lines:
        if len(line.strip()) == 0:
            continue

        if stacking:
            stacking = add_stack_layer(stacks, line)
        else:
            move_crates(stacks, line)

    result = []
    for stack in stacks:
        result.append(stack.pop())

    print("".join(result))