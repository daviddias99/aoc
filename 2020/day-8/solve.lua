FileHandling = require 'fileHandling'

local function parseInstructions(lines)
  local instructions = {}
  for _, line in pairs(lines) do
    local inst, arg = string.match(line, '(%a+) ([%+%-]%d+)')
    local instruction = {inst, arg, false}
    table.insert(instructions, instruction)
  end

  return instructions
end

local function execInstruction(instruction, pc, acc, visited) 
  
  if(visited[pc]) then
    return nil, acc
  end

  visited[pc] = true

  if instruction[1] == 'nop' then
    pc = pc + 1
  elseif instruction[1] == 'jmp' then
    pc = pc + tonumber(instruction[2])
  elseif instruction[1] == 'acc' then
    pc = pc + 1
    acc = acc + tonumber(instruction[2])  
  end

  return pc, acc
end

local function runProgram(instructions)

  local programCounter = 1
  local accumulator = 0
  local visited = {}
  while programCounter <= #instructions do
    local currentInstruction = instructions[programCounter]
    programCounter, accumulator = execInstruction(currentInstruction, programCounter, accumulator, visited)
  
    if programCounter == nil then
      return programCounter, accumulator
    end
  end
  return programCounter, accumulator
end

local lines = FileHandling.getLinesFromFile('input/inputA.txt')
local instructions = parseInstructions(lines)

print(runProgram(instructions))

for _, instruction in pairs(instructions) do
  if instruction[1] == 'jmp' then
    instruction[1] = 'nop'
    local pc, acc = runProgram(instructions)
    if pc ~= nil then
      print(pc, acc)
    end
    instruction[1] = 'jmp'
  elseif instruction[1] == 'nop' then
    instruction[1] = 'jmp'
    local pc, acc = runProgram(instructions)
    if acc  ~= nil then
      print(pc, acc)
    end
    instruction[1] = 'nop'
  end
end