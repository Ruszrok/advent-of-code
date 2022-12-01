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
    
    result.append(sum)

    return result

class MaxHeap:
    #Packed tree [sys.max, r, r_l, r_r,r_l_l, r_l_r, 0, 0]
    def __init__(self, maxsize) -> None:
        if(maxsize < 1):
            raise Exception('Invalid max size')

        self.maxsize = maxsize
        self.size = 0 
        self.Heap = [0] * (maxsize + 1)
        self.Heap[0] = sys.maxsize
    
    def __parent(self, index) -> int:
        return index // 2

    def __left(self, index) -> int:
        return index * 2

    def __right(self, index) -> int:
        return index * 2 + 1

    def insert(self, el :int) -> None:
        if(self.size >= self.maxsize):
            raise Exception('overflow')

        self.size += 1
        self.Heap[self.size] = el
        current = self.size
        
        while self.Heap[self.__parent(current)] < self.Heap[current]:
            self.Heap[self.__parent(current)], self.Heap[current] = self.Heap[current], self.Heap[self.__parent(current)]
            current = self.__parent(current)
        

    def getMax(self) -> int:
        if(self.size == 0):
            raise Exception("Empty heap getMax()")
        return self.Heap[1]
    
    def __heapify(self, index) -> None:
        left = self.__left(index)
        right = self.__right(index)

        if left > self.size and right > self.size:
            return
        
        if(self.Heap[left] > self.Heap[index]):
            self.Heap[left], self.Heap[index] = self.Heap[index], self.Heap[left]
            self.__heapify(left)        
        if (self.Heap[right] > self.Heap[index]):
            self.Heap[right], self.Heap[index] = self.Heap[index], self.Heap[right]
            self.__heapify(right)

    #I use it instead of extraction and heap rebuild
    def extractMax(self) -> int:
        max_value = self.Heap[1]
        self.Heap[1] = self.Heap[self.size]
        self.size -= 1
        self.__heapify(1)

        return max_value

path = "input.txt"

if __name__ == '__main__':
    input = parse_input(path)
    mh = MaxHeap(len(input))
    for i in input:
        mh.insert(i)
    print(mh.getMax())
    print(mh.extractMax() + mh.extractMax() + mh.extractMax())
#print(input)