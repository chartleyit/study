#!/usr/bin/env python3

def main():
    # input_file = "input.txt.test"
    # input_file = "input.txt.test2"
    input_file = "input.txt"

    file = open(input_file, "r")
    data = file.read()
    file.close()

    lines = data.split("\n")
    matrix = []
    for line in lines:
        matrix.append(line.split())

    for c in range(len(matrix[0])):
        for l in matrix:
            print(f"{l[c]}", end = " ")

            # conditions to handle * or +

        print()



if __name__ == "__main__":
    main()