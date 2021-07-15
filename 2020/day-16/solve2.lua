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
  intervals[name] = {{tonumber(i1a), tonumber(i1b)}, {tonumber(i2a), tonumber(i2b)}}
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

local function isInInterval(interval, number)
  if number >= interval[1][1] and number <= interval[1][2] then
    return true
  elseif number >= interval[2][1] and number <= interval[2][2] then
    return true
  end

  return false
end

local filteredTickets = {}

for _, ticket in pairs(tickets) do
  local ticketValid = true
  for _, number in pairs(ticket) do
      local numberValid = false
      for _, interval in pairs(intervals) do
        numberValid = isInInterval(interval, number)
        if numberValid then break end
      end

      if not numberValid then
        ticketValid = false
        break
      end
  end

  if ticketValid then
    table.insert(filteredTickets, #filteredTickets + 1, ticket)
  end
end

local candidates = {}

for i = 1, #filteredTickets[1] do
  for intervalName, interval in pairs(intervals) do
    local validIntervalForColummn = true
    for _, ticket in pairs(filteredTickets) do
      if not isInInterval(interval, ticket[i]) then
        validIntervalForColummn = false
        break
      end
    end

    if validIntervalForColummn then
      if candidates[i] == nil then
        candidates[i] = {intervalName}
      else
        table.insert(candidates[i], #candidates[i] + 1, intervalName)
      end
    end
  end
end

local function refactorCandidate(candidate, nameToRemove)
  local result = {}
  local changed = false

  for _, name in pairs(candidate) do
    if name ~= nameToRemove then
      table.insert(result, #result + 1, name)
    else
      changed = true
    end

  end

  return result, changed
end

local changed = false

repeat
  changed = false
  for key, candidate in pairs(candidates) do
    if #candidate == 1 then
      for candidateSweepKey, candidateSweepValue in pairs(candidates) do
        if candidateSweepKey ~= key then
          local didChange
          candidates[candidateSweepKey], didChange =  refactorCandidate(candidateSweepValue, candidate[1])
          changed = didChange or changed
        end
      end
    end
  end
until not changed

local interestingIndexes = {}
for key, value in pairs(candidates) do
  if(string.match(value[1], 'departure') ~= nil) then
    table.insert(interestingIndexes, #interestingIndexes + 1, key)
  end
end

local result = 1

for key, value in pairs(interestingIndexes) do
  result = result * myTicket[value]
end

print(result)