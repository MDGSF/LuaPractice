function expand(s)
  return (string.gsub(s, "$(%w+)", _G))
end

name = "Lua"; status = "great"

print(expand("$name is $status, isn't it?"))

print(expand("$othername is $status, isn't is?"))
