FileHandling = require 'fileHandling'

local function anyThatAddTo(table, value, startIdx, endIdx)
  for i = startIdx, endIdx - 1 do
    for j = i + 1, endIdx do
        if table[j] == value - table[i] then
          return true
        end
    end  
  end

  return false
end

local lines = FileHandling.getNumbersFromFile('input/inputA.txt')
local windowSize = 25
local neededNumber = 0

for i = 1, #lines - windowSize do
  if not anyThatAddTo(lines, lines[i+windowSize],i, i + windowSize - 1) then
    neededNumber = lines[i+windowSize]
    print(neededNumber)
    break
  end
end

local currentSum = lines[1]
local start = 1;
local endidx;

for i = 2, #lines do

  while currentSum > neededNumber and start < i - 1 do
    currentSum = currentSum - lines[start]
    start = start + 1  
  end

  if currentSum == neededNumber then
    endidx = i - 1
    break
  end

  if i <= #lines then
      currentSum = currentSum + lines[i]
  end

end

local min = math.maxinteger
local max = math.mininteger

for i = start, endidx do
  if lines[i] < min then
    min = lines[i]
  end

  if lines[i] > max then
    max = lines[i]
  end
end

print(max + min)
