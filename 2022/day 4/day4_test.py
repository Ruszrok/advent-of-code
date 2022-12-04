from day4 import Range
import pytest

def test_Contains() -> None:
    r1 = Range(1, 1)
    r2 = Range(0, 0)
    assert r1.contains(r2) == False
    assert r2.contains(r1) == False

def test_Contains_2() -> None:
    r1 = Range(1, 1)
    r2 = Range(1, 1)
    assert r1.contains(r2) == True
    assert r2.contains(r1) == True

def test_Contains_3() -> None:
    r1 = Range(1, 4)
    r2 = Range(2, 3)
    assert r1.contains(r2) == True
    assert r2.contains(r1) == False

def test_Contains_4() -> None:
    r1 = Range(1, 4)
    r2 = Range(1, 1)
    assert r1.contains(r2) == True
    assert r2.contains(r1) == False

def test_Contains_5() -> None:
    r1 = Range(1, 4)
    r2 = Range(1, 2)
    assert r1.contains(r2) == True
    assert r2.contains(r1) == False

def test_Contains_6() -> None:
    r1 = Range(1, 4)
    r2 = Range(2, 4)
    assert r1.contains(r2) == True
    assert r2.contains(r1) == False

def test_Contains_7() -> None:
    r1 = Range(1, 4)
    r2 = Range(2, 5)
    assert r1.contains(r2) == False
    assert r2.contains(r1) == False