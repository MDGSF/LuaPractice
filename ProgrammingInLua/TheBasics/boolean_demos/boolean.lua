#!/usr/bin/env lua

-- lua consider both the Boolean false and nil as false
-- and anything else as true.


-- and
-- first operand is false, result is first operand
-- otherwise, result is second operand
print(4 and 5) -- 5
print(nil and 13) -- nil
print(false and 13) -- false


-- or
-- first operand is true, result is first operand
-- otherwise, result is second operand
print(0 or 5) -- 0
print(false or "hi") -- hi
print(nil or false) -- false


-- x = x or v, which is equivalent to
-- if not x then x = v end
x = 10
x = x or 11
print(x) -- 10

x = false
x = x or 11
print(x) -- 11

-- ((a and b) or c)
-- (a and b or c)
-- a ? b : c
a = 100
b = 99
result = (a > b) and a or b
print(result) -- 100


-- not
print(not nil) -- true
print(not false) -- true
print(not 0) -- false
print(not "") -- false
print(not not 1) -- true
print(not not nil) -- false
