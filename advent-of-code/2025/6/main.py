#!/usr/bin/env python3

# part 2 is tricky; parser is annoying and needs to perserve spaces

def read_ceph(data):
    ceph_data = []
    for line in data:
        ceph_data.insert(0,line)

    # print(f"{ceph_data} {len(data[0])}")

    sum = 0
    numbers = []
    for col in reversed(range(len(data[0]))):
        number = ""
        op = ""
        for line in data:
            chr = line[col]
            if chr == " ":
                continue
            
            if chr != "*" and chr != "+":
                number += chr
            
            if chr == "*" or chr == "+":
                op = chr

        # print(f"{number}")

        if number != "":
            # this ends with the right number
            numbers.append(int(number))
    
        if op != "":
            total = 0
            for n in numbers:
                if op == "+":
                    total += n
                if op == "*":
                    if total == 0:
                        total = 1
                    total *= n
            sum += total
            # print(f"{numbers}")
            # print(f"operation {op}")
            # print(f"{total}")
            op = ""
            numbers = []

    return sum

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

    print(f"{read_ceph(lines)}")

if __name__ == "__main__":
    main()