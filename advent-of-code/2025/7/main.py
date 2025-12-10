#!/usr/bin/env python3

def scan(input):
    matrix = input
    splits = 0
    for line in range(len(matrix) - 1):
        for c in range(len(matrix[line])):
            col = matrix[line][c]
            if col == "S" or col == "|":
                if line < len(matrix):
                    if matrix[line+1][c] == ".":
                        matrix[line+1][c] = "|"
                    if matrix[line+1][c] == "^":
                        splits += 1
                        if c > 0:
                            matrix[line+1][c-1] = "|"
                        if c < len(matrix[line]):
                            matrix[line+1][c+1] = "|"
        print(f"{matrix[line]}")
                    

    return splits

def main():
    # input_file = "input.txt.test"
    # input_file = "input.txt.test2"
    input_file = "input.txt"

    lines = []
    with open(input_file) as file:
        for line in file:
            lines.append(list(line.strip()))

    splits = scan(lines)
    print(f"{splits}")

if __name__ == "__main__":
    main()