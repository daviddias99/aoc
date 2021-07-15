FileHandling = require 'fileHandling'

local function tableConcat(t1,t2)
  for i=1,#t2 do
      table.insert(t1, #t1 + 1, t2[i])
  end
  return t1
end

local function toBits(num,bits)
  -- returns a table of bits, most significant first.
  bits = bits or math.max(1, select(2, math.frexp(num)))
  local t = {} -- will contain the bits        
  for b = bits, 1, -1 do
      t[b] = math.fmod(num, 2)
      num = math.floor((num - t[b]) / 2)
  end
  return table.concat(t)
end

local function binToDec(bin)
  bin = string.reverse(bin)
  local sum = 0

  for i = 1, string.len(bin) do
    local num = string.sub(bin, i,i) == "1" and 1 or 0
    sum = sum + num * math.pow(2, i-1)
  end

  return math.floor(sum)
end

local function applyMask(value, mask)
  local result = ''
  local i = string.len(value)

  while i > 0 do
    local maskValue = string.char(string.byte(mask, i, i))
    local valueValue = string.char(string.byte(value, i, i))
    local resultValue = ''

    if maskValue == '1' then
      resultValue = '1'
    elseif maskValue == 'X' then
      resultValue = 'X'
    elseif maskValue == '0' then
      resultValue = valueValue
    end
    
    result = resultValue .. result
    i = i - 1
  end
  
  return result
end

local function findXAndReplace(value) 
  local result = {}
  for i = 1, string.len(value) do
    local valueValue = string.char(string.byte(value, i, i))

    if valueValue == 'X' then
      local temp0 = string.sub(value,0,i-1) .. '0' .. string.sub(value,i+1)
      local temp1 = string.sub(value,0,i-1) .. '1' .. string.sub(value,i+1)
      result = tableConcat(result, findXAndReplace(temp0))
      result = tableConcat(result, findXAndReplace(temp1))
      return result
    end
  end
  return {value}
end

string.gsub()
local lines = FileHandling.getLinesFromFile('input/inputA.txt')
local mask = string.match(lines[1], 'mask = (%w+)')
local memory = {}

for _, value in ipairs(lines) do
  if string.match(value, 'mask') then
    mask = string.match(value, 'mask = (%w+)')
  else
    local addr, val = string.match(value, "mem%[(%d+)] = (%d+)")
    local applied = applyMask(toBits(math.floor(addr),36), mask)
    local possibleAddresses = findXAndReplace(applied)
    for _, value in pairs(possibleAddresses) do
      memory[binToDec(value)] = val
    end
  end
end

local sum = 0
for _, value in pairs(memory) do
  sum = sum + value
end

print(sum)