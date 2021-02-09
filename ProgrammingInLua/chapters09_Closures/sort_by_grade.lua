names = {"Peter", "Paul", "Mary"}
grades = {Mary = 10, Paul = 7, Peter = 8}

function sort_by_grade(names, grades)
  table.sort(names, function (n1, n2)
    return grades[n1] > grades[n2]
  end)
end

-- 在上面的匿名函数中，grades 既不是全局变量，也不是局部变量，是 non-local
-- 变量，在 Lua 里面也叫 upvalues
