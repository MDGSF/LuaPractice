-- print the first non-empty input line
local line
repeat
  line = io.read()
until line ~= ""
print(line)
