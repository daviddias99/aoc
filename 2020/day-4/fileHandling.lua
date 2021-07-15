FileHandling = {}

-- see if the file exists
function file_exists(file)
  local f = io.open(file, "rb")
  if f then f:close() end
  return f ~= nil
end

-- get all lines from a file, returns an empty 
-- list/table if the file does not exist
function numbers_from_file(file)
  if not file_exists(file) then return {} end
  local lines = {}
  for line in io.lines(file) do 
    lines[#lines + 1] = tonumber(line)
  end
  return lines
end

function strings_from_file(file)
  if not file_exists(file) then return {} end
  local lines = {}
  for line in io.lines(file) do 
    lines[#lines + 1] = line
  end
  return lines
end

function grouped_strings_from_file(file)
  if not file_exists(file) then return {} end
  local lines = {}
  local iterator = 1
  for line in io.lines(file) do 

    if line == '' then
      iterator = iterator + 1
    else
      if lines[iterator] then
        lines[iterator] = lines[iterator] .. ' ' .. line
      else
        lines[iterator] = line
      end
    end
  end

  return lines
end

FileHandling.getNumbersFromFile = numbers_from_file
FileHandling.getLinesFromFile = strings_from_file
FileHandling.getGroupedLinesFromFile = grouped_strings_from_file
FileHandling.FileExists = file_exists

return FileHandling