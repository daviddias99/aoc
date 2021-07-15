FileHandling = require 'fileHandling'

local lines = FileHandling.getLinesFromFile('input/inputA.txt')
local numbers = {}
local i = 1
local last
for number in string.gmatch(lines[1], '(%d+)') do
  numbers[tonumber(number)] = {i, nil}
  i = i + 1
  last = tonumber(number)
end
local spoken
for i = i, 30000000 do
  spoken = nil
  if numbers[last][2] == nil then
    spoken = 0
  else
    spoken = numbers[last][1] - numbers[last][2]
  end

  if numbers[spoken] == nil then
    numbers[spoken] = {i, nil}
  else
    numbers[spoken][2] = numbers[spoken][1]
    numbers[spoken][1] = i
  end
  last = spoken
  -- print(spoken)
end

print(spoken)


