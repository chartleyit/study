#!/usr/bin/env python3

import re

# DFS -> recursion (bad)
# BFS -> breadth first
    # create a graph of data

class Panel:
    def __init__(self, raw_data):
        # col 1 light pattern
        # col 2 wiring
        # col 3 jolt req
        self.lights = re.search(r'\[(.*?)\]', raw_data).group(1)
        self.jolts = re.search(r'\{(.*?)\}', raw_data).group(1).split(",")
        raw_switches = re.findall(r'\((.*?)\)', raw_data)
        self.switches = [[int(x) for x in s.split(",")] for s in raw_switches]

        # self.lights = re.findall(r'(\[.*?\])', raw_data)
        # self.switches = re.findall(r'(\(.*?\))', raw_data)
        # self.jolts = re.findall(r'(\{.*?\})', raw_data)
        self.states = ["."*len(self.lights)]
    
    def __str__(self):
        return f"{self.lights} {self.switches} {self.jolts}"
    
    def lights(self):
        return self.lights
    
    def switches(self):
        return self.switches
    
    def states(self):
        return self.states


def update_lights(panel):
    for s in range(len(panel.states)):
        state = panel.states[s]
        for switch in panel.switches:
            # print(f"DEBUG: switch {switch} {state}")
            split_state = list(state)
            # print(f"{state} push switches -", end = "")
            for action in switch:
                # print(f" {action} ", end = "")
                if split_state[action] == ".":
                    split_state[action] = "#"
                else:
                    split_state[action] = "." 
            # print(f" {''.join(split_state)}")
            new_state = "".join(split_state)
            if new_state not in panel.states:
                # print(f"DEBUG: append new_state {new_state}")
                panel.states.append(new_state)
            if panel.lights in panel.states:
                # print(f"DEBUG: found {panel.lights} in {panel.states}")
                # print(f"UPDATED: {panel.states}")
                return
        # print(f"DEBUG: states {panel.states}")
        

def light_states(panel):
    count = 0
    print(f"STARTING {panel.states} SEARCHING for {panel.lights}")
    print(f"\tSWITCHES: {panel.switches}")
    while panel.lights not in panel.states:
        update_lights(panel)
        count += 1
        if count > 10:
            print(f"\tDEBUG: this is too deep")
            break
        print(f"\tUPDATED: {panel.states}")
    
    return count

def main():
    # input_file = "input.txt.test"
    # input_file = "input.txt.test2"
    input_file = "input.txt"
    
    panels = []
    total = 0

    with open(input_file) as file:
        for line in file:
            clean_in = line.strip()

            panel = Panel(clean_in)
            # print(f"DEBUG: {panel}")
            panels.append(panel)

    for p in panels:
        c = light_states(p)
        print(f"FOUND: in {c}")
        total += c

    print(f"total operations: {total}")
    
    # parse input
    # determine fewest button presses

if __name__ == "__main__":
    main()