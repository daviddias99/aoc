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
  local rotating = degrees
  local directions = {'N', 'E', 'S', 'W'}
  local dirModifier = side == 'L' and 1 or -1
  rotating = rotating * dirModifier
  local res = {x=pos.x * math.cos(math.pi * rotating / 180) - pos.y * math.sin(math.pi * rotating / 180), 
      y=pos.y * math.cos(math.pi * rotating / 180) + pos.x * math.sin(math.pi * rotating / 180), front=pos.front}

  return res;
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
local shipPos = {x=0,y=0}
local waypointPos = {x=10,y=1}

for _, command in pairs(commands) do
  waypointPos = execMove(waypointPos, command)
  if command.dir == 'L' or command.dir == 'R' then
    waypointPos = doRotation(waypointPos, command.dir, command.dist)
  elseif command.dir == 'F'then
    shipPos = {x=shipPos.x + waypointPos.x * command.dist,y= shipPos.y + waypointPos.y * command.dist}
  end

  print(string.format('Command: %s %d', command.dir, command.dist))
  print(string.format('ShipPos: %f %f', shipPos.x, shipPos.y))
  print(string.format('Waypoint: %f %f', waypointPos.x, waypointPos.y))
end

print(shipPos.x)
print(shipPos.y)
print(math.abs(shipPos.x) + math.abs(shipPos.y))