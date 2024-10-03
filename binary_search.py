"""
define a function which will take a list and a target value
binary search splits the list into two halves and compares the target value to the middle value
we start at 0 index
"""


def binary_search(list: list, target: int):
    first_val: int = 0
    last_val: int = len(list) - 1
