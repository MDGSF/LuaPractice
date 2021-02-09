function return_except_last(...)
  arg = {...}
  table.remove(arg)
  return table.unpack(arg)
end

print(return_except_last(1, 2, 3, 4, 5))
