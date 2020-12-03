import math

def getinput():
  with open("input", "r") as map_input:
    orig_map = map_input.read().splitlines()

  return orig_map

def scanmap(map: list, right: int, down: int):
  trees = colidx = 0

  for line in map[::down]:
    full_line = line

    if colidx >= len(line):
      full_line = line * math.ceil(colidx / len(line)+1)

    if full_line[colidx] == "#":
      trees += 1
    colidx += right

  return trees

def part1(map: list):
  return scanmap(map, 3, 1)

def part2(policy: list):
  case1 = scanmap(map, 1, 1)
  case2 = scanmap(map, 3, 1)
  case3 = scanmap(map, 5, 1)
  case4 = scanmap(map, 7, 1)
  case5 = scanmap(map, 1, 2)

  return [case1, case2, case3, case4, case5, (case1 * case2 * case3 * case4 * case5)]

map = getinput()
print("Part 1: {}".format(part1(map)))
print("Part 2: {}".format(part2(map)))
