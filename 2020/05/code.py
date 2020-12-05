import argparse

def parse_args():
    ap = argparse.ArgumentParser("Advent Of Code - Day 5")
    ap.add_argument("-f", "--filename", dest="fn", required=False, default="input",
                    help="Specifies which input file to use")
    args = ap.parse_args()

    return args

def get_input(input_file: str):
    seats = list()

    with open(input_file, "r") as data_input:
        for line in data_input.read().splitlines():
            seats.append(line.replace("F", "0").replace("B", "1").replace("L", "0").replace("R", "1"))

    return seats

def define_seat(seat: str):
    row = int(seat[:7], 2)
    col = int(seat[7:], 2)

    return [row, col]

def get_seat_ids(seats: list):
    seat_ids = list()

    for seat in seats:
        seat_pos = define_seat(seat)
        seat_ids.append(seat_pos[0] * 8 + seat_pos[1])

    return seat_ids

def part1(seats: list):
    return max(get_seat_ids(seats))

def part2(seats: list):
    seat_ids = get_seat_ids(seats)

    for seat in range(min(seat_ids), max(seat_ids)):
        if seat not in seat_ids and (seat + 1) in seat_ids and (seat - 1) in seat_ids:
            return seat

def main():
    args = parse_args()
    seats = get_input(args.fn)
    print("Part 1: {}".format(part1(seats)))
    print("Part 2: {}".format(part2(seats)))

if __name__ == '__main__':
    main()
