print(io.open("non-existent-file", "r"))

print(io.open("/etc/passwd", "w"))

-- local f = assert(io.open(filename, mode))
-- local t = f:read("a")
-- f:close()
