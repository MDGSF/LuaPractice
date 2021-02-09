function new_counter()
  local count = 0
  return function()
    count = count + 1
    return count
  end
end

c1 = new_counter()
print(c1())
print(c1())
print(c1())
