package state

/*
http://www.lua.org/manual/5.3/manual.html#lua_gettop

int lua_gettop (lua_State *L);

Returns the index of the top element in the stack. Because indices start at 1,
this result is equal to the number of elements in the stack; in particular, 0
means an empty stack.
*/
func (self *luaState) GetTop() int {
	return self.stack.top
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_absindex

int lua_absindex (lua_State *L, int idx);

Converts the acceptable index idx into an equivalent absolute index (that is,
one that does not depend on the stack top).
*/
func (self *luaState) AbsIndex(idx int) int {
	return self.stack.absIndex(idx)
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_checkstack

int lua_checkstack (lua_State *L, int n);

Ensures that the stack has space for at least n extra slots (that is, that you
can safely push up to n values into it). It returns false if it cannot fulfill
the request, either because it would cause the stack to be larger than a fixed
maximum size (typically at least several thousand elements) or because it cannot
allocate memory for the extra space. This function never shrinks the stack; if
the stack already has space for the extra slots, it is left unchanged.
*/
func (self *luaState) CheckStack(n int) bool {
	self.stack.check(n)
	return true
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_pop

void lua_pop (lua_State *L, int n);

Pops n elements from the stack.
*/
func (self *luaState) Pop(n int) {
	for i := 0; i < n; i++ {
		self.stack.pop()
	}
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_copy

void lua_copy (lua_State *L, int fromidx, int toidx);

Copies the element at index fromidx into the valid index toidx, replacing the
value at that position. Values at other positions are not affected.
*/
func (self *luaState) Copy(fromIdx, toIdx int) {
	val := self.stack.get(fromIdx)
	self.stack.set(toIdx, val)
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_pushvalue

void lua_pushvalue (lua_State *L, int index);

Pushes a copy of the element at the given index onto the stack.
*/
func (self *luaState) PushValue(idx int) {
	val := self.stack.get(idx)
	self.stack.push(val)
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_replace

void lua_replace (lua_State *L, int index);

Moves the top element into the given valid index without shifting any element
(therefore replacing the value at that given index), and then pops the top
element.
*/
func (self *luaState) Replace(idx int) {
	val := self.stack.pop()
	self.stack.set(idx, val)
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_insert

void lua_insert (lua_State *L, int index);

Moves the top element into the given valid index, shifting up the elements above
this index to open space. This function cannot be called with a pseudo-index,
because a pseudo-index is not an actual stack position.
*/
func (self *luaState) Insert(idx int) {
	self.Rotate(idx, 1) // 把 top 位置的元素旋转到 idx 位置
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_remove

void lua_remove (lua_State *L, int index);

Removes the element at the given valid index, shifting down the elements above
this index to fill the gap. This function cannot be called with a pseudo-index,
because a pseudo-index is not an actual stack position.
*/
func (self *luaState) Remove(idx int) {
	self.Rotate(idx, -1) // 把 idx 位置的元素旋转到 top 位置
	self.Pop(1)          // 弹出 top 位置的元素
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_rotate

void lua_rotate (lua_State *L, int idx, int n);

Rotates the stack elements between the valid index idx and the top of the stack.
The elements are rotated n positions in the direction of the top, for a positive
n, or -n positions in the direction of the bottom, for a negative n. The
absolute value of n must not be greater than the size of the slice being
rotated. This function cannot be called with a pseudo-index, because a
pseudo-index is not an actual stack position.
*/
func (self *luaState) Rotate(idx, n int) {
	t := self.stack.top - 1           // end of stack segment being rotated
	p := self.stack.absIndex(idx) - 1 // start of segment
	var m int                         // end of prefix
	if n >= 0 {
		m = t - n
	} else {
		m = p - n - 1
	}
	self.stack.reverse(p, m)   // reverse the prefix with length 'n'
	self.stack.reverse(m+1, t) // reverse the suffix
	self.stack.reverse(p, t)   // reverse the entire segment
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_settop

void lua_settop (lua_State *L, int index);

Accepts any index, or 0, and sets the stack top to this index. If the new top is
larger than the old one, then the new elements are filled with nil. If index is
0, then all stack elements are removed.
*/
func (self *luaState) SetTop(idx int) {
	newTop := self.stack.absIndex(idx)
	if newTop < 0 {
		panic("stack underflow!")
	}

	n := self.stack.top - newTop
	if n > 0 {
		for i := 0; i < n; i++ {
			self.stack.pop()
		}
	} else if n < 0 {
		for i := 0; i > n; i-- {
			self.stack.push(nil)
		}
	}
}
