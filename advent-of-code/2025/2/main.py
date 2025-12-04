#!/usr/bin/env python3

def valid(id):
    id_str = str(id)
    if id == 0:
        return False

    if len(id_str) % 2 == 0:
        id_split = len(id_str)//2
        # print(f"id_split: {id_split} => {id_str[:id_split]} == {id_str[id_split:]}")
        if id_str[:id_split] == id_str[id_split:]:
            return False
    
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
            chk = valid(id)
            if not chk:
                id_checksum += id
                # print(f"{id} {chk} - updated sum: {id_checksum}")
    
    print(f"invalid id sum: {id_checksum}")

if __name__ == "__main__":
    main()