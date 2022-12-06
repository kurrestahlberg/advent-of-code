with open("input.txt") as f:
    data = f.readline().strip()
    for i in range(len(data)):
        different = len(set(data[i:i+4]))
        if different == 4:
            print("start-of-packet: {0}".format(i + 4))
            break
