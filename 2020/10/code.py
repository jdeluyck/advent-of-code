import argparse

def parse_args():
    arg_parse = argparse.ArgumentParser("Advent Of Code - Day 10")
    arg_parse.add_argument("-f", "--filename", dest="fn", required=False, default="input",
                    type=str, help="Specifies which input file to use")
    args = arg_parse.parse_args()

    return args

def get_input(input_file: str):
    with open(input_file, "r") as data_input:
        return sorted([int(line) for line in data_input.readlines()])

def calculate_adapter_chain(input_data: list):
    jolts = 0
    count = [0, 0, 0]

    for adapter_jolts in input_data:
        count[adapter_jolts - jolts - 1] += 1
        jolts = adapter_jolts

    return [count[0], count[2], count[0] * count[2]]

def determine_possibilities(input_data: list, jolts: int, cache):
    if jolts in cache:
        return cache[jolts]
    tmp = []
    for idx in range(jolts + 1, jolts + 4):
        if idx in input_data:
            tmp.append(determine_possibilities(input_data, idx, cache))

    cache[jolts] = sum(tmp)
    return cache[jolts]

def part1(input_data: list):
    return calculate_adapter_chain(input_data)

def part2(input_data: list):
    cache = {input_data[-1]: 1}
    determine_possibilities(input_data, 0, cache)
    return cache[0]

def main():
    args = parse_args()
    input_data = get_input(args.fn)

    print("Part 1: {}".format(part1(input_data)))
    print("Part 2: {}".format(part2(input_data)))

if __name__ == '__main__':
    main()
