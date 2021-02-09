function process()
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

process()
