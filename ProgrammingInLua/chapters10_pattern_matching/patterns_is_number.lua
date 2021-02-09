function is_number(str)
  return string.find(str, "^[+-]?%d+$") and true or false
end

print(is_number("123")) -- true
print(is_number("+123")) -- true
print(is_number("-123")) -- true
print(is_number("asdf")) -- false
