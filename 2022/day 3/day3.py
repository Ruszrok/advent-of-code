
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

def first_task_solution(input: List[str]) -> int:
    result = 0
    for r in input:
        items = {}
        for i in r[:len(r)//2]:
            items[i] = 1
        
        for i in r[len(r)//2:]:
            if(i in items):
                result += to_priority(i)
                break
    return result

def second_task_solution(input: List[str]) -> int:
    result = 0
    for r1, r2, r3 in zip(*[iter(input)] * 3):
        items = {}
        for i in r1:
            items[i] = 1
        
        for i in r2:
            if(i in items):
                items[i] = 2
        
        for i in r3:
            if(i in items and items[i] == 2):
                result += to_priority(i)
                break


    return result


if __name__ == "__main__":
    input = parse_input("input.txt")
    result1 = first_task_solution(input)
    result2 = second_task_solution(input)
    print(result1)
    print(result2)
        