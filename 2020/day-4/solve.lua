FileHandling = require 'fileHandling'

function checkValid1(password)
  local byr = string.match(password, 'byr:') ~= nil
  local iyr = string.match(password, 'iyr:') ~= nil
  local eyr = string.match(password, 'eyr:') ~= nil
  local hgt = string.match(password, 'hgt:') ~= nil
  local hcl = string.match(password, 'hcl:') ~= nil
  local ecl = string.match(password, 'ecl:') ~= nil
  local pid = string.match(password, 'pid:') ~= nil
  return byr and iyr and eyr and hgt and hcl and ecl and pid
end

function checkValid2(password)
  local byr = string.match(password, 'byr:(%d%d%d%d)')
  local byrValid = byr and (tonumber(byr) >= 1920 and tonumber(byr) <= 2002)

  local iyr = string.match(password, 'iyr:(%d%d%d%d)')
  local iyrValid = iyr and (tonumber(iyr) >= 2010 and tonumber(iyr) <= 2020)

  local eyr = string.match(password, 'eyr:(%d%d%d%d)')
  local eyrValid = eyr and (tonumber(eyr) >= 2020 and tonumber(eyr) <= 2030)

  local hgt, unit = string.match(password, 'hgt:(%d+)(cm)')
  if hgt == nil then
    hgt, unit = string.match(password, 'hgt:(%d+)(in)')
  end
  local hgtValid = hgt ~= nil

  if hgt then
    if unit == 'cm' then
      hgtValid = tonumber(hgt) >= 150 and tonumber(hgt) <= 193
    else
      hgtValid = tonumber(hgt) >= 59 and tonumber(hgt) <= 76
    end
  end
  local hcl = string.match(password, 'hcl:#[0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f]')
  local hclValid = hcl ~= nil

  local ecl = string.match(password, 'ecl:(%a+)')
  local eclValid = (ecl == 'amb') or (ecl == 'blu') or (ecl == 'brn') or (ecl == 'gry') or (ecl == 'grn') or (ecl == 'hzl') or (ecl == 'oth')

  local pid = string.match(password, 'pid:%d%d%d%d%d%d%d%d%d%s')
  if pid == nil then
    pid = string.match(password, 'pid:%d%d%d%d%d%d%d%d%d$')
  end
  
  local pidValid = pid ~= nil
  return byrValid and iyrValid and eyrValid and hgtValid and hclValid and eclValid and pidValid
end

local lines = FileHandling.getGroupedLinesFromFile("input/inputA.txt")
local result = 0
for _, value in pairs(lines) do
  if(checkValid2(value)) then
    result = result + 1
  end
end

print(result)