#!/usr/bin/env python3

# DFS -> recursion (bad)
# BFS -> breadth first
    # create a graph of data

class Panel:
    def __init__(self, raw_data):
        fields = raw_data.split(" ")
        self.lights = fields[0]
        self.switches = fields[1:len(fields)-2]
        self.jolts = fields[len(fields)-1]
        self.states = []
    
    def __str__(self):
        return f"{self.lights} {self.switches} {self.jolts}"

def update_lights(panel):
    return

def main():
    input_file = "input.txt.test"
    # input_file = "input.txt.test2"
    # input_file = "input.txt"
    
    panels = []

    with open(input_file) as file:
        for line in file:
            clean_in = line.strip()

            panel = Panel(clean_in)
            print(f"{panel}")
            panels.append(panel)
    
    # parse input
    # col 1 light pattern
    # col 2 wiring
    # col 3 jolt req
    # determine fewest button presses

    file = open(input_file, "r")
    data = file.read()
    file.close()
    lines = data.split("\n")


if __name__ == "__main__":
    main()