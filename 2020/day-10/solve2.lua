FileHandling = require 'fileHandling'

local lines = FileHandling.getNumbersFromFile('input/inputA.txt')
table.insert(lines, #lines + 1, 0)
table.sort(lines)
table.insert(lines, #lines + 1, lines[#lines] + 3)

local cache = {}

local function traverse(lines, i)
  local localAcc = i + 1;
  local resAcc = 0
  
  if localAcc == #lines then
    return 1
  end

  while localAcc <= #lines and lines[localAcc] - lines[i]  <= 3 do

    if cache[localAcc] == nil then
      local score = traverse(lines, localAcc)
      resAcc = resAcc + score
      cache[localAcc] = score
    else
      resAcc = resAcc + cache[localAcc]
    end

    localAcc = localAcc + 1
  end
  return resAcc
end

print(traverse(lines,  1))