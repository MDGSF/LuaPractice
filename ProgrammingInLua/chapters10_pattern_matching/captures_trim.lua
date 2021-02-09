function trim(s)
  s = string.gsub(s, "^%s*(.-)%s*$", "%1")
  return s
end

print(trim(" hello world   "))
