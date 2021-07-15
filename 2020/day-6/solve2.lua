FileHandling = require 'fileHandling'

local function countChars(line)
  local charCount = {}
  local personCount = 1
  local count = 0

  for i = 1, #line do
    local c = line:sub(i,i)

    if c == ' ' then
      personCount = personCount + 1
    elseif charCount[c] == nil then
      charCount[c] = 1
    else
      charCount[c] = charCount[c] + 1
    end
  end

  for _, value in pairs(charCount) do
    if value == personCount then
      count = count + 1
    end  
  end

  return count
end

local lines = FileHandling.getGroupedLinesFromFile("input/inputA.txt", false)

local count = 0
for _, value in pairs(lines) do
  count = count + countChars(value)
end

print(count)