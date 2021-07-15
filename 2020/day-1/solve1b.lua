FileHandling = require 'fileHandling'

local file = 'input/input_large.txt'
local input = FileHandling.getNumbersFromFile(file)
table.sort(input)

for i = 1, #input - 1 do
  for j = i + 1, #input do
    if 2020 - input[i] < input[j] then
      print('Close, but no cigar')
      break
    end

    if 2020 - input[i] == input[j] then
      print(string.format('%d + %d = 2020, %d * %d = %d', input[i], input[j], input[i], input[j], input[i] * input[j]))
      return
    end
  end    
end