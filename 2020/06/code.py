import argparse

def parse_args():
    ap = argparse.ArgumentParser("Advent Of Code - Day 6")
    ap.add_argument("-f", "--filename", dest="fn", required=False, default="input",
                    help="Specifies which input file to use")
    args = ap.parse_args()

    return args

def get_input(input_file: str):
    data = list()
    tmp = list()

    with open(input_file, "r") as data_input:
        for line in data_input.read().splitlines():
            if len(line) == 0:
                data.append(tmp)
                tmp = list()
                continue

            tmp.append(line)

    data.append(tmp)

    return data

def count_answers(answers: list, everybody: bool):
    cnt = 0
    matrix = dict()

    for group in answers:
        for letter in range(ord('a'), ord('z')+1):
            matrix[chr(letter)] = 0

        for answers in group:
            for letter in answers:
                matrix[letter] += 1

        if everybody:
            cnt += sum([1 for letter in matrix if matrix[letter] == len(group)])
        else:
            cnt += sum([1 for letter in matrix if matrix[letter]])

    return cnt

def part1(answers: list):
    return count_answers(answers, False)

def part2(answers: list):
    return count_answers(answers, True)

def main():
    args = parse_args()
    answers = get_input(args.fn)
    print("Part 1: {}".format(part1(answers)))
    print("Part 2: {}".format(part2(answers)))

if __name__ == '__main__':
    main()
