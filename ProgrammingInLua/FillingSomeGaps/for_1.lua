-- for var = exp1, exp2, exp3 do
--   something
-- end

for i = 1, 3 do
  print(i)
end

for i = 3, 1, -1 do
  print(i)
end

for i = 10, 1, -2 do
  print(i)
end

for i = 1, math.huge do
  print(i)
  if i > 3 then
    break
  end
end
