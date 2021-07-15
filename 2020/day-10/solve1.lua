FileHandling = require 'fileHandling'

local lines = FileHandling.getNumbersFromFile('input/inputA.txt')
table.insert(lines, #lines + 1, 0)
table.sort(lines)
table.insert(lines, #lines + 1, lines[#lines] + 3)

local differences = {0, 0, 0}

for i = 1, #lines - 1 do
  local joltDifference = lines[i+1] - lines[i]
  differences[joltDifference] = differences[joltDifference] + 1
end

print(differences[1])
print(differences[3])
print(differences[3] * differences[1])
