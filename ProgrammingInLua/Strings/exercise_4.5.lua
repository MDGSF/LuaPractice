function remove(a, pos, len)
  return string.sub(a, 1, pos - 1) .. string.sub(a, pos + len, #a)
end

print(remove("hello world", 7, 4))
