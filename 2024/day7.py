from dataclasses import dataclass

from common import get_file_name


@dataclass
class Equation:
    test_value: int
    numbers: list[int]

    def is_full_solution(self, operators: list[str]) -> bool:
        if len(operators) != len(self.numbers) - 1:
            return False
        result = self.numbers[0]
        for i in range(1, len(self.numbers)):
            operator = operators[i - 1]
            if operator == '+':
                result += self.numbers[i]
            elif operator == '*':
                result *= self.numbers[i]
            else:
                result = int('{}{}'.format(result, self.numbers[i]))
            if result > self.test_value:
                return False
        # print('Checking {} on {} -> {}'.format(operators, self, result))
        return result == self.test_value

    def is_viable_partial_solution(self, operators: list[str]) -> bool:
        if len(operators) >= len(self.numbers):
            return False
        elif len(operators) == 0:
            return True
        result = self.numbers[0]
        for i in range(1, len(operators)):
            operator = operators[i - 1]
            if operator == '+':
                result += self.numbers[i]
            elif operator == '*':
                result *= self.numbers[i]
            else:
                result = int('{}{}'.format(result, self.numbers[i]))
            if result > self.test_value:
                return False
        return len(operators) <= len(self.numbers) - 1


def backtrack_dfs_solve(e: Equation, partial_solution: list[str], possible_operators: list[str]) -> bool:
    if e.is_full_solution(partial_solution):
        return True
    elif not e.is_viable_partial_solution(partial_solution):
        return False
    else:
        for operator in possible_operators:
            if backtrack_dfs_solve(e, partial_solution + [operator], possible_operators):
                return True
    return False


if __name__ == "__main__":
    answer: int = 0
    equations: list[Equation] = []
    with open(get_file_name(7, False)) as file:
        for line in file:
            test_value = line.strip().split(':')[0]
            numbers = line.strip().split()[1:]
            equations.append(Equation(int(test_value), [int(n) for n in numbers]))

    for equation in equations:
        if backtrack_dfs_solve(equation, [], ['+', '*']):
            answer += equation.test_value

    print("Part 1: {}".format(answer))

    answer = 0

    for idx, equation in enumerate(equations):
        print('checking equation {}'.format(idx))
        if backtrack_dfs_solve(equation, [], ['+', '*', '||']):
            answer += equation.test_value

    print("Part 2: {}".format(answer))
