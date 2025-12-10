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
    total = 0
    for line in lines:
        matrix.append(line.split())

    for c in range(len(matrix[0])):
        last = len(matrix) - 1
        output = ""
        if matrix[last][c] == "*":
            op = "*"
            result = 1
        if matrix[last][c] == "+":
            op = "+"
            result = 0

        for line in range(len(matrix)):
            l = matrix[line]
            if line != last:
                output += l[c] + " "
                # print(f"{l[c]}", end = " ")

            if op == "*" and line < last:
                # result *= int(l[c])
                result = result * int(l[c])
                output += "* "
            # conditions to handle * or +

            if op == "+" and line < last:
                result += int(l[c])
                output += "+ "

            if line == last:
                output += "="
                # print(f"=", end = "")

        total += result
        print(output, result)
        print(total)
    
    print(f"{len(matrix)}")


if __name__ == "__main__":
    main()