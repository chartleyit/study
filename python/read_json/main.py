#!/usr/bin/env python3

import json

def read_json(file_path):
    with open(file_path) as file:
        data = json.load(file)

    return data

if __name__ == "__main__":

    file = "data.json"
    file_out = read_json(file)

    if "domain" not in file_out:
        print(f"domain missing from file")
    else:
        print(f"domain in file")

    for k in file_out:
        print(f"{k}")
