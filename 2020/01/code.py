def getinput():
  with open("input", "r") as exp_input:
    return [int(line) for line in exp_input.readlines()]

def part1(items: list):
  for item in items:
    if (2020 - item) in items:
      return [item, (2020 - item), (2020 - item) * item]

def part2(items: list):
  for idx in range(len(items)):
    for idx2 in range(idx, len(items)):
      if (2020 - items[idx] - items[idx2]) in items:
        return [(2020 - items[idx]), (2020-items[idx2]), (2020 - items[idx] - items[idx2]), (2020 - items[idx] - items[idx2]) * items[idx] * items[idx2]]

 
expenses = getinput()
print("Part 1: {}".format(part1(expenses)))
print("Part 2: {}".format(part2(expenses)))

