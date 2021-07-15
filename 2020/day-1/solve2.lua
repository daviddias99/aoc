FileHandling = require 'fileHandling'

local file = 'input/input_large.txt'
local input = FileHandling.getNumbersFromFile(file)
table.sort(input)

for i = 1, #input - 2 do
  for j = i + 1, #input - 1 do

    if 2020 - input[i] < input[j] then
      print('Close, but no cigar')
      break
    end

    for k = j + 1, #input do

      if 2020 - input[i] - input[j] < input[k] then
        print('Close, but no cigar')
        break
      end

      if 2020 - input[i] - input[j] == input[k] then
        print(string.format('%d + %d + %d = 2020, %d * %d * %d = %d', input[i], input[j], input[k], input[i], input[j], input[k], input[i] * input[j] * input[k]))
        return
      end
    end
  end    
end