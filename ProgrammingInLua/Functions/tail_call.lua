function foo(n)
  if n > 0 then
    return foo(n - 1)
  end
end

foo(10000)
