length = 14

with open("input.txt") as f:
    data = f.readline().strip()
    for i in range(len(data)):
        different = len(set(data[i:i+length]))
        if different == length:
            print("start-of-packet: {0}".format(i + length))
            break
