
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

GAME_SCORE_1 = {(1,1):4, (1,2):8,(1,3):3,(2,1):1,(2,2):5,(2,3):9,(3,1):7,(3,2):2,(3,3):6}
GAME_SCORE_2 = {(1,1):3, (1,2):4,(1,3):8,(2,1):1,(2,2):5,(2,3):9,(3,1):2,(3,2):6,(3,3):7}

if __name__ == "__main__":
    input = parse_input("input.txt")
    total = 0
    total2 = 0
    for pair in input:
        total += GAME_SCORE_1[pair]
        total2 += GAME_SCORE_2[pair]

    print(total)
    print(total2)
