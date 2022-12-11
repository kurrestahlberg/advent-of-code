# https://adventofcode.com/2022/day/11
import re

def to_int(input: str) -> int:
    return int(input)

class Monkey:
    def __init__(self, spec: str):
        spec_lines = spec.splitlines()
        self.id = re.match("Monkey ([0-9]*):", spec_lines[0]).group(1)
        self.items = list(map(int, re.match("  Starting items: (([0-9]+,? ?)*)", spec_lines[1]).group(1).split(", ")))
        self.op = re.match("  Operation: new = ([a-zA-Z0-9\*\+ ]*)", spec_lines[2]).group(1)
        self.divisible_by = int(re.match("  Test: divisible by (\d+)", spec_lines[3]).group(1))
        self.if_true = int(re.match("    If true: throw to monkey (\d)+", spec_lines[4]).group(1))
        self.if_false = int(re.match("    If false: throw to monkey (\d)+", spec_lines[5]).group(1))
        self.inspect_count = 0

    def inspect_and_throw(self) -> tuple[int, int] | None:
        if len(self.items) == 0:
            #print("Monkey {0} all out of items!".format(self.id))
            return None

        self.inspect_count += 1

        level = self.items.pop(0)
        #print("Level before is {0}".format(level))
        level = int(int(eval(self.op.replace("old", str(level)))) / 3)
        #print("Level after is {0}".format(level))

        if level % self.divisible_by == 0:
            #print("Throwing to {0}".format(self.if_true))
            return (level, self.if_true)
        else:
            #print("Throwing to {0}".format(self.if_false))
            return (level, self.if_false)

    def add(self, value: int) -> None:
        self.items.append(value)

    def __str__(self) -> str:
        return "Monkey {0} - {1} - {2} - {3} - {4} - {5}".format(self.id, self.op, self.items, self.divisible_by, self.if_true, self.if_false)

with open("input.txt") as f:
    monkeys_specs = f.read().split("\n\n")

    monkeys: list[Monkey] = []
    for monkey_spec in monkeys_specs:
        monkey = Monkey(monkey_spec)
        monkeys.append(monkey)
        print(monkey)

    for i in range(20):
        for monkey in monkeys:
            while True:
                result = monkey.inspect_and_throw()
                if result is None:
                    break

                monkeys[result[1]].add(result[0])
        
        print()
        print("Round {0}".format(i + 1))
        for monkey in monkeys:
            print(monkey.inspect_count, monkey)

        counts = list(map(lambda m: m.inspect_count, monkeys))
        counts.sort(reverse=True)
        print("Level of monkey business is {0}".format(counts[0] * counts[1]))




