def getinput():
  policy = list()

  with open("input", "r") as policy_input:
    for line in policy_input.readlines():
      tmp = line.split()
      nrs = tmp[0].split("-")
      #              firstnr,       secondnr,    match letter,    string
      policy.append((int(nrs[0]), int(nrs[1]), tmp[1][:1],        tmp[2]))
  return policy

def part1(policy: list):
  count = 0
  for item in policy:
    min, max, letter, string = item
    if string.count(letter) >= min and string.count(letter) <= max:
      count += 1
  return count

def part2(policy: list):
  count = 0
  for item in policy:
    pos1, pos2, matchletter, string = item
    if (string[pos1 - 1] == matchletter) ^ (string[pos2 - 1] == matchletter):
        count += 1
  return count

policy = getinput()
print("Part 1: {}".format(part1(policy)))
print("Part 2: {}".format(part2(policy)))
