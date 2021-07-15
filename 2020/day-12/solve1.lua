FileHandling = require 'fileHandling'

local function extractCommands(lines)
  local commands = {}
  for _, line in pairs(lines) do
    local dir, dist = string.match(line, '(%a)(%d+)')
    table.insert(commands, #commands+1, {dir= dir,dist= dist})
  end
  
  return commands
end

local function doRotation(pos, side, degrees)
  local rotating = degrees/90
  local directions = {'N', 'E', 'S', 'W'}
  local dirModifier = side == 'L' and -1 or 1
  local res = {x=pos.x, y=pos.y, front=pos.front}

  for key, value in ipairs(directions) do
    if value == pos.front then
      res.front = directions[((key - 1) + rotating * dirModifier) % 4 + 1]
      return res
    end
  end

end

local function execMove(initPos, command)

  local res = {x=initPos.x, y=initPos.y, front=initPos.front}
  if command.dir == 'N' then
    res.y = initPos.y + command.dist
  elseif command.dir == 'S'then
    res.y = initPos.y - command.dist
  elseif command.dir == 'E'then
    res.x = initPos.x + command.dist
  elseif command.dir == 'W'then
    res.x = initPos.x - command.dist
  end
  
  return res
end

local lines = FileHandling.getLinesFromFile('input/inputA.txt')
local commands = extractCommands(lines)
local pos = {x=0,y=0,front='E'}

for _, command in pairs(commands) do
  pos = execMove(pos, command)
  if command.dir == 'L' or command.dir == 'R' then
    pos = doRotation(pos, command.dir, command.dist)
  elseif command.dir == 'F'then
    pos = execMove(pos, {dir=pos.front, dist=command.dist})
  end
end

print(pos.x)
print(pos.y)
print(math.abs(pos.x) + math.abs(pos.y))