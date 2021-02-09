-- computes the square root of 'x' using Newton-Raphson method
x = 9
local sqr = x / 2
repeat
  sqr = (sqr + x / sqr) / 2
  local error = math.abs(sqr^2 - x)
until error < x / 10000

print(sqr)
