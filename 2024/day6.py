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


def has_cycle(g: Grid, max_length: int) -> bool:
    o = (-1, 0)
    c: Point = find_starting_position(g)
    # print('starting from ', c)
    visited = []
    while c.in_bounds(g.x_max(), g.y_max()):
        visited.append(c)
        n = c.plus(o)
        while g.get(n) == '#':
            o = get_next_direction(o)
            n = c.plus(o)

        for x in range(1, len(visited) - 1):
            if visited[x] == n and visited[x - 1] == c:
                print('cycle on path', visited)
                return True
        c = n
        if len(visited) > max_length + 2:
            return False

    return False


if __name__ == "__main__":
    answer: int = 0
    grid = Grid(read_input_as_string_grid(6, False))

    starting_position: Point = find_starting_position(grid)
    path: list[Point] = []
    current_position: Point = starting_position

    print("Starting at {}".format(starting_position))
    offset = (-1, 0)
    while current_position.in_bounds(grid.x_max(), grid.y_max()):
        path.append(current_position)
        next_position = current_position.plus(offset)
        while grid.get(next_position) == '#':
            offset = get_next_direction(offset)
            next_position = current_position.plus(offset)
        current_position = next_position

    answer = len(set(path))
    print("Part 1: %s" % answer)

    possible_obstructions: list[Point] = []
    for i in range(1, len(path)):
        p: Point = path[i]
        if p == starting_position:
            continue
        print('Testing #{}'.format(i))
        grid.set(p, '#')
        # print('checking grid:\n', grid)
        if has_cycle(grid, i):
            possible_obstructions.append(p)
        grid.set(p, '.')

    answer = len(set(possible_obstructions))
    # For the sample input, should be:
    # (6, 3)
    # (7, 6)
    # (7, 7)
    # (8, 1)
    # (8, 3)
    # (9, 7)
    print(possible_obstructions)
    # 515 is too low
    print("Part 2: %s" % answer)
