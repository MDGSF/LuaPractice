package state

import (
	"fmt"
	. "luago/api"
)

/*
http://www.lua.org/manual/5.3/manual.html#lua_typename

const char *lua_typename (lua_State *L, int tp);

Returns the name of the type encoded by the value tp, which must be one the
values returned by lua_type.
*/
func (self *luaState) TypeName(tp LuaType) string {
	switch tp {
	case LUA_TNONE:
		return "no value"
	case LUA_TNIL:
		return "nil"
	case LUA_TBOOLEAN:
		return "boolean"
	case LUA_TNUMBER:
		return "number"
	case LUA_TSTRING:
		return "string"
	case LUA_TTABLE:
		return "table"
	case LUA_TFUNCTION:
		return "function"
	case LUA_TTHREAD:
		return "thread"
	default:
		return "userdata"
	}
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_type

int lua_type (lua_State *L, int index);

Returns the type of the value in the given valid index, or LUA_TNONE for a
non-valid (but acceptable) index. The types returned by lua_type are coded by
the following constants defined in lua.h: LUA_TNIL (0), LUA_TNUMBER,
LUA_TBOOLEAN, LUA_TSTRING, LUA_TTABLE, LUA_TFUNCTION, LUA_TUSERDATA,
LUA_TTHREAD, and LUA_TLIGHTUSERDATA.
*/
func (self *luaState) Type(idx int) LuaType {
	if self.stack.isValid(idx) {
		val := self.stack.get(idx)
		return typeOf(val)
	}
	return LUA_TNONE
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_isnone

int lua_isnone (lua_State *L, int index);

Returns 1 if the given index is not valid, and 0 otherwise.
*/
func (self *luaState) IsNone(idx int) bool {
	return self.Type(idx) == LUA_TNONE
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_isnil

int lua_isnil (lua_State *L, int index);

Returns 1 if the value at the given index is nil, and 0 otherwise.
*/
func (self *luaState) IsNil(idx int) bool {
	return self.Type(idx) == LUA_TNIL
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_isnoneornil

int lua_isnoneornil (lua_State *L, int index);

Returns 1 if the given index is not valid or if the value at this index is nil,
and 0 otherwise.
*/
func (self *luaState) IsNoneOrNil(idx int) bool {
	return self.Type(idx) <= LUA_TNIL
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_isboolean

int lua_isboolean (lua_State *L, int index);

Returns 1 if the value at the given index is a boolean, and 0 otherwise.
*/
func (self *luaState) IsBoolean(idx int) bool {
	return self.Type(idx) == LUA_TBOOLEAN
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_istable

int lua_istable (lua_State *L, int index);

Returns 1 if the value at the given index is a table, and 0 otherwise.
*/
func (self *luaState) IsTable(idx int) bool {
	return self.Type(idx) == LUA_TTABLE
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_isfunction

int lua_isfunction (lua_State *L, int index);

Returns 1 if the value at the given index is a function (either C or Lua),
and 0 otherwise.
*/
func (self *luaState) IsFunction(idx int) bool {
	return self.Type(idx) == LUA_TFUNCTION
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_isthread

int lua_isthread (lua_State *L, int index);

Returns 1 if the value at the given index is a thread, and 0 otherwise.
*/
func (self *luaState) IsThread(idx int) bool {
	return self.Type(idx) == LUA_TTHREAD
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_isstring

int lua_isstring (lua_State *L, int index);

Returns 1 if the value at the given index is a string or a number (which is
always convertible to a string), and 0 otherwise.
*/
func (self *luaState) IsString(idx int) bool {
	t := self.Type(idx)
	return t == LUA_TSTRING || t == LUA_TNUMBER
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_isnumber

int lua_isnumber (lua_State *L, int index);

Returns 1 if the value at the given index is a number or a string convertible
to a number, and 0 otherwise.
*/
func (self *luaState) IsNumber(idx int) bool {
	_, ok := self.ToNumberX(idx)
	return ok
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_isinteger

int lua_isinteger (lua_State *L, int index);

Returns 1 if the value at the given index is an integer (that is, the value is
a number and is represented as an integer), and 0 otherwise.
*/
func (self *luaState) IsInteger(idx int) bool {
	val := self.stack.get(idx)
	_, ok := val.(int64)
	return ok
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_toboolean

int lua_toboolean (lua_State *L, int index);

Converts the Lua value at the given index to a C boolean value (0 or 1). Like
all tests in Lua, lua_toboolean returns true for any Lua value different from
false and nil; otherwise it returns false. (If you want to accept only actual
boolean values, use lua_isboolean to test the value's type.)
*/
func (self *luaState) ToBoolean(idx int) bool {
	val := self.stack.get(idx)
	return convertToBoolean(val)
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_tointeger

lua_Integer lua_tointeger (lua_State *L, int index);

Equivalent to lua_tointegerx with isnum equal to NULL.
*/
func (self *luaState) ToInteger(idx int) int64 {
	i, _ := self.ToIntegerX(idx)
	return i
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_tointegerx

lua_Integer lua_tointegerx (lua_State *L, int index, int *isnum);

Converts the Lua value at the given index to the signed integral type
lua_Integer. The Lua value must be an integer, or a number or string
convertible to an integer (see §3.4.3); otherwise, lua_tointegerx returns 0.

If isnum is not NULL, its referent is assigned a boolean value that indicates
whether the operation succeeded.
*/
func (self *luaState) ToIntegerX(idx int) (int64, bool) {
	val := self.stack.get(idx)
	i, ok := val.(int64)
	return i, ok
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_tonumber

lua_Number lua_tonumber (lua_State *L, int index);

Equivalent to lua_tonumberx with isnum equal to NULL.
*/
func (self *luaState) ToNumber(idx int) float64 {
	n, _ := self.ToNumberX(idx)
	return n
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_tonumberx

lua_Number lua_tonumberx (lua_State *L, int index, int *isnum);

Converts the Lua value at the given index to the C type lua_Number (see
lua_Number). The Lua value must be a number or a string convertible to a number
(see §3.4.3); otherwise, lua_tonumberx returns 0.

If isnum is not NULL, its referent is assigned a boolean value that indicates
whether the operation succeeded.
*/
func (self *luaState) ToNumberX(idx int) (float64, bool) {
	val := self.stack.get(idx)
	switch x := val.(type) {
	case float64:
		return x, true
	case int64:
		return float64(x), true
	default:
		return 0, false
	}
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_tostring

const char *lua_tostring (lua_State *L, int index);

Equivalent to lua_tolstring with len equal to NULL.
*/
func (self *luaState) ToString(idx int) string {
	s, _ := self.ToStringX(idx)
	return s
}

func (self *luaState) ToStringX(idx int) (string, bool) {
	val := self.stack.get(idx)

	switch x := val.(type) {
	case string:
		return x, true
	case int64, float64:
		s := fmt.Sprintf("%v", x)
		self.stack.set(idx, s)
		return s, true
	default:
		return "", false
	}
}
