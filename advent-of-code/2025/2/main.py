#!/usr/bin/env python3

def repeat(id_str, numb):
    # n, x = 0, len(id_str)//numb
    n, x = 0, numb
    while n < len(id_str) - x:
        current = id_str[n:n+x]
        next = id_str[n+x:n + (2*x)]
        # as soon as we hit a pair that isn't a match it returns true
        if current != next:
            return True
        n += x
    
    # return false if we go through the set and don't return true
    # print(f"repeat: {id_str}")
    return False

def valid(id):
    id_str = str(id)

    # any numbers starting with 0 aren't valid
    if id_str[0] == "0":
        return False

    # any number that isn't greater than 1 can't repeat 
    if len(id_str) == 1:
        return True

    # * quick test if any string is just repeated single characters
    if id_str == id_str[0] * len(id_str):
        print(f"{id_str}")
        return False

    if len(id_str) % 2 == 0:
        id_split = len(id_str)//2
        # print(f"id_split: {id_split} => {id_str[:id_split]} == {id_str[id_split:]}")
        # think recursion; split each time and see if it repeats
        if id_str[:id_split] == id_str[id_split:]:
            print(f"{id_str}")
            return False
    
    # for the list of values we can subdivide the list into
    sets = range(2, len(id_str)//2 )
    for n in reversed(sets):
        # print(f"{id_str} starting at {n}")
        if len(id_str) % n == 0:
            # this is excessive but for time just make it work
            if not repeat(id_str, n):
                print(f"{id_str}")
                return False
    
    # if len(id_str) % 3 == 0:
    #     return repeat(id_str, 3)
    
    return True


def parse(ids):
    id_list = ids.strip().split(",")
    return id_list

def main():
    # input_file = "input.txt.test"
    # input_file = "input.txt.test2"
    input_file = "input.txt"

    file = open(input_file, "r")
    id_ranges = parse(file.read())
    file.close()

    id_checksum = 0
    for id_range in id_ranges:
        start, end = int(id_range.split("-")[0]), int(id_range.split("-")[1])
        # print(f"id_range: {start}, {end}")
        for id in range(start, end + 1):
            # print(f"{id}:")
            chk = valid(id)
            if not chk:
                id_checksum += id
                # print(f"{id} {chk} - updated sum: {id_checksum}")
    
    # 1227779271 is too low
    # 46270373637 is too high
    # 46270373637 is too high
    print(f"invalid id sum: {id_checksum}")

if __name__ == "__main__":
    main()