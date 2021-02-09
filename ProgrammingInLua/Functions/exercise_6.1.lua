function print_array(a)
  for _, v in ipairs(a) do
    print(v)
  end
end

print_array({1, 2, 3, 4, 5})
