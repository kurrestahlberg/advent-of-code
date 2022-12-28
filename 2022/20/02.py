
class Value:
    def __init__(self, value: int) -> None:
        self.value = value

    def __str__(self) -> str:
        return str(self.value)

def solve(file: str) -> int:
    values: list[Value] = []
    key = 811589153

    zero_pos = -1
    
    with open(file) as f:
        for val in f.readlines():
            int_val = int(val.strip()) * key
            if int_val == 0:
                zero_pos = len(values)
            values.append(Value(int_val))

    print("Zero at", zero_pos, values[zero_pos])

    mixed = values.copy()

    for i in range(10):
        for value in values:
            idx = mixed.index(value)
            mixed.remove(value)
            new_pos = idx + value.value
            new_pos %= (len(values) - 1)
            mixed.insert(new_pos, value)
            #print("Moved {0} from {1} to {2}".format(value.value, idx, new_pos))
        print("Round {0} done".format(i + 1))
        print(list(map(str, mixed)))


    zero_idx = mixed.index(values[zero_pos])
    result = 0
    for i in range(1000, 3001, 1000):
        add = mixed[(zero_idx + i)%len(mixed)].value
        print("Adding {0} to {1}".format(add, result))
        result += add

    print("---")
    print(list(map(str, values)))
    print(list(map(str, mixed)))
    return result

result = solve("test-input.txt")
print("Result:", result)

assert result == 1623178306

result = solve("input.txt")
print("Result:", result)
