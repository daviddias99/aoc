FileHandling = require 'fileHandling'

function binaryFind(input, low, high, highReducer, lowReducer)
  local currentChr = string.sub(input,1,1)
  local restOfInput = string.sub(input,2)
  if high - low <= 0 then
    return low, input
  end

  if currentChr == highReducer then
    high = low + (high - low)//2
  elseif currentChr == lowReducer then
    low = low + (high - low)//2 + 1   
  end
  
  return binaryFind(restOfInput, low, high, highReducer, lowReducer)
end

function getSeat(line)
  local row, rest = binaryFind(line, 0, 127, 'F', 'B')
  local column = binaryFind(rest, 0, 7, 'L', 'R')
  return row, column
end

function getScore(row, column)
  return row * 8 + column
end

function maxScore(lines)
  local max = 0
  local scores = {}
  for _, value in pairs(lines) do
    local row, column = getSeat(value)
    local score = getScore(row, column)

    table.insert(scores, #scores + 1, score)

    if score > max then
      max = score
    end
  end
  table.sort(scores)
  return max, scores
end
local lines = FileHandling.getLinesFromFile("input/inputA.txt")

local max, scores = maxScore(lines)

local old = nil
for key, value in pairs(scores) do
  if old ~= nil and value - old > 1 then
      print(string.format('%d %d %d', key, value, old))
  end

  old = value
end
