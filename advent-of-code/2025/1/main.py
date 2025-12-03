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
    
    def turn(self, input):
        direction = input[0]
        distance = int(input[1:])

        if direction not in self.DIRECTION:
            raise ValueError(f"direction must be one of {self.DIRECTION} got {direction}")

        if direction == "L":
            self.set(-abs(distance))
        if direction == "R":
            self.set(distance)
        
        if self.pos == 0:
            self.count += 1 
        
        # this logic needs work
        if self.pos > 99:
            self.pos -= 99
        if self.pos < 0:
            self.pos += 99
    
    def pos(self):
        return self.pos
    
    def count(self):
        return self.count
    
    def set(self, d):
        self.pos += d
        while self.pos > 99:
            self.pos -= 100
        while self.pos < 0:
            self.pos += 100
    
def main():
    d = Dial()
    with open('input.txt') as file:
        for line in file:
            action = line.strip()
            d.turn(action)
            print(f"position {d.pos}")

    print(f"count {d.count}")

if __name__ == "__main__":
    main()