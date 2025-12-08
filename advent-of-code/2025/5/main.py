#!/usr/bin/env python3

def main():
    # input_file = "input.txt.test"
    # input_file = "input.txt.test2"
    input_file = "input.txt"

    file = open(input_file, "r")
    input = file.read().split("\n")
    file.close()

    fresh_list, ingredients = parse_input(input)

    fresh = 0
    spoiled = 0

    print(f"{fresh_list}")
    print(f"{ingredients}")
    # evalutate ingredients in lists
    for i in ingredients:
        for l in fresh_list:
            # could do with split comparison
            start = l.split("-")[0]
            end = l.split("-")[1]
            if i >= int(start) and i <= int(end):
                # print(f"fresh - {i}")
                fresh += 1
                break
            else:
                # print(f"spoiled - {i}")
                spoiled += 1
            # could do a contains

    print(f"fresh {fresh}")

def parse_input(input):
    fresh_list = []
    ingredients = []
    part = 0

    for line in input:
        if line == "":
            part = 1
            continue

        if part == 0:
            fresh_list.append(line)
    
        if part == 1:
            ingredients.append(int(line))
        
    return fresh_list, ingredients

if __name__ == "__main__":
    main()