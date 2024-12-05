from math import floor

from common import get_file_name

if __name__ == "__main__":
    answer: int = 0
    rules: list[(int, int)] = []
    updates: list[list[int]] = []
    with open(get_file_name(5, False)) as file:
        past_rules = False
        for line in file:
            if line.strip() == "":
                past_rules = True
                continue
            if past_rules:
                updates.append([int(x) for x in line.strip().split(',')])
            else:
                a, b = line.strip().split('|')
                rules.append((int(a), int(b)))

    correctly_ordered_updates: list[list[int]] = []
    incorrectly_ordered_updates: list[list[int]] = []
    for update in updates:
        correct = True
        for rule in rules:
            try:
                if update.index(rule[0]) > update.index(rule[1]):
                    correct = False
                    break
            except ValueError:
                pass
        if correct:
            correctly_ordered_updates.append(update)
        else:
            incorrectly_ordered_updates.append(update)

    for update in correctly_ordered_updates:
        middle_page = update[floor((len(update) - 1) / 2)]
        answer += middle_page
    print("Part 1: %s" % answer)

    answer = 0

    for update in incorrectly_ordered_updates:
        while True:
            # Bubble sort
            changed = False
            for rule in rules:
                try:
                    first_item_idx = update.index(rule[0])
                    second_item_idx = update.index(rule[1])
                    if first_item_idx > second_item_idx:
                        # swap
                        temp = update[first_item_idx]
                        update[first_item_idx] = update[second_item_idx]
                        update[second_item_idx] = temp
                        changed = True
                except ValueError:
                    pass
            if not changed:
                break

    for update in incorrectly_ordered_updates:
        middle_page = update[floor((len(update) - 1) / 2)]
        answer += middle_page

    print("Part 2: %s" % answer)
