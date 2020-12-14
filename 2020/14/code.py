import argparse


def parse_args():
    arg_parse = argparse.ArgumentParser("Advent Of Code - Day 14")
    arg_parse.add_argument("-f", "--filename", dest="fn", required=False, default="input",
                    type=str, help="Specifies which input file to use")
    args = arg_parse.parse_args()

    return args


def get_input(input_file: str):
    with open(input_file, "r") as data_input:
        return data_input.read().splitlines()


def run_part1(input_data: list):
    mem = {}
    mask = ""

    for line in input_data:
        cmd, value = line.split(" = ")

        if cmd == "mask":
            mask = value.strip()
        else:
            print(value)
            address = int(cmd[4:-1])
            value = int(value)
            mem[address] = apply_mask_value(mask, value)

    return mem


def run_part2(input_data: list):
    mem = {}
    mask = ""

    for line in input_data:
        cmd, value = line.split(" = ")

        if cmd == "mask":
            mask = value.strip()
        else:
            address = int(cmd[4:-1])
            value = int(value)
            all_mem_addresses = calc_mask_address(mask, address)
            for idx in all_mem_addresses:
                mem[idx] = value

    return mem


def apply_mask_value(mask: str, value: int):
    or_mask = int(mask.replace('X', '0'), 2)
    and_mask = int(mask.replace('X', '1'), 2)

    result = (value | or_mask) & and_mask
    return result


def calc_mask_address(mask: str, address: int):
    tmp = int(mask.replace('X', '0'), 2)
    address |= tmp

    floaties = []
    for idx, char in enumerate(mask):
        if char == "X":
            floaties.append(len(mask) - 1 - idx)

    return do_flip(address, floaties)


def do_flip(address: int, floaties: list):
    if len(floaties) == 1:
        tmp1 = address | (1 << floaties[0])
        tmp2 = address & ~(1 << floaties[0])

        return [tmp1, tmp2]

    tmp1 = address | (1 << floaties[0])
    tmp2 = address & ~(1 << floaties[0])
    subfloaties = floaties[1:]

    tmp = do_flip(tmp1, subfloaties) + do_flip(tmp2, subfloaties)

    return tmp


def part1(input_data: list):
    mem = run_part1(input_data)
    return sum([val for val in mem.values()])


def part2(input_data: list):
    mem = run_part2(input_data)
    return sum([val for val in mem.values()])


def main():
    args = parse_args()
    data = get_input(args.fn)

    print("Part 1: {}".format(part1(data)))
    print("Part 2: {}".format(part2(data)))


if __name__ == '__main__':
    main()
