from day1 import MaxHeap
import pytest

def test_raise_getMaxHeap() -> None:
    with pytest.raises(Exception):
        mh = MaxHeap(0)


def test_raise_min_getMaxHeap() -> None:
    with pytest.raises(Exception):
        mh = MaxHeap(1)
        assert mh.getMax()

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
    mh = MaxHeap(3)
    mh.insert(1)
    mh.insert(2)
    mh.insert(3)
    mh.insert(3)
    assert 3 == mh.getMax(), "Four elements works wrong"

def test_four_inorder_getMaxHeap() -> None:
    mh = MaxHeap(3)
    mh.insert(1000)
    mh.insert(2)
    mh.insert(500)
    mh.insert(3)
    assert 1000 == mh.getMax(), "Four elements works wrong"
