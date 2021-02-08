function table_concat(t)
  result = ""
  for k, v in ipairs(t) do
    result = result .. v
  end
  return result
end

print(table_concat({"hello", " ", "world"}))
