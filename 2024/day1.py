from collections import Counter

left_list = []
right_list = []

with open("input/day1.txt") as file:
    for line in file:
        values = line.strip().split()
        left_list.append(int(values[0]))
        right_list.append(int(values[1]))

left_list.sort()
right_list.sort()
answer = 0
for i in range(len(left_list)):
    answer += abs(left_list[i] - right_list[i])

print("Part 1: %s" % answer)

answer = 0
right_list_occurrences = Counter(right_list)
for v in left_list:
    answer += v * right_list_occurrences.get(v, 0)

print("Part 2: %s" % answer)