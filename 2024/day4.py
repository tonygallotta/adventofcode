import re
from dataclasses import dataclass
from typing import Generic, TypeVar

from common import read_input_as_string_grid


@dataclass(frozen=True)  # make it hashable
class Point:
    x: int
    y: int

    def neighbors(self, x_max: int, y_max: int) -> set:
        possible_neighbors = [Point(self.x + x, self.y + y) for x in (-1, 0, 1) for y in (-1, 0, 1)]
        return set(filter(lambda p: p != self and p.in_bounds(x_max, y_max), possible_neighbors))

    def plus(self, offset: (int, int)):
        return Point(self.x - offset[0], self.y - offset[1])

    def difference(self, other) -> (int, int):
        return self.x - other.x, self.y - other.y

    def in_bounds(self, x_max, y_max):
        return 0 <= self.x < x_max and 0 <= self.y < y_max

    def top_left(self):
        return Point(self.x - 1, self.y - 1)

    def top_right(self):
        return Point(self.x + 1, self.y - 1)

    def bottom_left(self):
        return Point(self.x - 1, self.y + 1)

    def bottom_right(self):
        return Point(self.x + 1, self.y + 1)


@dataclass(frozen=True)
class Grid:
    data: list[list]

    def get(self, point: Point):
        if not point.in_bounds(self.x_max(), self.y_max()):
            return None
        return self.data[point.x][point.y]

    def x_max(self):
        return len(self.data)

    def y_max(self):
        return len(self.data[0])


if __name__ == "__main__":
    answer: int = 0
    grid = Grid(read_input_as_string_grid(4, False))

    x_max = grid.x_max()
    y_max = grid.y_max()
    for x in range(x_max):
        for y in range(y_max):
            current = Point(x, y)
            if grid.get(current) == 'X':
                for m_maybe in current.neighbors(x_max, y_max):
                    if grid.get(m_maybe) == 'M':
                        offset = current.difference(m_maybe)
                        a_maybe = m_maybe.plus(offset)
                        if grid.get(a_maybe) == 'A':
                            s_maybe = a_maybe.plus(offset)
                            if grid.get(s_maybe) == 'S':
                                answer += 1

    print("Part 1: %s" % answer)

    answer = 0

    # Start 1 row and 1 column in since any "A" that starts an X must fall in this bound
    for x in range(1, grid.x_max() - 1):
        for y in range(1, grid.y_max() - 1):
            current = Point(x, y)
            if grid.get(current) == 'A':
                left_diagonal = grid.get(current.top_left()) + grid.get(current.bottom_right())
                right_diagonal = grid.get(current.top_right()) + grid.get(current.bottom_left())
                if (left_diagonal == 'MS' or left_diagonal == 'SM') and (
                        right_diagonal == 'MS' or right_diagonal == 'SM'):
                    answer += 1

    print("Part 2: %s" % answer)
