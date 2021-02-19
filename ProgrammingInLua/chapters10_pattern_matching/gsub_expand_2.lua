function expand(s)
  return (string.gsub(s, "$(%w+)", function(n)
    return tostring(_G[n])
  end))
end

print(expand("print = $print; a = $a"))
