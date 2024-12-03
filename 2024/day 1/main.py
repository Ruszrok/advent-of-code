import os
import argparse
import sys

def read_input(file_path: str) -> str:
    l, r = [], []
    with open(file_path, 'r') as f:
        lines = f.readlines()
        for line in lines:
            parts = line.split('   ')
            a1 = int(parts[0].strip())
            b1 = int(parts[1].strip())
            l.append(a1)
            r.append(b1)    

    return l, r 

def distance(a, b) -> int:
    if (len(a) != len(b)):
        print('Error: lists must have the same length')
        sys.exit(1)

    distance = 0
    for i in range(len(a)):
        distance += abs(a[i] - b[i])
    return distance

def similarity(a, b) -> int:
    if (len(a) != len(b)):
        print('Error: lists must have the same length')
        sys.exit(1)

    similarity = 0
    b_map = {}
    for s in b:
        if s in b_map:
            b_map[s] += 1
        else:
            b_map[s] = 1

    for n in a:
        if n in b_map:
            similarity += n * b_map[n]
    return similarity

if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Advent of Code 2024 - Day 1')
    parser.add_argument('-i', type=str, help='Input file')
    args = parser.parse_args()

    if not os.path.exists(args.i):
        print(f'Error: file not found: {args.i}')
        sys.exit(1)

    a1,b1 = read_input(args.i)
    a_s = sorted(a1)
    b_s = sorted(b1)
    print(distance(a_s, b_s))
    print(similarity(a_s, b_s))
    
