def read_input_as_int_arrays(day, sample = False):
    results = []
    with open(get_file_name(day, sample)) as file:
        for line in file:
            values = line.strip().split()
            results.append([int(v) for v in values])
    return results

def read_input_as_string_arrays(day, sample):
    results = []
    with open(get_file_name(day, sample)) as file:
        for line in file:
            values = line.strip().split()
            results.append(values)
    return results

def read_input_as_string_grid(day, sample):
    results = []
    with open(get_file_name(day, sample)) as file:
        for line in file:
            values = line.strip()
            line_chars = []
            for c in values:
                line_chars.append(c)
            results.append(line_chars)
    return results

def read_input_as_strings(day, sample):
    results = []
    with open(get_file_name(day, sample)) as file:
        for line in file:
            results.append(line.strip())
    return results

def get_file_name(day, sample):
    return "input/day{}{}.txt".format(day, "_sample" if sample else "")