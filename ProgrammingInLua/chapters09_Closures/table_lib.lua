Lib = {}
Lib.foo = function (x, y) return x + y end
Lib.goo = function (x, y) return x - y end

Lib1 = {
  foo = function (x, y) return x + y end,
  goo = function (x, y) return x - y end
}

Lib2 = {}
function Lib2.foo (x, y) return x + y end
function Lib2.goo (x, y) return x - y end

print(Lib.foo(2, 3), Lib.goo(2, 3))
print(Lib1.foo(2, 3), Lib1.goo(2, 3))
print(Lib2.foo(2, 3), Lib2.goo(2, 3))
