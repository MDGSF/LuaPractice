function return_except_first(...)
  return select(2, ...)
end

print(return_except_first(1, 2, 3, 4, 5))
