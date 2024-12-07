import itertools
from dataclasses import dataclass

from common import get_file_name


@dataclass
class Equation:
    test_value: int
    numbers: list[int]


if __name__ == "__main__":
    answer: int = 0
    equations: list[Equation] = []
    with open(get_file_name(7, False)) as file:
        for line in file:
            test_value = line.strip().split(':')[0]
            numbers = line.strip().split()[1:]
            equations.append(Equation(int(test_value), [int(n) for n in numbers]))

    valid_test_values = set()
    counter = 0
    for equation in equations:
        # No way this will scale for pt 2
        num_operators = len(equation.numbers) - 1
        possible_operator_combos = list(itertools.product(*[['+', '*'] for i in range(num_operators)]))
        solved = False
        for ops in possible_operator_combos:
            result = equation.numbers[0]
            for i in range(1, len(equation.numbers)):
                counter += 1
                if ops[i - 1] == '+':
                    result += equation.numbers[i]
                else:
                    result *= equation.numbers[i]
            if result == equation.test_value:
                # print('{} is solved with {}'.format(equation, ops))
                valid_test_values.add(result)
                solved = True
                break
            if solved:
                break

    print("Part 1: {}".format(sum(valid_test_values)))
