#!/usr/bin/env python3

# DFS
# count paths

class Server:
    def __init__(self, data):
        link = data.split(":")

        self.node = link[0]
        self.next = link[1]
    
    def next(self):
        return self.next

def trace(target, servers):
    count = 0
    if target == "out":
        count += 1
        return count

    for s in servers[target]:
        count += trace(s, servers)
    
    return count

def main():
    # input_file = "input.txt.test"
    # input_file = "input.txt.test2"
    input_file = "input.txt"

    servers = {}
    with open(input_file) as file:
        for line in file:
            clean_in = line.strip().split(":")
            servers[clean_in[0]] = clean_in[1].split(" ")[1:]

    print(f"{servers}")
    print(f"START {servers['you']}")

    total = trace('you', servers)
    print(f"{total}")


if __name__ == "__main__":
    main()