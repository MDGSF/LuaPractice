local f = io.popen("ls", "r")
local dir = {}
for entry in f:lines() do
  print(entry)
  dir[#dir + 1] = entry
end

