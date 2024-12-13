from dataclasses import dataclass


def read_input_as_int_arrays(day, sample=False):
    results = []
    with open(get_file_name(day, sample)) as file:
        for line in file:
            values = line.strip().split()
            results.append([int(v) for v in values])
    return results


def read_input_as_string_arrays(day, sample):
    results = []
    with open(get_file_name(day, sample)) as file:
        for line in file:
            values = line.strip().split()
            results.append(values)
    return results


def read_input_as_string_grid(day, sample, sample_num=1):
    results = []
    with open(get_file_name(day, sample, sample_num)) as file:
        for line in file:
            values = line.strip()
            line_chars = []
            for c in values:
                line_chars.append(c)
            results.append(line_chars)
    return results


def read_input_as_strings(day, sample):
    results = []
    with open(get_file_name(day, sample)) as file:
        for line in file:
            results.append(line.strip())
    return results


def get_file_name(day, sample, sample_num=1):
    return "input/day{}{}{}.txt".format(day, "_sample" if sample else "", sample_num if sample_num > 1 else "")


@dataclass(frozen=True)  # make it hashable
class Point:
    x: int
    y: int

    def neighbors(self, x_max: int, y_max: int) -> set:
        return set(filter(lambda p: p != self and p.in_bounds(x_max, y_max), self.neighbors_unbounded()))

    def neighbors_unbounded(self) -> set:
        return set([Point(self.x + x, self.y + y) for x in (-1, 0, 1) for y in (-1, 0, 1)])

    def neighbors_unbounded_cross(self) -> set:
        return set([Point(self.x + x, self.y + y) for x, y in [(-1, 0), (0, -1), (0, 1), (1, 0)]])

    def neighbors_cross(self, x_max: int, y_max: int) -> set:
        return set(filter(lambda p: p != self and p.in_bounds(x_max, y_max), self.neighbors_unbounded_cross()))

    def plus(self, offset: (int, int)):
        return Point(self.x + offset[0], self.y + offset[1])

    def minus(self, offset: (int, int)):
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

    def set(self, point: Point, value):
        self.data[point.x][point.y] = value

    def x_max(self):
        return len(self.data)

    def y_max(self):
        return len(self.data[0])

    def print(self, print_fn) -> None:
        for i in range(self.x_max()):
            for j in range(self.y_max()):
                point = Point(i, j)
                if print_fn is None:
                    print(self.get(point), end='')
                else:
                    print(print_fn(point), end='')
            print('')
