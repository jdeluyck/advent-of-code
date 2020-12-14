import argparse
import pprint
import copy


def parse_args():
    arg_parse = argparse.ArgumentParser("Advent Of Code - Day 11")
    arg_parse.add_argument("-f", "--filename", dest="fn", required=False, default="input",
                    type=str, help="Specifies which input file to use")
    args = arg_parse.parse_args()

    return args


def get_input(input_file: str):
    data = []
    with open(input_file, "r") as data_input:
        for line in data_input.read().splitlines():
            data.append(list(line))

    return data


def check_seat(input_data: list, x: int, y: int, max_taken: int):
    cnt_taken_adj = cnt_taken_vis = 0
    result_adj = result_vis = seat = input_data[x][y]

    if seat == ".":
        return [seat, seat]

    for x_iter in range(x-1, x+2):
        if x_iter < 0 or x_iter >= len(input_data):
            continue
        for y_iter in range(y-1, y+2):
            if y_iter < 0 or y_iter >= len(input_data[x_iter]):
                continue
            if x_iter == x and y_iter == y:
                continue

            adjacent_seat = input_data[x_iter][y_iter]

            if adjacent_seat == "#":
                cnt_taken_adj += 1
                cnt_taken_vis += 1
            elif adjacent_seat == "L":
                pass
            else:
                diff_x = x_iter - x
                diff_y = y_iter - y

                vis_x = x_iter + diff_x
                vis_y = y_iter + diff_y

                while vis_x >= 0 and vis_x < len(input_data) and vis_y >= 0 and vis_y < len(input_data[vis_x]):
                    visible_seat = input_data[vis_x][vis_y]
                    vis_x += diff_x
                    vis_y += diff_y

                    if visible_seat == "#":
                        cnt_taken_vis += 1
                        break
                    elif visible_seat == "L":
                        break

    if seat == "L":
        if cnt_taken_adj == 0:
            result_adj = "#"
        if cnt_taken_vis == 0:
            result_vis = "#"
    elif seat == "#":
        if cnt_taken_adj >= max_taken:
            result_adj = "L"
        if cnt_taken_vis >= max_taken:
            result_vis = "L"

    return [result_adj, result_vis]


def iterate_seats(input_data: list, max_taken: int, adjacent: bool = True):
    seat_chart = input_data

    while True:
        old_seat_chart = copy.deepcopy(seat_chart)
        changed = False

        for x_idx in range (0, len(seat_chart)):
            for y_idx in range(0, len(seat_chart[x_idx])):
                adj, vis = check_seat(old_seat_chart, x_idx, y_idx, max_taken)

                if adjacent:
                    tmp = adj
                else:
                    tmp = vis

                if tmp != seat_chart[x_idx][y_idx]:
                    seat_chart[x_idx][y_idx] = tmp
                    changed = True

        if not changed:
            break

    return seat_chart


def count_seats(input_data: list, seat_type: str):
    cnt = 0
    for x_idx in range(0, len(input_data)):
        for y_idx in range(0, len(input_data[x_idx])):
            if input_data[x_idx][y_idx] == seat_type:
                cnt += 1

    return cnt


def part1(input_data: list):
    return count_seats(iterate_seats(input_data, 4), "#")


def part2(input_data: list):
    return count_seats(iterate_seats(input_data, 5, False), "#")


def main():
    args = parse_args()
    input_data = get_input(args.fn)
    print("Part 1: (should be 37/2412) {}".format(part1(input_data)))

    input_data = get_input(args.fn)
    print("Part 2: {}".format(part2(input_data)))

if __name__ == '__main__':
    main()
