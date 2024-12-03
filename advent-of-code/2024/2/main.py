#!/usr/bin/env python3

def dumb(l):
    safe = 0
    for i in range(len(l)):
        if i + 1 == len(l):
            print("safe")
            safe += 1
            break
        elif abs(int(l[i]) - int(l[i+1])) > 3:
            print(abs(int(l[i]) - int(l[i+1])))
            print("unsafe")
            break
        elif l[i] == l[i+1]:
            print(abs(int(l[i]) - int(l[i+1])))
            print("unsafe")
            break
    
    return safe

def test(a, b, prev):
    change = a - b
    if change == 0:
        # print("unsafe: no change")
        return 0, change
        
    if abs(change) > 3:
        # print("unsafe: change greater than 3")
        return 0, change

    if change < 0 and prev > 0:
        # print("unsafe: decreasing change")
        return 0, change

    if change > 0 and prev < 0:
        # print("unsafe: increasing change")
        return 0, change
    
    return 1, change


def check(lvls):
    # checking to see if we have a sorted list that
    # has elements that don't increase by more than 3

    # if last diff was opposite sign
    change = 0
    prev_change = 0
    prev = 0
    safe = 0
    tolerate = 0
    
    lvls = [int(x) for x in lvls]
    i = 0
    while i < len(lvls) - 1:
        safe, test_prev = test(lvls[i], lvls[i+1], prev) 
        if safe == 0:
            print(f"{lvls}")
            print(f"{tolerate} : {safe} : {prev} : {i} of {len(lvls)}")
            if i + 3 > len(lvls):
                print("index out of range")
                return safe
            if tolerate > 0:
                return safe
            safe, test_prev = test(lvls[i], lvls[i+2], prev) 
            tolerate += 1
        prev = test_prev

        # change = lvls[i] - lvls[i + 1]
        # if change == 0:
        #     print("unsafe: no change")
        #     return safe

        # if abs(change) > 3:
        #     print("unsafe: change greater than 3")
        #     return safe

        # if change < 0 and prev_change > 0:
        #     print("unsafe: decreasing change")
        #     return safe

        # if change > 0 and prev_change < 0:
        #     print("unsafe: increasing change")
        #     return safe

        # prev_change = change

        i += 1

    # print("safe")
    safe = 1
    return safe


def main():
    with open('input.txt') as file:
        reports = 0
        results = 0
        for report in file:
            reports += 1
            # levels = [int(x) for x in report.strip().split()]
            levels = report.strip().split()
            # safe = dumb(levels)
            safe = check(levels)
            # if safe != 0:
            #     print(levels)
            results += safe

    print(f"reports: {reports}")
    print(f"results: {results}")

if __name__ == "__main__":
    main()
