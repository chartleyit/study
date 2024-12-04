#!/usr/bin/env python3

import re

def valid_mul(input):
    cmd_list = re.findall(r"mul\((\d{1,3}),(\d{1,3})\)", input)
    print(cmd_list)
    return cmd_list

# create something to multiply standard sequenced of mul(X,Y)
def mul(x,y):
    return x * y

# create something to reconcile corrupted instructions
def parse(input):
    mul_list = valid_mul(input)
    result = 0
    for pair in mul_list:
        result += mul(int(pair[0]), int(pair[1]))
    return result

def main():
    total = 0
    with open('input.txt') as file:
        for line in file:
            print(line)
            total += parse(line)
    print(f"result: {total}")

if __name__ == "__main__":
    main()
