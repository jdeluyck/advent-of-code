import argparse
from collections import deque


def parse_args():
    arg_parse = argparse.ArgumentParser("Advent Of Code - Day 12")
    arg_parse.add_argument("-f", "--filename", dest="fn", required=False, default="input",
                    type=str, help="Specifies which input file to use")
    args = arg_parse.parse_args()

    return args


def get_input(input_file: str):
    data = []
    values = []
    with open(input_file, "r") as data_input:
        for line in data_input.read().splitlines():
            data.append(line[0])
            values.append(int(line[1:]))

    return data, values


def calc_distance(orders: list, amount: list, start_direction: str, wp_mode: bool = False):
    ship_distance = {"N": 0, "E": 0, "S": 0, "W": 0}
    wp_distance = {"N": 1, "E": 10, "S": 0, "W": 0}

    ship_direction = start_direction

    for idx, order in enumerate(orders):
        if order == "L" or order == "R":
            if wp_mode:
                wp_distance = rotate_waypoint(wp_distance, order, amount[idx])
            else:
                ship_direction = calc_direction(ship_direction, order, amount[idx])
        else:
            if wp_mode:
                if order == "F":
                    for k,v in wp_distance.items():
                        ship_distance = move_around(ship_distance, k, v * amount[idx])
                else:
                    wp_distance = move_around(wp_distance, order, amount[idx])
            else:
                if order == "F":
                    order = ship_direction

                ship_distance = move_around(ship_distance, order, amount[idx])

    ship_distance['Total'] = sum([i for i in ship_distance.values()])

    return ship_distance


def move_around(distances: dict, direction: str, amount: int):
    rev = calc_direction(direction, "R", 180)
    tmp = distances[direction] - distances[rev] + amount

    if tmp < 0:
        distances[rev] = abs(tmp)
    else:
        distances[rev] = 0
        distances[direction] = tmp

    return distances


def rotate_waypoint(waypoints: dict, direction: str, amount: int):
    compass = ["N", "E", "S", "W"]
    tmp = deque(waypoints.values())

    if direction == "L":
        amount = 0 - int(amount / 90)
    else:
        amount = int(amount / 90)

    tmp.rotate(amount)

    for c in compass:
        waypoints[c] = tmp[compass.index(c)]

    return waypoints


def calc_direction(start: str, direction: str, amount: int):
    compass = ["N", "E", "S", "W"]

    if direction == "L":
        compass.reverse()

    idx = compass.index(start)

    return compass[(idx + int(amount / 90)) % len(compass)]


def part1(input_data: list, amount: list):
    return calc_distance(input_data, amount, "E")


def part2(input_data: list, amount: list):
    return calc_distance(input_data, amount, "E", True)


def main():
    args = parse_args()
    direction, amount = get_input(args.fn)

    print("Part 1: {}".format(part1(direction, amount)))
    print("Part 2: {}".format(part2(direction, amount)))

if __name__ == '__main__':
    main()
