function nonils(...)
  local arg = table.pack(...)
  for i = 1, arg.n do
    if arg[i] == nil then
      return false
    end
  end
  return true
end

print(nonils(2, 3, nil))
print(nonils(2, 3))
print(nonils())
print(nonils(nil))

