import re

from common import read_input_as_strings

if __name__ == "__main__":
    answer: int = 0
    program_lines = read_input_as_strings(3, False)
    for program_line in program_lines:
        matches: list = re.findall("mul\(\d+,\d+\)", program_line)
        for match in matches:
            first, second = re.findall("(\d+)", match)
            answer += int(first) * int(second)

    print("Part 1: %s" % answer)

    answer = 0

    # NOTE: The input is (critically) all 1 line. Left this loop anyway
    for i in range(len(program_lines)):
        program_line = program_lines[i]
        # Remove all don't()...do() blocks, lazy matching
        program_line = re.sub("don't\(\).*?do\(\)", "", program_line)
        # Remove everything after the last unclosed don't()
        program_line = re.sub("don't\(\).*$", "", program_line)
        matches: list = re.findall("mul\(\d+,\d+\)", program_line)
        for match in matches:
            first, second = re.findall("(\d+)", match)
            answer += int(first) * int(second)
    print("Part 2: %s" % answer)
