-- floor, round towards minus infinite
print(math.floor(3.3)) -- 3
print(math.floor(-3.3)) -- -4

-- ceil, round towards plus infinite
print(math.ceil(3.3)) -- 4
print(math.ceil(-3.3)) -- -3

-- modf, round towards zero
print(math.modf(3.3))  --  3    0.3
print(math.modf(-3.3)) -- -3   -0.3

-- round, rounds towards nearest integer
-- always round half-integers up
function round (x)
  local f = math.floor(x)
  if x == f then return f
  else return math.floor(x + 0.5)
  end
end

print(round(1))   -- 1
print(round(1.1)) -- 1
print(round(1.5)) -- 2
print(round(1.9)) -- 2

-- round, rounds towards nearest integer
-- but round half-integers to the nearest even integer
function round_even (x)
  local f = math.floor(x)
  if (x == f) or (x % 2.0 == 0.5) then
    return f
  else
    return math.floor(x + 0.5)
  end
end

print(round_even(2.5))  -- 2
print(round_even(3.5))  -- 4
print(round_even(-2.5)) -- -2
print(round_even(-3.5)) -- -4
