function unpack (t, i, n)
  i = i or 1
  n = n or #t
  if i <= n then
    return t[i], unpack(t, i + 1, n)
  end
end

print(unpack({"Sunday", "Monday", "Tuesday", "Wednesday"}, 2, 3))

