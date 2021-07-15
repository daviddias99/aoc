FileHandling = require 'fileHandling'

local lines = FileHandling.getLinesFromFile('input/inputA.txt')
local arrivalTime = tonumber(lines[1])
local buses = {}
local iterator = string.gmatch(lines[2], '(%d+)')
local closestBus = 0
local closestBusTime = math.maxinteger

for value in iterator do
  local bus = tonumber(value)
  table.insert(buses, #buses + 1, bus)

  local value = arrivalTime // bus + 1
  if value * bus - arrivalTime < closestBusTime - arrivalTime then
    closestBus = bus
    closestBusTime = value * bus
  end

end

print(closestBus)
print(closestBusTime)
print((closestBusTime - arrivalTime) * closestBus)