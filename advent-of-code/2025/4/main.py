#!/usr/bin/env python3

NEIGHBOR = [
    [-1, -1],
    [-1,0],
    [0, -1],
    [1, 0],
    [0,1],
    [1,1],
    [1,-1],
    [-1,1]
]

def chk_paper(data):
    accessible = 0

    n = 0
    print(data.split("\n"))
    for line in data.split("\n"):
        print(n, line)
        n+=1

    return accessible

def find_neighbors(data, target):
    rows = len(data)
    cols = len(data[0])
    m = data
    offsets = NEIGHBOR
    total = 0 # is this even working?

    new_m = [[0 for _ in range(rows)] for _ in range(cols)]

    for c in range(0, cols):
        for r in range(0, rows):
            # print(f"STARTING - col {c}, row: {r}")

            count = 0
            n = []
            accessible = True


            # for each position in the matrix
            location = m[r][c]

            # if location is not TP=
            if location != "@":
                # print(f"({r}, {c}): not an @")
                continue

            for dr, dc in offsets:
                ref_row, ref_col = r + dr, c + dc
                # if neighor row is out of bounds skip
                if ref_row < 0 or ref_row >= rows:
                    continue
                # if neighor col is out of bounds skip
                if ref_col < 0 or ref_col >= cols:
                    continue


                # if neighbor space is "@"
                spc = m[ref_row][ref_col]
                n.append(spc)
                if spc == "@":
                    count += 1

                if count >= target:
                    # print(f"({r}, {c}): not accessible")
                    accessible = False
                    continue


            if accessible:
                # print(f"\t{r}, {c} neighbors: {n} - {count}")
                total += 1
                # remove the roll
                new_m[r][c] = "x"
            else:
                new_m[r][c] = m[r][c]

            # print(f"({r}, {c}): accessible")
            # print(f"\t{n}")

    print(f"accessible: {total}")
    return new_m, total

def main():
    # input_file = "input.txt.test"
    # input_file = "input.txt.test2"
    input_file = "input.txt"

    file = open(input_file, "r")
    data = file.read()
    file.close()

    target = 4
    lines = data.split("\n")
    matrix = []
    for line in lines:
        matrix.append(list(line))

    # print_matrix(matrix)
    removed, c = 1, 0
    total = 0
    while removed > 0:
        c += 1
        matrix, removed = find_neighbors(matrix, target)

        total += removed

        if c > 100:
            print(f"ESACPE!!!!")
            break
    
    # print_matrix(matrix)
    print(f"total removed: {total}")

def print_matrix(matrix):
    r = 0
    for l in matrix:
        print(f"{r}: {l}")
        r += 1

if __name__ == "__main__":
    main()