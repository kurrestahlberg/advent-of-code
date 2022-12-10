# https://adventofcode.com/2022/day/10

with open("input.txt") as f:
    cmds = [cmd.strip() for cmd in f.readlines()]
    x: list[int] = [1]
    active_op: list[str] = ["init"]
    current = 1

    for cmd in cmds:
        op = cmd.split(" ")

        if op[0] == "noop":
            x.append(current)
            active_op.append(cmd)
        elif op[0] == "addx" and len(op) == 2:
            x.append(current)
            active_op.append(cmd)
            current += int(op[1])
            x.append(current)
            active_op.append(cmd)
        else:
            print("WTF?!", cmd, len(x))
            exit(1)

    interesting_cycles = [20, 60, 100, 140, 180, 220]
    print("Total number of cycles: {0}, x at 20: {1}".format(len(x), x[20]))

    total = 0
    for cycle in interesting_cycles:
        print("X at cycle {0} is {1} and signal strength is {2}".format(cycle, x[cycle - 1], x[cycle - 1] * cycle))
        total += x[cycle - 1] * cycle

    for i, val in enumerate(x):
        print(i, "\t", val, "\t", active_op[i])

    print("Total:", total)


