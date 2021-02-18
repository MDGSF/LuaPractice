package state

/*
http://www.lua.org/manual/5.3/manual.html#lua_len

void lua_len (lua_State *L, int index);

Returns the length of the value at the given index. It is equivalent to the '#'
operator in Lua (see §3.4.7) and may trigger a metamethod for the "length" event
(see §2.4). The result is pushed on the stack.
*/
func (self *luaState) Len(idx int) {
	val := self.stack.get(idx)

	if s, ok := val.(string); ok {
		self.stack.push(int64(len(s)))
	} else {
		panic("length error!")
	}
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_concat

void lua_concat (lua_State *L, int n);

Concatenates the n values at the top of the stack, pops them, and leaves the
result at the top. If n is 1, the result is the single value on the stack (that
is, the function does nothing); if n is 0, the result is the empty string.
Concatenation is performed following the usual semantics of Lua (see §3.4.6).
*/
func (self *luaState) Concat(n int) {
	if n == 0 {
		self.stack.push("")
	} else if n >= 2 {
		for i := 1; i < n; i++ {
			if self.IsString(-1) && self.IsString(-2) {
				s2 := self.ToString(-1)
				s1 := self.ToString(-2)
				self.stack.pop()
				self.stack.pop()
				self.stack.push(s1 + s2)
				continue
			}

			panic("concatenation error!")
		}
	}
	// n == 1, do nothing
}
