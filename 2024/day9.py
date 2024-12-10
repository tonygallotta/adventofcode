from math import floor

from common import get_file_name, read_input_as_strings

if __name__ == "__main__":
    answer: int = 0
    condensed_disk_map: list[str] = list(read_input_as_strings(9, False)[0])
    full_disk_map: list[(int, int)] = []

    file_id = 0
    full_disk_map_idx = 0
    for idx, block_count in enumerate(condensed_disk_map):
        if idx % 2 == 0:
            full_disk_map.append((file_id, int(block_count)))
            file_id += 1
        else:
            full_disk_map.append((None, int(block_count)))
    original_full_disk_map = full_disk_map.copy()

    left_idx: int = 0
    right_idx: int = len(full_disk_map) - 1

    while full_disk_map[right_idx][0] is None:
        right_idx -= 1
    while full_disk_map[left_idx][0] != None:
        left_idx += 1

    while left_idx < right_idx:
        left_slot = full_disk_map[left_idx]
        right_file = full_disk_map[right_idx]
        left_free_remaining = left_slot[1]
        right_file_remaining = right_file[1]

        if right_file_remaining == left_free_remaining:
            full_disk_map[left_idx] = right_file
            full_disk_map[right_idx] = (None, right_file_remaining)
        elif right_file_remaining < left_free_remaining:
            full_disk_map[left_idx] = right_file
            full_disk_map.insert(left_idx + 1, (None, left_free_remaining - right_file_remaining))
            full_disk_map[right_idx + 1] = (None, right_file[1])
            left_idx += 1
            right_idx -= 1
        else:
            # right_file_remaining > left_free_remaining
            filled = left_free_remaining
            # fill the whole slot
            full_disk_map[left_idx] = (right_file[0], filled)
            full_disk_map[right_idx] = (right_file[0], right_file_remaining - filled)
            left_idx += 1

        while full_disk_map[right_idx][0] is None:
            right_idx -= 1
        while full_disk_map[left_idx][0] != None:
            left_idx += 1

    position = 0
    for idx, f in enumerate(full_disk_map):
        if f[0] is None:
            break
        for _ in range(f[1]):
            answer += position * f[0]
            position += 1

    print("Part 1: {}".format(answer))

    answer = 0

    full_disk_map = original_full_disk_map
    left_idx: int = 0
    right_idx: int = len(full_disk_map) - 1

    while full_disk_map[right_idx][0] is None:
        right_idx -= 1
    while full_disk_map[left_idx][0] != None:
        left_idx += 1
    while right_idx > 0:
        left_slot = full_disk_map[left_idx]
        right_file = full_disk_map[right_idx]
        left_free_remaining = left_slot[1]
        right_file_remaining = right_file[1]

        if left_idx >= right_idx:
            left_idx = 0
            right_idx -= 1
        elif right_file_remaining == left_free_remaining:
            full_disk_map[left_idx] = right_file
            full_disk_map[right_idx] = (None, right_file_remaining)
            left_idx = 0
            right_idx -= 1
        elif right_file_remaining < left_free_remaining:
            full_disk_map[left_idx] = right_file
            full_disk_map.insert(left_idx + 1, (None, left_free_remaining - right_file_remaining))
            full_disk_map[right_idx + 1] = (None, right_file[1])
            left_idx = 0
            right_idx -= 1
        else:
            left_idx += 1

        while full_disk_map[right_idx][0] is None:
            right_idx -= 1
        while full_disk_map[left_idx][0] != None:
            left_idx += 1

    position = 0
    for idx, f in enumerate(full_disk_map):
        if f[0] is None:
            position += f[1]
            continue
        for _ in range(f[1]):
            answer += position * f[0]
            position += 1

    # 6287317016845
    print("Part 2: {}".format(answer))
