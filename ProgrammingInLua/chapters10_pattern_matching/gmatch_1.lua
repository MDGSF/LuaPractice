s = "some thing"
words = {}
for w in string.gmatch(s, "%a+") do -- %a+ matches word
  words[#words + 1] = w
end

for k, v in pairs(words) do
  print(k, v)
end
