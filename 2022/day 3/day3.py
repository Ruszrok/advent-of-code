
import io
from typing import List


def parse_input(path : str) -> List[str]:
    result = []
    with io.open(path, 'r') as file:
        for line in file.readlines():
            result.append(line.strip())

    return result

def to_priority(code: str) -> int:
    c = ord(code)
    if(c > 90):
        return c - 96
    else:
        return c - 64 + 26

if __name__ == "__main__":
    input = parse_input("input.txt")
    result = 0
    for r in input:
        items = {}
        for i in r[:len(r)//2]:
            items[i] = 1
        
        for i in r[len(r)//2:]:
            if(i in items):
                result += to_priority(i)
                break
    #print(result, "Expected: 157")
    print(result)
        