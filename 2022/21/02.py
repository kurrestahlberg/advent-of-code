from __future__ import annotations
import re

class Monkey:
    def __init__(self, name: str, action: str, monkeys: dict[str, Monkey]) -> None:
        action = action.strip()
        self.name = name.strip()
        self.monkeys = monkeys
        if action.isnumeric():
            self.value = int(action)
            self.action = None
        else:
            self.value = None
            m = re.match("([a-z]{4}) ([\+\-\*\/]) ([a-z]{4})", action)
            if m != None:
                self.action = m.groups()
            else:
                print("FAILURE: \"{0}\"".format(action))

    def get_value(self) -> float:
        if self.value != None:
            return self.value

        assert self.action != None

        if self.name == "root":
            left = self.monkeys[self.action[0]].get_value()
            right = self.monkeys[self.action[2]].get_value()

            #print("Left: {0}, right: {1}".format(left, right))

            return left - right

        match self.action[1]:
            case "+":
                return self.monkeys[self.action[0]].get_value() + self.monkeys[self.action[2]].get_value()
            case "-":
                return self.monkeys[self.action[0]].get_value() - self.monkeys[self.action[2]].get_value()
            case "/":
                return self.monkeys[self.action[0]].get_value() / self.monkeys[self.action[2]].get_value()
            case "*":
                return self.monkeys[self.action[0]].get_value() * self.monkeys[self.action[2]].get_value()
            case _: return 0

def solve(file: str) -> int:
    monkeys: dict[str, Monkey] = {}

    with open(file) as f:
        data = f.readlines()

    for d in data:
        [name, value] = d.split(": ")
        monkeys[name] = Monkey(name, value, monkeys)

    value = 0
    for _ in range(1000):
        monkeys["humn"].value = round(value)
        first = monkeys["root"].get_value()
        if first == 0:
            return round(value)
        monkeys["humn"].value = round(value + 1)
        second = monkeys["root"].get_value()

        print(value, first, second)

        value += first / (first - second)


#result = solve("test-input.txt")
#print("Result", result)

#assert result == 301

result = solve("input.txt")
print("Result", result)
