import argparse


def parse_args():
    arg_parse = argparse.ArgumentParser("Advent Of Code - Day 13")
    arg_parse.add_argument("-f", "--filename", dest="fn", required=False, default="input",
                    type=str, help="Specifies which input file to use")
    args = arg_parse.parse_args()

    return args


def get_input(input_file: str):
    result = []
    with open(input_file, "r") as data_input:
        ts = int(data_input.readline().strip())
        busses = data_input.readline().strip().split(",")
        for bus_id in busses:
            if bus_id != 'x':
                result.append({'id': int(bus_id), 'idx': busses.index(bus_id)})

    return ts, result


def part1(input_data: list, ts: int):
    min_wait = ts
    min_wait_bus_id = None

    for bus in input_data:
        wait = bus['id'] - (ts % bus['id'])
        if wait < min_wait:
            min_wait = wait
            min_wait_bus_id = bus['id']

    return [min_wait, min_wait_bus_id, min_wait * min_wait_bus_id]


def part2(input_data: list):
    ts = 1
    jump = 1

    for bus in input_data:
        while (ts + bus['idx']) % bus['id'] != 0:
            ts += jump

        print ("Bus {} hit 0 at interval {}+{}".format(bus['id'], ts,bus['idx']))
        jump *= bus['id']

    return ts


def main():
    args = parse_args()
    ts, bus_info = get_input(args.fn)

    print("Part 1: {}".format(part1(bus_info, ts)))
    print("Part 2: {}".format(part2(bus_info)))

if __name__ == '__main__':
    main()
