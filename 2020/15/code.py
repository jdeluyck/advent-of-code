import argparse

def parse_args():
    arg_parse = argparse.ArgumentParser("Advent Of Code - Day 13")
    arg_parse.add_argument("-f", "--filename", dest="fn", required=False, default="input",
                           type=str, help="Specifies which input file to use")
    args = arg_parse.parse_args()

    return args


def get_input(input_file: str):
    with open(input_file, "r") as data_input:
        return [int(line) for line in data_input.read().strip().split(",")]


def do_calc(input_data: list, max_nr: int):
    cache = {}
    cnt = len(input_data) + 1

    for idx in range(len(input_data)-1):
        cache[input_data[idx]] = idx + 1

    last_nr = input_data[-1]
    while cnt <= max_nr:
        if last_nr not in cache:
            new_nr = 0
        else:
            new_nr = cnt - cache[last_nr] - 1

        cache[last_nr] = cnt - 1
        cnt += 1
        last_nr = new_nr

    return last_nr


def part1(input_data: list, max_nr: int = 2020):
    return do_calc(input_data, max_nr)


def part2(input_data: list, max_nr: int = 30000000):
    return do_calc(input_data, max_nr)


def main():
    args = parse_args()
    input = get_input(args.fn)

    print("Part 1: {}".format(part1(input)))
    print("Part 2: {}".format(part2(input)))

if __name__ == '__main__':
    main()