function add(...)
  local s = 0
  for _, v in ipairs{...} do
    s = s + v
  end
  return s
end

function add2(...)
  local s = 0
  for i = 1, select("#", ...) do
    s = s + select(i, ...)
  end
  return s
end

print(add(3, 4, 10, 25, 12))
print(add2(3, 4, 10, 25, 12))

