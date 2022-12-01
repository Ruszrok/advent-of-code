#Idea is max heap building after reading - https://www.educative.io/blog/data-structure-heaps-guide
import io
import sys
from typing import List

def parse_input(path : str) -> List[int]:
    file = io.open(path, 'r')
    sum = 0
    result = []
    for line in file.readlines():
        if(line == '\n'):
            result.append(sum)
            sum = 0
            continue

        sum += int(line)

    return result

class MaxHeap:
    def __init__(self, maxsize) -> None:
        if(maxsize < 1):
            raise Exception('Invalid max size')

        self.maxsize = maxsize
        self.size = 0 
        self.Heap = [0] * maxsize
        self.Heap[0] = sys.maxsize
        self.FRONT = 1
    
    def insert(self, el :int) -> None:
        self.FRONT += 1
        pass

    def getMax(self) -> int:
        if(self.FRONT == 1):
            raise Exception("Empty heap getMax()")
        return self.Heap[0]

path = "test.txt"

if __name__ == 'main':
    input = parse_input(path)
    mh = MaxHeap(len(input))
    map(mh.insert, input)
    print(mh.getMax())

#print(input)