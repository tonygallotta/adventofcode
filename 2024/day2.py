from common import read_input_as_int_arrays


def check_is_safe(report: list[int], check_increasing: bool):
    for i in range(1, len(report)):
        prev_level = report[i - 1]
        current_level = report[i]
        diff = abs(current_level - prev_level)
        if check_increasing and current_level < prev_level:
            return False
        elif (not check_increasing) and current_level > prev_level:
            return False
        elif diff == 0 or diff > 3:
            return False
    return True


if __name__ == "__main__":
    reports = read_input_as_int_arrays(2)
    answer: int = 0
    for r in range(len(reports)):
        report = reports[r]
        if check_is_safe(report, report[0] < report[1]):
            answer += 1

    print("Part 1: %s" % answer)

    answer = 0
    for r in range(len(reports)):
        report = reports[r]
        is_safe = check_is_safe(report, True) or check_is_safe(report, False)
        if is_safe:
            # print("report %s is safe AS IS" % r)
            answer += 1
            continue
        for i in range(len(report)):
            report_with_removal = report.copy()
            del report_with_removal[i]
            is_safe = check_is_safe(report_with_removal, True) or check_is_safe(report_with_removal, False)
            if is_safe:
                # print("report %s is safe by deleting %s" % (r, i))
                answer += 1
                break
        else:
            continue

    print("Part 2: %s" % answer)
