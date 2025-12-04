#!/usr/bin/env python3


def max_jolts(bank):
    x, y = 0, 1
    i = 1
    # minus 1 because we can't set i to the last digit
    while i < len(bank)-1:
        # increase the tens
        if int(bank[x]) < int(bank[i]):
            x, y = i, i+1
        # increase the ones
        if int(bank[y]) < int(bank[i+1]):
            y = i + 1
        i += 1
        

    result = int("".join([bank[x], bank[y]]))
    print(f"{bank[x:]} - {result}")
    return result

def main():
    # input_file = "input.txt.test"
    # input_file = "input.txt.test2"
    input_file = "input.txt"

    total_jolts = 0

    # file = open(input_file, "r")
    # input = file.read()
    # file.close()
    # print(f"{input}")

    with open(input_file) as file:
        line_num = 1
        for line in file:
            clean_in = line.strip()
            print(f"{line_num}: ", end='')
            total_jolts += max_jolts(clean_in)
            line_num += 1

    # 16625 is too low
    print(f"total_jolts: {total_jolts}")

if __name__ == "__main__":
    main()