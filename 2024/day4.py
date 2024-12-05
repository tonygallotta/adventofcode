from common import read_input_as_string_grid, Grid, Point

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
