from day1 import MaxHeap
import pytest

def test_raise_invalid_maxsize_MaxHeap() -> None:
    with pytest.raises(Exception):
        mh = MaxHeap(0)


def test_raise_min_getMaxHeap() -> None:
    with pytest.raises(Exception):
        mh = MaxHeap(1)
        mh.getMax()

def test_raise_overflow_MaxHeap() -> None:
    with pytest.raises(Exception):
        mh = MaxHeap(1)
        mh.insert(0)
        mh.insert(1)

def test_base_getMaxHeap() -> None:
    mh = MaxHeap(3)
    mh.insert(1)
    assert 1==mh.getMax(), "One element works wrong"

def test_three_getMaxHeap() -> None:
    mh = MaxHeap(3)
    mh.insert(1)
    mh.insert(2)
    mh.insert(3)
    assert 3 == mh.getMax(), "Three elements works wrong"

def test_four_getMaxHeap() -> None:
    mh = MaxHeap(4)
    mh.insert(1)
    mh.insert(2)
    mh.insert(3)
    mh.insert(3)
    assert 3 == mh.getMax(), "Four elements works wrong"

def test_four_inorder_getMaxHeap() -> None:
    mh = MaxHeap(4)
    mh.insert(1000)
    mh.insert(2)
    mh.insert(500)
    mh.insert(3)
    assert 1000 == mh.getMax(), "Four elements works wrong"

def test_four_inorder_getMaxHeap() -> None:
    mh = MaxHeap(4)
    mh.insert(1000)
    mh.insert(2)
    mh.insert(500)
    mh.insert(5000)
    assert 5000 == mh.getMax(), "Four elements works wrong"

def test_four_inorder_extractMaxHeap() -> None:
    mh = MaxHeap(4)
    mh.insert(1000)
    mh.insert(2)
    mh.insert(500)
    mh.insert(5000)
    assert 5000 == mh.extractMax(), "Four elements works wrong"
    assert 1000 == mh.extractMax(), "Four elements works wrong"

def test_five_inorder_bug_fix_extractMaxHeap() -> None:
    mh = MaxHeap(5)
    mh.insert(1000)
    mh.insert(2)
    mh.insert(500)
    mh.insert(5000)
    mh.insert(300)
    assert 5000 == mh.extractMax(), "Four elements works wrong"
    assert 1000 == mh.extractMax(), "Four elements works wrong"
    assert 500 == mh.extractMax(), "Four elements works wrong"