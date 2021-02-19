function expandTab(s, tab)
  tab = tab or 8 -- tab "size" (default is 8)
  local corr = 0
  s = string.gsub(s, "()\t", function(p)
    local sp = tab - (p - 1 + corr) % tab
    corr = corr - 1 + sp
    return string.rep(" ", sp)
  end)
  return s
end

-- print(expandTab("a\tb\tc", 3))

function unexpandTab(s, tab)
  tab = tab or 8
  s = expandTab(s, tab)
  local pat = string.rep(".", tab)
  s = string.gsub(s, pat, "%0\1")
  s = string.gsub(s, " +\1", "\t")
  s = string.gsub(s, "\1", "")
  return s
end

print(unexpandTab("a  b  c", 3))
