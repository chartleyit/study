#!/usr/bin/env python3

import re

def valid_mul(input):
    cmd_list = re.findall(r"mul\((\d{1,3}),(\d{1,3})\)|(do)\(|(don't)\(", input)
    print(cmd_list)
    return cmd_list

# create something to multiply standard sequenced of mul(X,Y)
def mul(x,y):
    return x * y

# create something to reconcile corrupted instructions
def parse(input, enabled):
    mul_list = valid_mul(input)
    result = 0
    for cmd in mul_list:
        if enabled:
            if cmd[3] != "":
                enabled = False
            else:
                print("enabled")
                print(int(cmd[0] or 0), int(cmd[1] or 0))
                result += mul(int(cmd[0] or 0), int(cmd[1] or 0))
        else:
            if cmd[2] != "":
                enabled = True
            else:
                print("disabled")
                print(cmd)
            
    return result, enabled

def main():
    total = 0
    with open('input.txt') as file:
        enabled = True
        for line in file:
            print(line)
            out, enabled = parse(line, enabled)
            total += out
    print(f"result: {total}")

if __name__ == "__main__":
    main()


# result: 69246921
# too high
# result: result: 69247082
# too high
