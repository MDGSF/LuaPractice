function readFile(file)
  local f = assert(io.open(file, "rb"))
  local content = f:read("*all")
  f:close()
  return content
end

local content = readFile("read_whole_file.lua")
print(content)
