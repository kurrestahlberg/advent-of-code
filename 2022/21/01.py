from __future__ import annotations
import re

class Monkey:
    def __init__(self, action: str, monkeys: dict[str, Monkey]) -> None:
        action = action.strip()
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

    def get_value(self) -> int:
        if self.value != None:
            return self.value

        assert self.action != None

        match self.action[1]:
            case "+":
                return self.monkeys[self.action[0]].get_value() + self.monkeys[self.action[2]].get_value()
            case "-":
                return self.monkeys[self.action[0]].get_value() - self.monkeys[self.action[2]].get_value()
            case "/":
                return int(self.monkeys[self.action[0]].get_value() / self.monkeys[self.action[2]].get_value())
            case "*":
                return self.monkeys[self.action[0]].get_value() * self.monkeys[self.action[2]].get_value()
            case _: return 0

def solve(file: str) -> int:
    monkeys: dict[str, Monkey] = {}

    with open(file) as f:
        data = f.readlines()

    for d in data:
        [name, value] = d.split(": ")
        monkeys[name] = Monkey(value, monkeys)

    return monkeys["root"].get_value()


result = solve("test-input.txt")
print("Result", result)

assert result == 152

result = solve("input.txt")
print("Result", result)
