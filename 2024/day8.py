from collections import defaultdict
from itertools import combinations

from common import read_input_as_string_grid, Grid, Point

if __name__ == "__main__":
    answer: int = 0
    grid = Grid(read_input_as_string_grid(8, False))

    points_by_frequency: dict[str, list[Point]] = defaultdict(list)

    for i in range(grid.x_max()):
        for j in range(grid.y_max()):
            point = Point(i, j)
            value = grid.get(point)
            if value != '.':
                points_by_frequency[value].append(point)

    antinodes: set[Point] = set()
    for antenna, points in points_by_frequency.items():
        for p1, p2 in list(combinations(points, 2)):
            slope = p1.difference(p2)
            antinode1 = p1.plus(slope)
            antinode2 = p2.minus(slope)
            if antinode1.in_bounds(grid.x_max(), grid.y_max()):
                antinodes.add(antinode1)

            if antinode2.in_bounds(grid.x_max(), grid.y_max()):
                antinodes.add(antinode2)

    grid.print(lambda p: '#' if p in antinodes else grid.get(p))

    answer = len(antinodes)
    print("Part 1: {}".format(answer))

    answer = 0

    for antenna, points in points_by_frequency.items():
        for p1, p2 in list(combinations(points, 2)):
            slope = p1.difference(p2)
            p1_antinode = p1
            while p1_antinode.in_bounds(grid.x_max(), grid.y_max()):
                antinodes.add(p1_antinode)
                p1_antinode = p1_antinode.plus(slope)
            p2_antinode = p2
            while p2_antinode.in_bounds(grid.x_max(), grid.y_max()):
                antinodes.add(p2_antinode)
                p2_antinode = p2_antinode.minus(slope)
            iter += 1

    answer = len(antinodes)
    print("Part 2: {}".format(answer))
