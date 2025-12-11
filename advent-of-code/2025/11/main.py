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

# pt2
def trace_svr(target, stops, servers):
    count = 0
    if target == "out" and 'dac' in stops and 'fft' in stops:
        stops.append(target)
        print(f"DEBUG: {stops}")
        count += 1
        return count

    if target == "out":
        stops.append(target)
        # print(f"reached out with stops: {stops}")
        return count
    
    if len(stops) > len(servers) * 2:
        print(f"reached limit with stops {stops}")
        return count
    
    stops.append(target)
    for s in servers[target]:
        count += trace_svr(s, stops, servers)
    
    return count

def main():
    input_file = "input.txt.test"
    # input_file = "input.txt.test2"
    # input_file = "input.txt"

    servers = {}
    with open(input_file) as file:
        for line in file:
            clean_in = line.strip().split(":")
            servers[clean_in[0]] = clean_in[1].split(" ")[1:]

    print(f"{servers}")
    print(f"START {servers['you']}")

    total = trace('you', servers)
    
    # pt2
    svrs = {}
    input2_file = "input2.txt"
    # input2_file = "input.txt"

    with open(input2_file) as file:
        for line in file:
            clean_in = line.strip().split(":")
            svrs[clean_in[0]] = clean_in[1].split(" ")[1:]

    total_svr = trace_svr('svr', [], svrs)

    print(f"PT1: total from you {total}")
    print(f"PT2: total from svr with stops: {total_svr}")


if __name__ == "__main__":
    main()