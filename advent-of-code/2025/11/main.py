#!/usr/bin/env python3

# DFS
# count paths

class Server:
    def __init__(self):
        self.adj = {}
        self.num_nodes = 0

    def add_edge(self, src, dst):
        neighbors = self.adj.setdefault(src, set())

        neighbors.add(dst)

        self.adj.setdefault(dst, set())

    def display(self):
        for node, neighbors in sorted(self.adj.items()):
            print(f"{node} -> {list(neighbors)}")
    
    def count_paths(self, start, target, visit):
        targets = set(target)
        visits = set(visit)

        return self._dfs_recusrive_count(start, targets, visits, visited={})
    
    def _dfs_recusrive_count(self, current, targets, visits, visited):
        # print(f"{current}, {targets}, {visited}")
        if current in targets:
            return 1
        
        if current in visited:
            return 0
        
        visited[current] = True

        total_paths = 0

        neighbors = self.adj.get(current, set())

        for neighbor in neighbors:
            total_paths += self._dfs_recusrive_count(neighbor, targets, visits, visited)
        
        del visited[current]

        return total_paths

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
    # stops.append(target)
    # if target == "out":
    #     if 'dac' in stops and 'fft' in stops:
    #         print(f"DEBUG: {stops}")
    #         stops = []
    #         count += 1
    #         return count
        
    #     stops = []
    #     return count
    
    if len(stops) > len(servers) * 4:
        print(f"reached limit with stops {stops}")
        return count
    
    for s in servers[target]:
        stops.append(s)
        if s == "out":
            if 'dac' in stops and 'fft' in stops:
                count += 1
            return count
        count += trace_svr(s, stops, servers)
    
    return count

def main():
    input_file = "input.txt.test"
    # input_file = "input.txt.test2"
    # input_file = "input.txt"

    servers = {}
    servers_graph = Server()
    with open(input_file) as file:
        for line in file:
            clean_in = line.strip().split(":")
            servers[clean_in[0]] = clean_in[1].split(" ")[1:]
            for d in clean_in[1].split(" ")[1:]:
                servers_graph.add_edge(clean_in[0], d)

    print(f"{servers}")
    print(f"START {servers['you']}")

    total = trace('you', servers)
    
    # pt2
    svrs = {}
    input2_file = "input2.txt"
    # input2_file = "input.txt"

    graph = Server()
    with open(input2_file) as file:
        for line in file:
            clean_in = line.strip().split(":")
            # graph.add_edge(clean_in[0], clean_in[1].split(" ")[1:])
            for d in clean_in[1].split(" ")[1:]:
                graph.add_edge(clean_in[0], d)
            # svrs[clean_in[0]] = clean_in[1].split(" ")[1:]

    # total_svr = trace_svr('svr', [], svrs)
    # graph.display()

    svr_total = graph.count_paths('svr', ['out'], ['dac', 'fft'])

    test = servers_graph.count_paths('you', ['out'], ['dac', 'fft'])
    print(f"PT1v2: test base case with graph {test}")

    print(f"PT1: total from you {total}")
    # print(f"PT2: total from svr with stops: {total_svr}")
    print(f"No infinite loop paths {svr_total}")



if __name__ == "__main__":
    main()