#!/usr/bin/env python3

# Secret Entrance
# dial with arrow
# 0-99
# puzzel input sequenc of rotations 1 per line
# L or R and distance
# Start 11 > R8 > 19
# 19 > L19 > 0 > L1 > 99
# 99 > R1 > 0

# we could probably uses a fixed length memory allocation of 100 (0-99)
# count the number of times the dial is left pointing at 0

# starting at 50
# L68
# password is number of times it stops at 0

target = 0
start = 50

class Dial:
    DIRECTION = ["L", "R"]
    def __init__(self):
        self.pos = 50
        self.count = 0
        self.new_count = 0
        self.old_count = 0
        self.actions = 0
    
    def turn(self, input):
        self.actions += 1

        direction = input[0]
        distance = int(input[1:])

        if self.pos > 99:
            raise ValueError(f"position can not be greater than 99")
        
        if self.pos < 0:
            raise ValueError(f"position can not be less than 0")

        if direction not in self.DIRECTION:
            raise ValueError(f"direction must be one of {self.DIRECTION} got {direction}")

        print(f"[{self.actions}] start: {self.pos} dir: {direction} change: {distance} \t count: {self.count}")
        self.set(distance, direction)
    
    def pos(self):
        return self.pos
    
    def count(self):
        return self.count
    
    def new_count(self):
        return self.new_count
    
    def old_count(self):
        return self.old_count
    
    def actions(self):
        return self.actions
    
    def set(self, dis, dir):

        # Part 2 count any time we pass 0
        # count the number of times it guaranteed to pass 0
        trips = dis // 100
        self.count += trips
        self.new_count += trips

        if trips > 0:
            print(f"trips {trips}")

        # logic always reduce the action to a less than 100 dist
        final = dis % 100
        if dir == "L":
            self.pos -= final
        if dir == "R":
            self.pos += final

        # we can't end in a position greater than 99; wrap around
        if self.pos > 99:
            print(f"{self.pos} - 100 = {self.pos - 100}")
            # print(f"dir: {dir} change: {dis} adj: {final} times: {trips}")
            self.count += 1
            self.new_count += 1
            self.pos -= 100
        # we cant end in a negative position; wrap around
        if self.pos < 0:
            print(f"{self.pos} + 100 = {self.pos + 100}")
            # print(f"dir: {dir} change: {dis} adj: {final} times: {trips}")
            self.count += 1
            self.new_count += 1
            self.pos += 100
        
        # Part 1 = count when ending position is 0
        if self.pos == 0:
            self.old_count += 1

        # print(f"{self.pos} {self.count}")
    
def main():
    # input_file = "input.txt.test"
    # input_file = "input.txt.test2"
    input_file = "input.txt"
    d = Dial()

    with open(input_file) as file:
        for line in file:
            action = line.strip()
            # print(f"{d.actions}: pos: {d.pos} count: {d.count}")
            d.turn(action)
            # print(f"position {d.pos}")

    # 7815 is too high
    # 4229 is too low
    # 6447 is not right
    # 6700 is someone elses answer
    # 6676 is not right
    print(f"count {d.count}")
    print(f"new count {d.new_count}")
    print(f"old count {d.old_count}")

if __name__ == "__main__":
    main()