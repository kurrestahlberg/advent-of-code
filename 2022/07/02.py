from __future__ import annotations

class Directory:
    subdirectories: dict[str, Directory]
    total_file_size: int

    def __init__(self, name: str, parent: Directory | None = None):
        self.name = name
        self.parent = parent
        self.total_file_size = 0
        self.subdirectories = {}

    def add_subdirectory(self, subdir: str) -> None:
        self.subdirectories[subdir] = Directory(subdir, self)

    def add_file(self, name: str, size: int) -> None:
        self.total_file_size += size

    def get_total_size(self) -> int:
        total_size = self.total_file_size
        for subdir in self.subdirectories.values():
            subdir_size = subdir.get_total_size()
            total_size += subdir_size

        return total_size

    def get_or_add_subdirectory(self, name: str) -> Directory:
        if self.subdirectories.get(name) != None:
            return self.subdirectories.get(name)

        self.add_subdirectory(dir)
        return dir

    def get_absolute_path(self) -> str:
        path = ""
        if self.parent != None:
            path = self.parent.get_absolute_path()

        if not path.endswith("/"):
            path = path + "/"
        return path + self.name

def get_small_dir_size_sum(start: Directory) -> int:
    total = 0
    print("{0}: {1}".format(start.get_absolute_path(), start.get_total_size()))
    for dir in start.subdirectories.values():
        if dir.get_total_size()  < 100000:
            total += dir.get_total_size()

        total += get_small_dir_size_sum(dir)

    return total

def find_smallest_bigger_than(start: Directory, size: int, smallest: int | None = None) -> int | None:
    start_size = start.get_total_size()
    if start_size < size:
        print("{0} is not big enough ({1}/{2})".format(start.get_absolute_path(), start_size, size))
        return None

    if smallest == None or start_size < smallest:
        print("Currently smallest is {0} at {1}".format(start.get_absolute_path(), start_size))
        smallest = start_size

    for dir in start.subdirectories.values():
        s = find_smallest_bigger_than(dir, size, smallest)
        if s != None:
            if smallest == None or s < smallest:
                smallest = s

    return smallest

with open("input.txt") as f:
    root = Directory("")
    current = root
    lines = [l.strip() for l in f.readlines()]
    for line in lines:
        parts = line.split()
        if parts[0] == "$": # command
            if parts[1] == 'cd':
                if parts[2] == '..':
                    current = current.parent if current.parent != None else root
                elif parts[2] == '/':
                    current = root
                else:
                    current = current.get_or_add_subdirectory(parts[2])
#                print("Switched to {0}".format(current.get_absolute_path()))
#            else:
#                print("Command: {0}".format(parts[1:]))
        elif parts[0] == "dir": # directory
            current.add_subdirectory(parts[1])
        elif parts[0].isnumeric():
            current.add_file(parts[1], int(parts[0]))

    print("Total size: {0}".format(root.get_total_size()))
    print("Smalls size: {0}".format(get_small_dir_size_sum(root)))
    spaceNeeded = root.get_total_size() - (70000000 - 30000000)
    print("Space needed: {0}".format(spaceNeeded))
    print("Smallest big enough size: {0}".format(find_smallest_bigger_than(root, spaceNeeded)))


