FileHandling = require 'fileHandling'

local function create1DTableWithSize(xSize)
  local result = {}

  for i = 1, xSize do
    table.insert(result, #result + 1, 0)
  end

  return result
end

local function create2DTableWithSize(ySize, xSize)
  local result = {}

  for i = 1, ySize do
    table.insert(result, #result + 1, create1DTableWithSize(xSize))
  end

  return result
end

local function create3DTableWithSize(zSize, ySize, xSize)
  local result = {}

  for i = 1, zSize do
    table.insert(result, #result + 1, create2DTableWithSize(ySize, xSize))
  end

  return result
end

local function create4DTableWithSize(wSize, zSize, ySize, xSize)
  local result = {}

  for i = 1, wSize do
    table.insert(result, #result + 1, create4DTableWithSize(zSize, ySize, xSize))
  end

  return result
end

local function expandSpace(space)

  for y, row in ipairs(space[#space]) do
    for x, value in ipairs(row) do
      if value == 1 then
        table.insert(space, #space + 1, create2DTableWithSize(#space[#space], #row))
        print('expanded higher z')
        goto endCycle1
      end
    end
  end
  ::endCycle1::

  for y, row in ipairs(space[1]) do
    for x, value in ipairs(row) do
      if value == 1 then
        table.insert(space, 1, create2DTableWithSize(#space[1], #row))
        print('expanded lower z')
        goto endCycle2
      end
    end
  end
  ::endCycle2::

  for z, slice in ipairs(space) do
    for key, value in ipairs(slice[1]) do
      if value == 1 then
        for zInner, sliceInner in pairs(space) do
          table.insert(sliceInner, 1, create1DTableWithSize(#sliceInner[1]))
        end
        print('expanded lower y')
        goto endCycle3
      end
    end
  end
  ::endCycle3::
  print(#space)
  for z, slice in ipairs(space) do
    for key, value in ipairs(slice[#slice]) do
      if value == 1 then
        for zInner, sliceInner in pairs(space) do
          table.insert(sliceInner, #sliceInner + 1, create1DTableWithSize(#sliceInner[1]))
        end
        print('expanded higher y')
        goto endCycle4
      end
    end
  end

  ::endCycle4::
  for z, slice in ipairs(space) do
    for key, value in ipairs(slice) do
      if value[1] == 1 then
        for zInner, sliceInner in pairs(space) do
          for yInner, rowInner in ipairs(sliceInner) do
            table.insert(rowInner, 1, 0)
          end
        end
        print('expanded lower y')

        goto endCycle5
      end
    end
  end
  ::endCycle5::
  
  for z, slice in ipairs(space) do
    for key, value in ipairs(slice) do
      if value[#value] == 1 then
        for zInner, sliceInner in pairs(space) do
          for yInner, rowInner in ipairs(sliceInner) do
            table.insert(rowInner, #rowInner + 1, 0)
          end
        end
        goto endCycle6
      end
    end
  end
  ::endCycle6::

end

local function countNear(space, x, y, z)
  local result = 0
  local xLower = x == 1 and 0 or -1
  local xUpper = x == #space[1][1] and 0 or 1
  local yLower = y == 1 and 0 or -1
  local yUpper = y == #space[1] and 0 or 1
  local zLower = z == 1 and 0 or -1
  local zUpper = z == #space and 0 or 1

  for i = xLower, xUpper do
    for j = yLower, yUpper do
      for k = zLower, zUpper do
        if not(i == 0 and j == 0 and k == 0) then
          result = result + space[z + k][y + j][x + i]
        end
      end
    end
  end

  return result
end


local lines = FileHandling.getLinesFromFile('input/inputA.txt')
local space = {{{}}}

for y, line in ipairs(lines) do
  local x = 1
  table.insert(space[1][1], y, {})
  for element in string.gmatch(line,'.') do
    table.insert(space[1][1][y], x, element == '#' and 1 or 0)
    x = x + 1
  end
end

expandSpace(space)

local nextSpace = create4DTableWithSize(#space, #space[1][1], #space[1][1][1])

-- for z, plane in ipairs(space) do
--   print('>', z)
--   for y, row in ipairs(plane) do
--     for x, value in ipairs(row) do
--       io.write(value)
--       io.write(' ')
--     end
--     print()
--   end
-- end

-- for z = 1, #space do
--   for y = 1, #space[1] do
--     for x = 1, #space[1][1] do
--       local near = countNear(space, x, y ,z)
--       nextSpace[z][y][x] = space[z][y][x]

--       if space[z][y][x] == 1 and not (near == 2 or near == 3) then
--         nextSpace[z][y][x] = 0
--       elseif space[z][y][x] == 0 and near == 3 then
--         nextSpace[z][y][x] = 1
--       end
--     end
--   end
-- end

for i = 1, 6 do
  local nextSpace = create3DTableWithSize(#space, #space[1], #space[1][1])

  for z = 1, #space do
    for y = 1, #space[1] do
      for x = 1, #space[1][1] do
        local near = countNear(space, x, y ,z)
        nextSpace[z][y][x] = space[z][y][x]
  
        if space[z][y][x] == 1 and not (near == 2 or near == 3) then
          nextSpace[z][y][x] = 0
        elseif space[z][y][x] == 0 and near == 3 then
          nextSpace[z][y][x] = 1
        end
      end
    end
  end
  for z, plane in ipairs(nextSpace) do
    print('>', z)
    for y, row in ipairs(plane) do
      for x, value in ipairs(row) do
        io.write(value)
        io.write(' ')
      end
      print()
    end
  end
  space = nextSpace
  expandSpace(space)
end

local sum = 0

for z, plane in ipairs(space) do
  print('>', z)
  for y, row in ipairs(plane) do
    for x, value in ipairs(row) do
      io.write(value)
      io.write(' ')

      sum = sum + value
    end
    print()
  end
end

print(sum)


