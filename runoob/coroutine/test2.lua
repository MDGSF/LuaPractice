function foo (a)
  print("foo a = ", a)
  return coroutine.yield(2 * a)
end

co = coroutine.create(function (a, b)
  print("first ", a, b)
  local r = foo(a + 1)

  print("second ", r)
  local r, s = coroutine.yield(a + b, a - b)

  print("third ", r, s)
  return b, "finished"
end)

print("main", coroutine.resume(co, 1, 10))
print("------------")
print("main", coroutine.resume(co, "r"))
print("------------")
print("main", coroutine.resume(co, "x", "y"))
print("------------")
print("main", coroutine.resume(co, "x", "y"))
