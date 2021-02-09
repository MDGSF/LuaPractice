-- fourth parameter(optional): limits the number of substitution
-- string.gsub(a subject string, a pattern, a replacement string, [fourth])

s = string.gsub("Lua is cute", "cute", "great")
print(s) -- Lua is great

s = string.gsub("all lii", "l", "x")
print(s) -- axx xii

s = string.gsub("Lua is great", "Sol", "Sun")
print(s) -- Lua is great



s = string.gsub("all lii", "l", "x", 1)
print(s) -- axl lii
s = string.gsub("all lii", "l", "x", 2)
print(s) -- axx lii
