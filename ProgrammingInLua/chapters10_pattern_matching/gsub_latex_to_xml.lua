function toxml(s)
  s = string.gsub(s, "\\(%a+)(%b{})", function(tag, body)
    body = string.sub(body, 2, -2) -- remove the brackets
    body = toxml(body) -- handle nested command
    return string.format("<%s>%s</%s>", tag, body, tag)
  end)
  return s
end

print(toxml("\\title{The \\bold{big} example}"))

