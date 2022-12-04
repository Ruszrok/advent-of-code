from __future__ import annotations

import io
import re
from typing import List
from typing import Tuple

class Range:
    def __init__(self, start, finish) -> None:
        self.start = start
        self.finish = finish
    
    def contains(self, r: Range) -> bool:
        return self.start <= r.start and self.finish >= r.finish
        
def parse_input(path : str) -> List[Tuple[Range, Range]]:
    result = []
    with io.open(path, 'r') as file:
        for line in file.readlines():
            values = re.split(r',|-', line)
            result.append((Range(values[0], values[1]), Range(values[2], values[3])))

    return result

if __name__ == "__main__":
    name = "input.txt"
    ranges = parse_input(name)
    count = 0 
    for r in ranges:
        if(r[0].contains(r[1]) or r[1].contains(r[0])):
            count += 1
    
    print(count)