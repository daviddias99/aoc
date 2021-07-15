FileHandling = require 'fileHandling'

local function convertTo2DTable(lines)
  local result = {}
  
    for _, line in pairs(lines) do
      local row = {}
      line:gsub(".",function(c) table.insert(row,c) end)
      table.insert(result,row)
    end
  return result
end

local function calculateAdjacentToPos(grid, row, column)
  local res = 0

  if row > 1 then

    if column > 1 then
      res = res + (grid[row-1][column-1] == '#' and 1 or 0)
    end

    res = res + (grid[row-1][column] == '#' and 1 or 0)

    if column < #grid[1] then
      res = res + (grid[row-1][column+1] == '#'  and 1 or 0)
    end
  end

  if column > 1 then
    res = res + (grid[row][column-1] == '#'  and 1 or 0)
  end

  if column < #grid[1] then
    res = res + (grid[row][column+1] == '#'  and 1 or 0)
  end
  
  if row < #grid then
    if column > 1 then
      res = res + (grid[row+1][column-1] == '#'  and 1 or 0)
    end
    res = res + (grid[row+1][column] == '#'  and 1 or 0)

    if column < #grid[1] then
      res = res + (grid[row+1][column+1] == '#'  and 1 or 0)
    end
  end

  return res
end

local function createNextRound(grid)
  local nextGrid = {}
  local changed = false
  local numOccupied = 0
  for i = 1, #grid do
    local row = grid[i]
    nextGrid[i] = {}
    for j = 1, #row do
      if grid[i][j] == 'L' and calculateAdjacentToPos(grid, i, j) == 0 then
        nextGrid[i][j] = '#'
        changed = true
      elseif grid[i][j] == '#' and calculateAdjacentToPos(grid, i, j) >= 4 then
        nextGrid[i][j] = 'L'
        changed = true
      else
        nextGrid[i][j] = grid[i][j]
      end 

      if nextGrid[i][j] == '#' then
        numOccupied = numOccupied + 1
      end
    end
  end

  return nextGrid, changed, numOccupied
end

local lines = FileHandling.getLinesFromFile('input/inputA.txt')
local grid = convertTo2DTable(lines)
local nextGrid, changed, numOccupied = createNextRound(grid)
while changed do
  nextGrid, changed, numOccupied = createNextRound(nextGrid)
end

print(numOccupied)
