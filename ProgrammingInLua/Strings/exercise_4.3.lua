function insert (a, i, b)
  return string.sub(a, 0, i - 1) .. b .. string.sub(a, i, #a)
end

print(insert("hello world", 1, "start: "))
print(insert("hello world", 7, "small "))
