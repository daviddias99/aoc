FileHandling = require 'fileHandling'

local function convertTo2DTable(lines)
  local result = {}
  
    for i, line in pairs(lines) do
      local row = {}
      line:gsub(".",function(c) table.insert(row,c) end)
      table.insert(result,row)
    end

  return result
end

local function checkForSlope(matrix, xInc, yInc)

  local currentPos = {x = 0, y = 0}
  local treeHitCount = 0
  while currentPos.y < #matrix do
    if(matrix[currentPos.y + 1][currentPos.x + 1] == '#') then
      treeHitCount = treeHitCount + 1
    end
  
    currentPos.x = (currentPos.x + xInc) % #matrix[1]
    currentPos.y = currentPos.y + yInc
  end

  return treeHitCount
end

local lines = FileHandling.getLinesFromFile("input/inputA.txt")
local matrix = convertTo2DTable(lines)

local res1 = checkForSlope(matrix, 1, 1)
local res2 = checkForSlope(matrix, 3, 1)
local res3 = checkForSlope(matrix, 5, 1)
local res4 = checkForSlope(matrix, 7, 1)
local res5 = checkForSlope(matrix, 1, 2)

print(res1)
print(res2)
print(res3)
print(res4)
print(res5)
print(res1 * res2 * res3 * res4 * res5)