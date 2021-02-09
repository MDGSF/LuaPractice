s = "a (enclosed (in) parentheses) line"
print((string.gsub(s, "%b()", "")))
