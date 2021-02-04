N = 8

-- check whether position (new_row, new_col) is free from attack
function isplaceok (a, new_row, new_col)
  for row = 1, new_row - 1 do
    if (a[row] == new_col) or
      (a[row] - row == new_col - new_row) or
      (a[row] + row == new_col + new_row) then
      return false
    end
  end
  return true
end

-- print a board
function printsolution (a)
  for row = 1, N do
    for col = 1, N do
      io.write(a[row] == col and "X" or "-", " ")
    end
    io.write("\n")
  end
  io.write("\n")
end

-- add to board 'a' all queens from 'n' to 'N'
function addqueen (a, n)
  if n > N then
    printsolution(a)
  else
    for c = 1, N do
      if isplaceok(a, n, c) then
        a[n] = c
        addqueen(a, n + 1)
      end
    end
  end
end

addqueen({}, 1)
