import argparse
from itertools import product

def parse_args():
    arg_parse = argparse.ArgumentParser("Advent Of Code - Day 17")
    arg_parse.add_argument("-f", "--filename", dest="fn", required=False, default="input",
                           type=str, help="Specifies which input file to use")
    args = arg_parse.parse_args()

    return args


def get_input(input_file: str):
    cubes = set()

    with open(input_file, "r") as data_input:
        data_file = data_input.read().splitlines()
        col_len = len(data_file[0])
        for line in range(len(data_file)):
            for col in range(col_len):
                if data_file[col][line] == "#":
                    cubes.add((line, col, 0, 0))

        return cubes


def run_simulation(cubes: set, grid_size: int, cycles: int = 6):
    positions = product(range(-1, 2), repeat=grid_size)

    cnt_pos = []
    if grid_size == 3:
        for tmp in positions:
            if tmp != (0, 0, 0):
                cnt_pos.append((tmp[0], tmp[1], tmp[2], 0))
    else:
        for tmp in positions:
            if tmp != (0, 0, 0, 0):
                cnt_pos.append(tmp)

    cnt_pos = tuple(cnt_pos)

    for _ in range(cycles):
        next_cube = set()
        neighbours = set()

        for cube in cubes:
            cnt = 0
            for n in cnt_pos:
                tmp = []
                for i in range(4):
                    tmp.append((cube[i] + n[i]))
                tmp = tuple(tmp)

                if tmp in cubes:
                    cnt += 1
                else:
                    neighbours.add(tmp)
            if 2 <= cnt <= 3:
                next_cube.add(cube)

        for cube in neighbours:
            cnt = 0
            for n in cnt_pos:
                tmp = []
                for i in range(4):
                    tmp.append((cube[i] + n[i]))
                tmp = tuple(tmp)

                if tmp in cubes:
                    cnt += 1

            if cnt == 3:
                next_cube.add(cube)

        cubes = next_cube

    return len(cubes)


def part1(cubes: set, grid_size: int):
    return run_simulation(cubes, grid_size)


def part2(cubes: set, grid_size: int):
    return run_simulation(cubes, grid_size)


def main():
    args = parse_args()
    cubes = get_input(args.fn)

    print("Part 1: {}".format(part1(cubes, 3)))
    print("Part 2: {}".format(part2(cubes, 4)))


if __name__ == '__main__':
    main()
