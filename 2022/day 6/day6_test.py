import day6
import pytest

testdata = [
    ("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4, 7),
    ("bvwbjplbgvbhsrlpgdmjqwftvncz", 4, 5),
    ("nppdvjthqldpwncqszvftbrmjlhg", 4, 6),
    ("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4, 10),
    ("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4, 11),
    ("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14, 19),
    ("bvwbjplbgvbhsrlpgdmjqwftvncz", 14, 23),
    ("nppdvjthqldpwncqszvftbrmjlhg", 14, 23),
    ("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14, 29),
    ("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 14, 26),
]


@pytest.mark.parametrize("a,size, expected", testdata)
def test_findposition_v0_test(a, size, expected):
    res = day6.find_start_position(a, size)
    assert res == expected