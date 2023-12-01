file = open("input_small", "r")

numbers = map(x -> parse(Int, x), split(readline(file), ","))
boards = []
currentBoard = []

line = readline(file)
line = readline(file)
while !eof(file)

  if line == ""
    append!(boards, [currentBoard])
    global line = readline(file)
    global currentBoard = []
    continue
  end
  append!(currentBoard, [map(x -> parse(Int, x), split(strip(line), r"\s+"))])
  
  line = readline(file)
end

append!(boards, [currentBoard])

display(boards)