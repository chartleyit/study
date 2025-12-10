#!/usr/bin/env python3

def main():
    # input_file = "input.txt.test"
    # input_file = "input.txt.test2"
    input_file = "input.txt"

    with open(input_file) as file:
        for line in file:
            clean_in = line.strip()

    file = open(input_file, "r")
    data = file.read()
    file.close()
    lines = data.split("\n")


if __name__ == "__main__":
    main()