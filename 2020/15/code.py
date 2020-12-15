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
    cnt = 0
    spoken_nrs = []

    for idx, nr in enumerate(input_data):
        tmp = [idx]
        cache[nr] = tmp
        spoken_nrs.append(nr)
        cnt += 1

    while cnt < max_nr:
        last_nr = spoken_nrs[-1]
        if last_nr in cache and len(cache[last_nr]) == 1:
                new_nr = 0
        else:
            new_nr = cache[last_nr][-1] - cache[last_nr][-2]

        spoken_nrs.append(new_nr)

        if new_nr in cache:
            tmp = [cache[new_nr][-1], len(spoken_nrs) - 1]
        else:
            tmp = [len(spoken_nrs) - 1]

        cache[new_nr] = tmp
        cnt += 1

    return spoken_nrs[-1]


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