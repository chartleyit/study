#!/usr/bin/env python3

def main():
    # input_file = "input.txt.test"
    # input_file = "input.txt.test2"
    input_file = "input.txt"

    file = open(input_file, "r")
    input = file.read().split("\n")
    file.close()

    # fresh_list, ingredients = parse_input(input)
    # check_fresh(fresh_list, ingredients)

    fresh_list, _ = parse_input(input)
    ranges = build_ranges(fresh_list)

    initial_len = len(ranges)
    reduced_list, depth = recurse_list(ranges[0], ranges[1:])
    update_len = len(reduced_list)

    loop = 0
    while initial_len - update_len > 0:
        loop += 1
        initial_len = update_len
        reduced_list, depth = recurse_list(reduced_list[0], reduced_list[1:])
        update_len = len(reduced_list)
        print(f"{len(reduced_list)}")
        if loop > 100:
            print(f"too many loops")
            break
    print(f"{loop}")

    # 422008276760409 too high
    # 419644187909662 too high
    # 413313071812045 too high
    print(f"\n{len(reduced_list)}")
    print(f"depth {depth}")
    print(f"total ingredients: {count_list(reduced_list)}")

def recurse_list(comp, fresh):
    # we've hit the bottom start to return data
    count = 1
    # print(f"{fresh}")
    if len(fresh) > 1:
        fresh, count = recurse_list(fresh[0], fresh[1:])

    for i in range(len(fresh)):
        # print(f"{comp} {fresh[i]}")
        if fresh[i][0] < comp[0] and comp[1] < fresh[i][1]:
            # print(f"range is within range {comp} {fresh[i]}")
            break

        # if the start is within the range extend the end
        if fresh[i][0] < comp[0] and comp[0] < fresh[i][1]:
            if comp[1] > fresh[i][1]:
                # print(f"extend end range {fresh[i][1]} by {comp[1]}")
                fresh[i][1] = comp[1]
                return fresh, count + 1

        # if the end is within the range extend the start
        if fresh[i][0] < comp[1] and comp[1] < fresh[i][1]:
            if fresh[i][0] > comp[0]:
                # print(f"extend start range {fresh[i][0]} by {comp[0]}")
                fresh[i][0] = comp[0]
                return fresh, count + 1

    # start and end are less then start
    if comp[0] < fresh[i][0] and comp[1] < fresh[i][0]:
        # print(f"range is less than {comp[1]} {fresh[i][0]}")
        fresh.insert(0, comp)
        return fresh, count + 1

    # start and end are greater than end
    if comp[0] > fresh[i][1] and comp[1] > fresh[i][1]:
        # print(f"range is greater than {comp[0]} {fresh[i][1]}")
        fresh.append(comp)
        return fresh, count + 1

    # print(f"fresh: {fresh}")
    return fresh, count + 1

def count_list(data):
    count = 0
    for i in data:
        count += len(range(i[0] - 1, i[1]))

    return count

def check_fresh(fresh_list, ingredients):
    fresh = 0
    spoiled = 0

    # print(f"{fresh_list}")
    # print(f"{ingredients}")
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

def build_ranges(input):
    output = []
    for r in input:
        start = int(r.split("-")[0])
        end = int(r.split("-")[1])
        output.append([start, end])
        # print(f"output {output}")
    
    return output

if __name__ == "__main__":
    main()