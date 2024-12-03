#!/usr/bin/env python3

# 2 sorted lists
# iterate through them to produce 

from types import SimpleNamespace


def main():
    left = []
    right = []
    dist = []
    sim_score = 0
    total = 0
    with open('input.txt') as file:
        for line in file:
            nums = line.strip().split()
            left.append(int(nums[0]))
            right.append(int(nums[1]))
    
    left.sort()
    right.sort()

    for i in range(len(left)):
        dist.append(abs(left[i] - right[i]))
        sim_score += right.count(left[i]) * left[i]
        total += dist[i]

    print(f"list: {dist}\ntotal: {total}\nsim_score: {sim_score}")

if __name__ == "__main__":
    main()
