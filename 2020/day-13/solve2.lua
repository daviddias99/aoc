FileHandling = require 'fileHandling'

local lines = FileHandling.getLinesFromFile('input/inputA.txt')
local buses = {}
local iterator = string.gmatch(lines[2], '(%w+)')

local function findFrom(buses, idx, number)

  if idx > #buses then
    return true
  end

  if buses[idx] < 0 then
    return findFrom(buses, idx + 1, number - buses[idx])
  end

  if number % buses[idx] == 0 then
    return findFrom(buses, idx + 1 , number + 1)
  end

  return false
end

local maxNumber = 0
local maxIndex = -1
local xIterator = 0
for value in iterator do
  local bus
  if value == 'x' then
    -- bus = 'x'
    xIterator = xIterator + 1
  else

    if xIterator ~= 0 then
      table.insert(buses, #buses + 1, -xIterator)
      xIterator = 0
    end

    bus = tonumber(value)
    if bus > maxNumber then
      maxNumber = bus
      maxIndex = #buses + 1
    end
  end
  table.insert(buses, #buses + 1, bus)
end

for key, value in pairs(buses) do
  print(value)
end

local currentIdx = (100000000000000 // maxNumber) * maxNumber - maxIndex + 1

while true do
  if currentIdx % buses[1] == 0 and findFrom(buses, 2, currentIdx + 1) then
    break
  end
  currentIdx = currentIdx + maxNumber - maxIndex + 1
end

print(currentIdx)
