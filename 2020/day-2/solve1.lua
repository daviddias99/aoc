FileHandling = require 'fileHandling'

local function parseLine(line) 
  return string.match(line, '(%d+)%-(%d+) (%a): (%a+)')
end

local function isValid(password, char, lowerBound, upperBound)
  local count = select(2, string.gsub(password,char,""))
  return count <= upperBound and count >= lowerBound;
end

local lines = FileHandling.getLinesFromFile("input/inputA.txt")
local result = 0
for _, line in pairs(lines) do
  local lowerBound, upperBound, char, password = parseLine(line)
  result = result + (isValid(password, char, tonumber(lowerBound), tonumber(upperBound)) and 1 or 0)
end

print(result)