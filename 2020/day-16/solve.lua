FileHandling = require 'fileHandling'

local lines = FileHandling.getLinesFromFile('input/inputA.txt')
local intervals = {}
local tickets = {}
local myTicket = {}
local idx

for i = 1, #lines do
  idx = i
  if lines[i] == '' then
    break
  end
  local name, i1a, i1b, i2a, i2b = string.match(lines[i], '(.+): (%d+)-(%d+) or (%d+)-(%d+)')
  table.insert(intervals, #intervals + 1, {tonumber(i1a), tonumber(i1b)})
  table.insert(intervals, #intervals + 1, {tonumber(i2a), tonumber(i2b)})
end

for i = idx + 2, #lines do
  idx = i
  if lines[i] == '' then
    break
  end

  for value in string.gmatch(lines[i], '(%d+)') do
    table.insert(myTicket, #myTicket + 1, tonumber(value))
  end
end

for i = idx + 2, #lines do
  idx = i
  if lines[i] == '' then
    break
  end

  local currentNums = {}
  for value in string.gmatch(lines[i], '(%d+)') do
    table.insert(currentNums, #currentNums + 1, tonumber(value))
  end

  table.insert(tickets, #tickets + 1, currentNums)
end

local sum = 0
local ticektsToRemove = {}

for key, ticket in pairs(tickets) do
  for _, number in pairs(ticket) do
      local numberValid = false
      for _, interval in pairs(intervals) do
        if number >= interval[1] and number <= interval[2] then
            numberValid = true
            break
        end
      end

      if not numberValid then
        sum = sum + number
      end
  end
end

print(sum)