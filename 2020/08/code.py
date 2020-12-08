import argparse

def parse_args():
    ap = argparse.ArgumentParser("Advent Of Code - Day 8")
    ap.add_argument("-f", "--filename", dest="fn", required=False, default="input",
                    help="Specifies which input file to use")
    args = ap.parse_args()

    return args

def get_input(input_file: str):
    data = []
    values = []
    with open(input_file, "r") as data_input:
        for line in data_input.read().splitlines():
            tmp = line.split()
            data.append(tmp[0])
           values.append(int(tmp[1]))

    return data, values
            
def parse_code(input_code: list, values: list):
    acc = 0
    idx = 0
    executed = []

    while (idx) < len(input_code):
        if idx in executed:
                break
        code = input_code[idx]
        value = values[idx]

        executed.append(idx)

        if code == "nop":
            idx += 1
        elif code == "acc":
            acc += value
            idx += 1
        elif code == "jmp":
            idx += value

    return acc, idx

def part1(codes: list, values: list):
    acc, idx = parse_code(codes, values)
    return acc

def part2(codes: list, values: list):
    for instruction in range(len(codes)):
        if codes[instruction] == "jmp":
            codes[instruction] = "nop"
            acc, idx = parse_code(codes, values)
            codes[instruction] = "jmp"
        elif codes[instruction] == "nop":
            codes[instruction] = "jmp"
            acc, idx = parse_code(codes, values)
            codes[instruction] = "nop"

        if idx == len(codes):
            return acc

def main():
    args = parse_args()
    codes, values = get_input(args.fn)

    print("Part 1: {}".format(part1(codes, values)))
    print("Part 2: {}".format(part2(codes, values)))

if __name__ == '__main__':
    main()
