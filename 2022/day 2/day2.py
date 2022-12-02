
import io
from typing import List
from typing import Tuple

def parse_input(path : str) -> List[Tuple[int, int]]:
    result = []
    with io.open(path, 'r') as file:
        for line in file.readlines():
            parts = [s.strip() for s in line.split(' ')]
            result.append((parts[0],parts[1]))

    return result

GAME_SCORE_1 = {('A','X'):4, ('A','Y'):8,('A','Z'):3,
                    ('B','X'):1,('B','Y'):5,('B','Z'):9,
                    ('C','X'):7,('C','Y'):2,('C','Z'):6}
GAME_SCORE_2 = {('A','X'):3, ('A','Y'):4,('A','Z'):8,
                    ('B','X'):1,('B','Y'):5,('B','Z'):9,
                    ('C','X'):2,('C','Y'):6,('C','Z'):7}

if __name__ == "__main__":
    input = parse_input("input.txt")
    total = 0
    total2 = 0
    for pair in input:
        total += GAME_SCORE_1[pair]
        total2 += GAME_SCORE_2[pair]

    print(total)
    print(total2)
