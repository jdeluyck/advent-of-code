import argparse
from collections import Counter

def parse_args():
    arg_parse = argparse.ArgumentParser("Advent Of Code - Day 16")
    arg_parse.add_argument("-f", "--filename", dest="fn", required=False, default="input",
                    type=str, help="Specifies which input file to use")
    args = arg_parse.parse_args()

    return args


def get_input(input_file: str):
    rules = {}
    your_ticket = []
    nearby_tickets = []

    with open(input_file, "r") as data_input:
        data_file = data_input.read().splitlines()
        idx = 0
        while idx in range(len(data_file)):
            tmp = data_file[idx].split(':')
            if tmp[0] == 'your ticket':
                idx += 1
                your_ticket = [int(i) for i in data_file[idx].split(',')]
            elif tmp[0] == 'nearby tickets':
                idx += 1
                while idx < len(data_file) and data_file[idx] != "":
                    nearby_tickets.append([int (i) for i in data_file[idx].split(',')])
                    idx += 1
            elif len(tmp) > 1 and tmp[0] != '':
                rules[tmp[0]] = split_input(tmp[1])

            idx += 1

    return rules, your_ticket, nearby_tickets


def split_input(input_data: str):
    tmp = []
    input_data = input_data.replace(" or ", "|").strip()
    for chunk in input_data.split("|"):
        chunk2 = chunk.split("-")
        tmp.append(list(range(int(chunk2[0]), int(chunk2[1]) + 1)))

    tmp = [i for j in tmp for i in j]
    return tmp


def filter_valid_tickets(rules: dict, nearby_tickets: list, your_ticket: list):
    result = 0
    tmp = []
    good_tickets = []

    for x in rules.values():
        tmp += x

    nearby_tickets.append(your_ticket)

    for ticket in nearby_tickets:
        tmp2 = set(ticket) - set(tmp)

        if tmp2:
            result += sum(tmp2)
        else:
            good_tickets.append(ticket)

    return good_tickets, result


def part1(rules: dict, nearby_tickets: list, your_ticket: list):
    valid_tickets, result = filter_valid_tickets(rules, nearby_tickets, your_ticket)
    return result


def find_pos(rules: dict, nearby_tickets: list):
    order_idx = []
    order_name = []
    cnt = {}
    ticket_cnt = len(nearby_tickets)

    for ticket in nearby_tickets:
        for idx, ticket_field_value in enumerate(ticket):
            foo = 0
            for rule_key, rule_value in rules.items():
                if rule_key not in cnt:
                    cnt[rule_key] = {}

                if ticket_field_value in rule_value:
                    if idx in cnt[rule_key]:
                        cnt[rule_key][idx] += 1
                    else:
                        cnt[rule_key][idx] = 1

                    foo += 1

    while len(cnt) > 0:
        for rule_key, rule_value in cnt.items():
            tmp = Counter(rule_value.values())[ticket_cnt]

            if tmp == 1:
                val = list(rule_value.keys())[list(rule_value.values()).index(ticket_cnt)]
                order_idx.append(val)
                order_name.append(rule_key)

                del cnt[rule_key]

                for x in cnt:
                    if val in cnt[x]:
                        del cnt[x][val]

                break


    return order_name, order_idx


def part2(rules: dict, good_tickets: list, your_ticket: list):
    total = 1
    order_name, order_idx = find_pos(rules, good_tickets)

    for idx, order in enumerate(order_name):
        if 'departure' in order:
            total *= your_ticket[order_idx[idx]]

    return total

def main():
    args = parse_args()
    rules, your_ticket, nearby_tickets = get_input(args.fn)

    good_tickets, results = filter_valid_tickets(rules, nearby_tickets, your_ticket)

    print("Part 1: {}".format(results))
    print("Part 2: {}".format(part2(rules, good_tickets, your_ticket)))

if __name__ == '__main__':
    main()
