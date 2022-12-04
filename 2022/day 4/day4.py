from __future__ import annotations

import io
import re
from typing import List
from typing import Tuple

class Range:
    def __init__(self, start : int, finish : int) -> None:
        self.start = start
        self.finish = finish
    
    def contains(self, r: Range) -> bool:
        return self.start <= r.start and self.finish >= r.finish

    def __str__(self) -> str:
        return f'{self.start}-{self.finish}'
        
def parse_input(path : str) -> List[Tuple[Range, Range]]:
    result = []
    with io.open(path, 'r') as file:
        for line in file.readlines():
            values = re.split(r',|-', line)
            result.append((Range(int(values[0].strip()), int(values[1].strip())), Range(int(values[2].strip()), int(values[3].strip()))))

    return result

if __name__ == "__main__":
    name = "input.txt"
    ranges = parse_input(name)
    count = 0 
    for r in ranges:
        if(r[0].contains(r[1]) or r[1].contains(r[0])):
            print(r[0], r[1])
            count += 1
    
    print(count)