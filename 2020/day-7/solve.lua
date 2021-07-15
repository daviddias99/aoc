FileHandling = require 'fileHandling'

local function parseLine(bags, line)
  local outerBag = line:match('([%a%s]+) bags contain')
  local innerBags = line:gmatch('(%d) ([%a%s]+) bag') 
  local innerBagsTable = {}
  for quant, color in innerBags do
    innerBagsTable[color] = quant
  end
  bags[outerBag] = innerBagsTable
end

local function searchBagForColor(bag, color, bags)
  local result = 0

  for bagColor, _ in pairs(bag) do
      if bagColor == color then return 1 end

      result = searchBagForColor(bags[bagColor], color, bags)

      if result == 1 then return result end
  end

  return 0
end

local function countBagsIn(bag, bags)
  local result = 0

  for key, value in pairs(bag) do
    result = result + value + value * countBagsIn(bags[key],bags)
  end

  return result
end

local lines = FileHandling.getLinesFromFile('input/inputA.txt')
local bags = {}

for _, value in pairs(lines) do
    parseLine(bags,value)
end

local res = 0
for key, value in pairs(bags) do
  if key ~= 'shiny gold' then
    res = res + searchBagForColor(value, 'shiny gold', bags)
  end
end

print(res)

res = countBagsIn(bags['shiny gold'], bags)

print(res)

