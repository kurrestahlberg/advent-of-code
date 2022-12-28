
class Value:
    def __init__(self, value: int) -> None:
        self.value = value

    def __str__(self) -> str:
        return str(self.value)

def solve(file: str) -> int:
    values: list[Value] = []

    zero_pos = -1
    
    with open(file) as f:
        for val in f.readlines():
            int_val = int(val.strip())
            if int_val == 0:
                zero_pos = len(values)
            values.append(Value(int_val))

    print("Zero at", zero_pos, values[zero_pos])

    mixed = values.copy()

    for value in values:
        idx = mixed.index(value)
        mixed.remove(value)
        new_pos = idx + value.value
        while new_pos < 0:
            new_pos += len(values) - 1
        while new_pos > len(values) - 1:
            new_pos -= len(values) - 1
        mixed.insert(new_pos, value)

        #print("Moved {0} from {1} to {2}".format(value.value, idx, new_pos))
        #print(list(map(str, mixed)))

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

assert result == 3

result = solve("input.txt")
print("Result:", result)
