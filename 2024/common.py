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


def read_input_as_string_grid(day, sample):
    results = []
    with open(get_file_name(day, sample)) as file:
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


def get_file_name(day, sample):
    return "input/day{}{}.txt".format(day, "_sample" if sample else "")


@dataclass(frozen=True)  # make it hashable
class Point:
    x: int
    y: int

    def neighbors(self, x_max: int, y_max: int) -> set:
        possible_neighbors = [Point(self.x + x, self.y + y) for x in (-1, 0, 1) for y in (-1, 0, 1)]
        return set(filter(lambda p: p != self and p.in_bounds(x_max, y_max), possible_neighbors))

    def plus(self, offset: (int, int)):
        return Point(self.x + offset[0], self.y + offset[1])

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
