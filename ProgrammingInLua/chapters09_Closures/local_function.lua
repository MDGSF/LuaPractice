-- local function foo (params) body end
-- expands to
-- local foo; foo = function (params) body end

local function fact(n)
  if n == 0 then return 1
  else return n * fact(n - 1)
  end
end

print(fact(3))
