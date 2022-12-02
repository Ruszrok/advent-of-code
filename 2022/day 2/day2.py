
import io
from typing import List
from typing import Tuple
                 #Rock, Paper, Scisors
DICT_FIRST = {'A' : 1, 'B' : 2, 'C' : 3}
DICT_SECOND = {'X' : 1, 'Y' : 2, 'Z' : 3}

def parse_input(path : str) -> List[Tuple[int, int]]:
    result = []
    with io.open(path, 'r') as file:
        for line in file.readlines():
            parts = [s.strip() for s in line.split(' ')]
            result.append((DICT_FIRST[parts[0]], DICT_SECOND[parts[1]]))

    return result

if __name__ == "__main__":
    input = parse_input("input.txt")
    #input = [(1, 2), (2,1), (3,3)]
    total = 0
    for pair in input:
        total += pair[1]
        if(pair[0] == pair[1]):
            total += 3
        elif(pair in [(1,2), (2,3), (3,1)]):
            total += 6
    print(total)
