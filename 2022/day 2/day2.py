
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
    total2 = 0
    for pair in input:
        #first total calculation
        total += pair[1]
        if(pair[0] == pair[1]):
            total += 3
        elif(pair in [(1,2), (2,3), (3,1)]):
            total += 6
        
        #second strategy calculation
        #1 - lose, 2 - draw, 3 - win
        total2 += 3 * (pair[1] - 1)
        if(pair[1] == 2):
            total2 += pair[0]
        elif(pair[1] == 3):
            if(pair[0] == 1):
                total2 += 2
            elif(pair[0] == 2):
                total2 += 3
            else:
                total2 += 1
        else:
            if(pair[0] == 1):
                total2 += 3
            elif(pair[0] == 2):
                total2 += 1
            else:
                total2 += 2
#expected - 14163
#12091
    print(total)
    print(total2)
