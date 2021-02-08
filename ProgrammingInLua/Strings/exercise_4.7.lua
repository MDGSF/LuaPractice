function ispali(str)
  i, j = 1, #str
  while i < j do
    if string.byte(str, i) ~= string.byte(str, j) then
      return false
    end
    i, j = i + 1, j - 1
  end
  return true
end

print(ispali("step on no pets"))
print(ispali("banana"))
