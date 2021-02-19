function unexcape(s)
  s = string.gsub(s, "+", " ")
  s = string.gsub(s, "%%(%x%x)", function(h)
    return string.char(tonumber(h, 16))
  end)
  return s
end

-- print(unexcape("a%2Bb+%3D+c")) -- a+b = c

cgi = {}
function decode(s)
  for name, value in string.gmatch(s, "([^&=]+)=([^&=]+)") do
    name = unexcape(name)
    value = unexcape(value)
    cgi[name] = value
  end
end

function escape(s)
  s = string.gsub(s, "[&=+%%%c]", function(c)
    return string.format("%%%02X", string.byte(c))
  end)
  s = string.gsub(s, " ", "+")
  return s
end

function encode(t)
  local b = {}
  for k, v in pairs(t) do
    b[#b + 1] = (escape(k) .. "=" .. escape(v))
  end
  return table.concat(b, "&")
end

t = {name = "al", query = "a+b = c", q = "yes or no"}
print(encode(t))
