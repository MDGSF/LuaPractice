print(#arg)
for k, v in ipairs(arg) do
  print(k, v)
end

function test1()
  if #arg == 0 then
    local lines = {}
    for line in io.lines() do
      lines[#lines + 1] = line
    end
    table.sort(lines)
    for _, line in ipairs(lines) do
      io.write(line, "\n")
    end
  elseif #arg == 1 then
    local lines = {}
    local f = assert(io.open(arg[1], "r"))
    for line in f:lines() do
      lines[#lines + 1] = line
    end
    table.sort(lines)
    for _, line in ipairs(lines) do
      io.write(line, "\n")
    end
  elseif #arg == 2 then
    local lines = {}
    local fin = assert(io.open(arg[1], "r"))
    local fout = assert(io.open(arg[2], "w"))
    for line in fin:lines() do
      lines[#lines + 1] = line
    end
    table.sort(lines)
    for _, line in ipairs(lines) do
      fout:write(line, "\n")
    end
  else
    os.exit(string.format("invalid argument number %d", #arg))
  end
end

function test2()
  if #arg == 0 then
    sort_stream(io.stdin, io.stdout)
  elseif #arg == 1 then
    local fin = assert(io.open(arg[1], "r"))
    sort_stream(fin, io.stdout)
  elseif #arg == 2 then
    local fin = assert(io.open(arg[1], "r"))
    local fout = assert(io.open(arg[2], "w"))
    sort_stream(fin, fout)
  else
    os.exit(string.format("invalid argument number %d", #arg))
  end
end

function sort_stream(fin, fout)
  local lines = {}
  for line in fin:lines() do
    lines[#lines + 1] = line
  end
  table.sort(lines)
  for _, line in ipairs(lines) do
    fout:write(line, "\n")
  end
end

test2()
