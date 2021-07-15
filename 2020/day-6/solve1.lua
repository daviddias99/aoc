FileHandling = require 'fileHandling'

local function countChars(line)
  local count = 0
  local charCount = {}
  for i = 1, #line do
    local c = line:sub(i,i)

    if charCount[c] == nil then
      count = count + 1
      charCount[c] = 1
    end
  end

  return count
end

local lines = FileHandling.getGroupedLinesFromFile("input/inputA.txt", true)

local count = 0
for _, value in pairs(lines) do
  count = count + countChars(value)
end

print(count)