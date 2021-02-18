package state

/*
http://www.lua.org/manual/5.3/manual.html#lua_pushnil

void lua_pushnil (lua_State *L);

Pushes a nil value onto the stack.
*/
func (self *luaState) PushNil() {
	self.stack.push(nil)
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_pushboolean

void lua_pushboolean (lua_State *L, int b);

Pushes a boolean value with value b onto the stack.
*/
func (self *luaState) PushBoolean(b bool) {
	self.stack.push(b)
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_pushinteger

void lua_pushinteger (lua_State *L, lua_Integer n);

Pushes an integer with value n onto the stack.
*/
func (self *luaState) PushInteger(n int64) {
	self.stack.push(n)
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_pushnumber

void lua_pushnumber (lua_State *L, lua_Number n);

Pushes a float with value n onto the stack.
*/
func (self *luaState) PushNumber(n float64) {
	self.stack.push(n)
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_pushstring

const char *lua_pushstring (lua_State *L, const char *s);

Pushes the zero-terminated string pointed to by s onto the stack. Lua makes
(or reuses) an internal copy of the given string, so the memory at s can be
freed or reused immediately after the function returns.

Returns a pointer to the internal copy of the string.

If s is NULL, pushes nil and returns NULL.
*/
func (self *luaState) PushString(s string) {
	self.stack.push(s)
}
