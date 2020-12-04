import cerberus

required_fields = ['byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid']
optional_fields = ['cid']

validation_schema = {
  'byr': { 'type': 'integer', 'min': 1920, 'max': 2002, 'coerce': int, 'required': True },
  'iyr': { 'type': 'integer', 'min': 2010, 'max': 2020, 'coerce': int, 'required': True },
  'eyr': { 'type': 'integer', 'min': 2020, 'max': 2030, 'coerce': int, 'required': True },
  'hgt': { 'type': 'string', 'required': True, 'oneof': [ {'regex': '^1[5-8][0-9]cm'}, {'regex': '^19[0-3]cm'},
                                      {'allowed': ['59in']}, {'regex': '^6[0-9]in'}, { 'regex': '^7[0-6]in'} ],
  },
  'hcl': { 'type': 'string', 'regex': '^#[a-f,0-9]{6}', 'required': True },
  'ecl': { 'type': 'string', 'allowed': ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth'], 'required': True },
  'pid': { 'type': 'string', 'regex': '^\d{9}', 'required': True },
  'cid': { 'required': False }
}

def getinput():
  passports = list()
  tmp = list()

  with open("input", "r") as data_input:
    for line in data_input.read().splitlines():
      if len(line) == 0:
        passports.append(dict(tmp))
        tmp = list()
        continue

      for chunk in line.split():
        tmp.append(map(str.strip, chunk.split(':')))

  passports.append(dict(tmp))
       
  return passports

def validate(passports: list, required_fields: list, optional_fields: list, validation_schema: dict):
  valid = invalid = 0

  if validation_schema:
    validator = cerberus.Validator(validation_schema)  
  else:
    all_fields = required_fields + optional_fields

  for passport in passports:
    if validation_schema:
      if validator.validate(passport):
        valid += 1
      else:
        invalid += 1

    else:
      pp_set = set(list(passport.keys()))

      if pp_set == set(all_fields): 
        valid += 1
      else:
        if pp_set == set(required_fields):
          valid += 1
        else:
          invalid += 1
  
  return [valid, invalid]

def part1(passports: list, required_fields: list, optional_fields: list):
  return validate(passports, required_fields, optional_fields, None)
 
def part2(passports: list, validation_schema: dict):
  return validate(passports, None, None, validation_schema)

passports = getinput()

print("Part 1: {}".format(part1(passports, required_fields, optional_fields)))
print("Part 2: {}".format(part2(passports, validation_schema)))
