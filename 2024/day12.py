from collections import defaultdict, OrderedDict
from itertools import combinations

from common import read_input_as_string_grid, Grid, Point


def explore_region(grid: Grid, point: Point, value: str, region: set[Point], visited: set[Point]) -> frozenset[Point]:
    visited.add(point)
    for n in point.neighbors_cross(grid.x_max(), grid.y_max()):
        if grid.get(n) == value and not n in visited:
            region.add(n)
            explore_region(grid, n, value, region, visited)

    return frozenset(region)

if __name__ == "__main__":
    answer: int = 0
    grid = Grid(read_input_as_string_grid(12, True, 1))
    print(grid)

    regions: set[frozenset[Point]] = set()

    for i in range(grid.x_max()):
        for j in range(grid.y_max()):
            point = Point(i, j)
            regions.add(explore_region(grid, point, grid.get(point), {point}, {point}))

    for idx, region in enumerate(regions):
        perimeter: OrderedDict[Point, int] = OrderedDict()
        print("Checking region {} of {}".format(idx, len(region)))
        for point in region:
            # print('checking neighbors of ', point, point.neighbors_unbounded_cross())
            for n in point.neighbors_unbounded_cross():
                if not n in region:
                    border_count = sum(1 if n2 in region else 0 for n2 in n.neighbors_cross(grid.x_max(), grid.y_max()))
                    perimeter[n] = max(perimeter.get(n, 0), border_count)
        # print('Region {} ({}) has size {} and perimeter {}'.format(idx, grid.get(list(region)[0]), len(region),
        #                                                            sum(perimeter.values())))
        # grid.print(lambda p: grid.get(p) if p in region else ('|' if p in perimeter else '.'))
        answer += len(region) * sum(perimeter.values())

    print("Part 1: {}".format(answer))

    answer = 0

    print("Part 2: {}".format(answer))
