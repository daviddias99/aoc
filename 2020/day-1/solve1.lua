FileHandling = require 'fileHandling'

local function filterAndSplit(list, test) 

  local even = {}
  local odd = {}

  for _, value in pairs(list) do
    if value > test then
      goto skip
    elseif value % 2 == 0 then
      table.insert(even, value)
    else
      table.insert(odd, value)
    end
    ::skip::
  end

  return even, odd
end

local function checkValues(list, value)
  for i = 1, #list - 1 do
    for j = i + 1, #list do
      if list[i] + list[j] == value then
        return list[i], list[j]
      end
    end
  end

  return 0,0
end

-- tests the functions above

local file = 'input/input_large.txt'
local lines = FileHandling.getNumbersFromFile(file)
local even, odd = filterAndSplit(lines, 2020 - math.min(table.unpack(lines)))

local value1, value2 = checkValues(even, 2020)
print(string.format("%d + %d, %d", value1, value2, value1 * value2))

if value1 * value2 ~= 0 then
  return
end

value1, value2 = checkValues(odd, 2020)
print(string.format("%d + %d, %d", value1, value2, value1 * value2))