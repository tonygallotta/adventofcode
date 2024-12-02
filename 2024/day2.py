from collections import Counter

reports = []

with open("input/day2.txt") as file:
    for line in file:
        values = line.strip().split()
        reports.append([int(v) for v in values])


def check_is_safe(report, check_increasing):
    is_safe = True
    for i in range(1, len(report)):
        prev_level = report[i - 1]
        current_level = report[i]
        diff = abs(current_level - prev_level)
        if check_increasing and current_level < prev_level:
            is_safe = False
        elif (not check_increasing) and current_level > prev_level:
            is_safe = False
        elif diff > 3 or diff == 0:
            # print("Diff is %s between %s and %s, not safe" % (diff, current_level, prev_level))
            is_safe = False
    return is_safe

print(reports)
answer = 0
for r in range(len(reports)):
    report = reports[r]
    check_increasing = report[0] < report[1]
    if check_is_safe(report, check_increasing):
        answer += 1
        # print("SAFE")
    # else:
    #     # print("UNSAFE")

print("Part 1: %s" % answer)

answer = 0
for r in range(len(reports)):
    report = reports[r]
    is_safe = check_is_safe(report, True) or check_is_safe(report, False)
    if is_safe:
        print("report %s is safe AS IS" % r)
        answer += 1
        continue
    for i in range(len(report)):
        report_with_removal = report.copy()
        del report_with_removal[i]
        is_safe = check_is_safe(report_with_removal, True) or check_is_safe(report_with_removal, False)
        if is_safe:
            print("report %s is safe by deleting %s" % (r, i))
            answer += 1
            break
    else:
        continue

print("Part 2: %s" % answer)