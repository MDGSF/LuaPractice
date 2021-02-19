x = string.gsub("hello world", "(%w+)", "%1 %1")
print(x) -- hello hello world world

x = string.gsub("hello world", "%w+", "%0 %0", 1)
print(x) -- hello hello world

x = string.gsub("hello world", "%w+", "%0 %0")
print(x) -- hello hello world world

x = string.gsub("hello world from lua", "(%w+)%s*(%w+)", "%2 %1")
print(x) -- world hello lua from

x = string.gsub("home = $HOME, user = $USER", "%$(%w+)", os.getenv)
print(x) -- home = /home/huangjian, user = huangjian

x = string.gsub("4+5 = $return 4+5$", "%$(.-)%$", function(s)
  return load(s)()
end)
print(x) -- 4+5 = 9

local t = { name = "lua", version = "5.3" }
x = string.gsub("$name-$version.tar.gz", "%$(%w+)", t)
print(x) -- lua-5.3.tar.gz
