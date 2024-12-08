import sys

from common import read_input_as_string_grid, Grid, Point


def find_starting_position(grid: Grid) -> Point:
    for x in range(grid.x_max()):
        for y in range(grid.y_max()):
            current = Point(x, y)
            if grid.get(current) == '^':
                return current
    return None


def get_next_direction(current_direction: (int, int)) -> (int, int):
    if current_direction == (-1, 0):
        return 0, 1
    elif current_direction == (0, 1):
        return 1, 0
    elif current_direction == (1, 0):
        return 0, -1
    else:
        return -1, 0


def backtrack_has_cycle(g: Grid, current: Point, bearing: (int, int), visited: set[(Point, (int, int))], ) -> bool:
    if (current, bearing) in visited:
        # print('cycle on path', visited)
        return True
    elif not current.in_bounds(g.x_max(), g.y_max()):
        # print('out bounds')
        return False
    else:
        visited.add((current, bearing))
        while g.get(current.plus(bearing)) == '#':
            bearing = get_next_direction(bearing)
        # print(("recursing on {} + {}".format(current, bearing)))
        return backtrack_has_cycle(g, current.plus(bearing), bearing, visited)


if __name__ == "__main__":
    sys.setrecursionlimit(10000)
    answer: int = 0
    grid = Grid(read_input_as_string_grid(6, False))

    starting_position: Point = find_starting_position(grid)
    path: list[Point] = []
    current_position: Point = starting_position

    offset = (-1, 0)
    while current_position.in_bounds(grid.x_max(), grid.y_max()):
        path.append(current_position)
        next_position = current_position.plus(offset)
        while grid.get(next_position) == '#':
            offset = get_next_direction(offset)
            next_position = current_position.plus(offset)
        current_position = next_position

    distinct_path = list(set(path))
    answer = len(distinct_path)
    print("Part 1: %s" % answer)

    possible_obstructions: list[Point] = []
    for i in range(1, len(distinct_path)):
        p: Point = distinct_path[i]
        if p == starting_position:
            continue
        print('Testing #{} - {}'.format(i, p))
        grid.set(p, '#')
        if backtrack_has_cycle(grid, starting_position, (-1, 0), set()):
            possible_obstructions.append(p)
        grid.set(p, '.')

    answer = len(possible_obstructions)
    # For the sample input, should be:
    # (6, 3)
    # (7, 6)
    # (7, 7)
    # (8, 1)
    # (8, 3)
    # (9, 7)
    print(possible_obstructions)
    print("Part 2: {}".format(answer))
