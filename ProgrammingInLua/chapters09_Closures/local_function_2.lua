local g

local function f(n)
  if n == 0 then return end
  print("f = ", n)
  g(n-1)
end

function g(n)
  if n == 0 then return end
  print("g = ", n)
  f(n-1)
end

f(10)
