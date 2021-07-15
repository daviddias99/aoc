FileHandling = require 'fileHandling'

local lines = FileHandling.getLinesFromFile('input/inputA.txt')
local buses = {}
local indexes = {}
local iterator = string.gmatch(lines[2], '(%w+)')
local N = 1
local index = 0
for value in iterator do
  index = index + 1
  local bus
  if value ~= 'x' then
    bus = tonumber(value)
    N = N * bus
    table.insert(buses, #buses + 1, bus)
    table.insert(indexes, #indexes + 1, (bus - (index - 1)) % bus)
  end
end

local function extendedEuclid(x, y)
  local x0, x1, y0, y1, q = 1, 0, 0 ,1, nil

  while y > 0 do
    q, x ,y = math.floor(x / y), y, x % y
    x0, x1 = x1, x0 - q * x1
    y0, y1 = y1, y0 - q * y1
  end

  return q, x0 , y0
end

local function invMod(a, m)
  local _, x, _ = extendedEuclid(a, m)

  return x % m
end

local function chineseRemainderEuclid(n, Num, a)
  local result = 0

  for i = 1, #n do
    local ai = a[i]
    local ni = n[i]
    
    local _, _, si = extendedEuclid(ni, Num // ni)
    result = result + ai * si * (Num // ni)
  end

  return (result * -1 + Num) % Num
end

local function chineseRemainderGauss(n, Num, a)
  local result = 0

  for i = 1, #n do
    local ai = a[i]
    local ni = n[i]
    local bi = Num // ni
    
    result = result + ai * bi * invMod(bi,ni)
  end

  return result % Num
end

print(chineseRemainderGauss(buses, N, indexes))
print(chineseRemainderEuclid(buses, N, indexes))