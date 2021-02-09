-- string.find("a [word]", "[")
-- lua: find_2.lua:1: malformed pattern (missing ']')

-- third parameter: start index to search
-- fourth parameter: use plain search, ignores patterns.
print(string.find("a [word]", "[", 1, true))
