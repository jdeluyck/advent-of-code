import argparse

def parse_args():
    ap = argparse.ArgumentParser("Advent Of Code - Day 7")
    ap.add_argument("-f", "--filename", dest="fn", required=False, default="input",
                    help="Specifies which input file to use")
    args = ap.parse_args()

    return args

def get_input(input_file: str):
    with open(input_file, "r") as data_input:
        #return [line.strip() for line in data_input.read().strip().replace("bags","").replace("bag","").split("\n")]
        return [line.strip() for line in data_input.read().strip().replace("bags","").replace("bag","").split("\n")]

def parse_rules(input_rules: list):
    rules = {}
    amount = {}
    for line in input_rules:
        rule = line.strip('.').split(' contain ')
        bag = rule[0][:-1]
        if rule[1] == 'no other ':
            rules[bag] = []
            amount[bag] = []
        else:
            amount[bag] = []
            rules[bag] = []
            for tmp in rule[1].split(', '):
                x = tmp.split()
                amount[bag].append(int(x.pop(0)))
                rules[bag].append(' '.join(x))

    return rules, amount

def bag_can_contain(rules: dict, rule, bag, cache):
    if rule in cache:
        return cache[rule]
    if bag in rules[rule]:
        cache[rule] = True
    else:
        cache[rule] = any(bag_can_contain(rules, b, bag, cache) for b in rules[rule])
    return cache[rule]

def count_bags(rules: dict, amount: dict, bag, cache):
    if bag in cache:
        return cache[bag]

    if len(rules[bag]) == 0:
        cache[bag] = 0
        return 0
    else:
        sum = 0
        for i in range(len(rules[bag])):
            sum += amount[bag][i] * (count_bags(rules, amount, rules[bag][i], cache) + 1)
        cache[bag] = sum
        return cache[bag]

def part1(rules:dict):
    total = 0
    for rule in rules:
        total += bag_can_contain(rules, rule, 'shiny gold', dict())

    return total

def part2(rules: dict, amounts: dict):
    return count_bags(rules, amounts, 'shiny gold', dict())

def main():
    args = parse_args()
    rules, amounts = parse_rules(get_input(args.fn))

    print("Part 1: {}".format(part1(rules)))
    print("Part 2: {}".format(part2(rules, amounts)))

if __name__ == '__main__':
    main()