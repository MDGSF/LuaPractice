a = { p = print }
a.p("Hello World")
print = math.sin
a.p(print(1))
math.sin = a.p
math.sin(10, 20)
