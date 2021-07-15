FileHandling = require 'fileHandling'

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

    if maskValue == 'X' then
      resultValue = valueValue
    else
      resultValue = maskValue
    end
    
    result = resultValue .. result
    i = i - 1
  end
  
  return result
end


local lines = FileHandling.getLinesFromFile('input/inputA.txt')
local mask = string.match(lines[1], 'mask = (%w+)')
local memory = {}

for _, value in ipairs(lines) do
  if string.match(value, 'mask') then
    mask = string.match(value, 'mask = (%w+)')
  else
    local addr, val = string.match(value, "mem%[(%d+)] = (%d+)")
    local applied = applyMask(toBits(math.floor(val),36), mask)
    memory[addr] = binToDec(applied)
  end
end

local sum = 0
for _, value in pairs(memory) do
  sum = sum + value
end

print(sum)