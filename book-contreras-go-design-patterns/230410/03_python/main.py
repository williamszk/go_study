"""experiment with closures in python
"""

from typing import Callable

def main():

    def add_n(n: int) -> Callable[[int], int]:
        def inner(m: int) -> int:
            return m + n
        return inner
    
    add_five = add_n(5)
    out1 = add_five(10)
    print(f"{out1 = }")

if __name__ == "__main__":
    main()