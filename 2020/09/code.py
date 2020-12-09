import argparse

def parse_args():
    ap = argparse.ArgumentParser("Advent Of Code - Day 9")
    ap.add_argument("-f", "--filename", dest="fn", required=False, default="input",
                    type=str, help="Specifies which input file to use")
    ap.add_argument("-p", "--preamble", dest="preamble", required=False, default=25,
                    type=int, help="Preamble for calculations")
    ap.add_argument("-l", "--length", dest="length", required=False, default=25,
                    type=int, help="Length for calculations")
    args = ap.parse_args()

    return args

def get_input(input_file: str):
    with open(input_file, "r") as data_input:
        return [int(line) for line in data_input.readlines()]

def find_value(input: list, preamble: int, length: int):
    nomatchlist = []

    for idx in range(preamble, len(input)):
        nomatch = True

        tmp = input[idx-length:idx]

        for tmp_val in tmp:
            if (input[idx] - tmp_val) in tmp and ((input[idx] - tmp_val) != tmp_val):
                nomatch = False

        if nomatch:
            nomatchlist.append(input[idx])

    return nomatchlist

def find_min_max(input: list, total: int):
    for preamble in range(0, len(input)):
        for length in range(2, len(input)):
            tmp = input[preamble:preamble+length]
            if sum(tmp) == total:
                return [min(tmp), max(tmp), min(tmp) + max(tmp)]
            elif sum(tmp) > total:
                continue

def part1(input: list, preamble: int, length: int):
    return(find_value(input, preamble, length)[0])

def part2(input: list, preamble: int, length: int):
    total = find_value(input, preamble, length)[0]
    return (find_min_max(input, total))

def main():
    args = parse_args()
    input = get_input(args.fn)

    print("Part 1: {}".format(part1(input, args.preamble, args.length)))
    print("Part 2: {}".format(part2(input, args.preamble, args.length)))

if __name__ == '__main__':
    main()
