while true do
  local block = io.read(2^13)
  if not block then break end
  io.write(block)
end
