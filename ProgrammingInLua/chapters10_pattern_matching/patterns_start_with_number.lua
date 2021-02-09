function is_start_with_number(str)
  return string.find(str, "^%d") and true or false
end

print(is_start_with_number("6asdf")) -- true
print(is_start_with_number("asdf")) -- false
