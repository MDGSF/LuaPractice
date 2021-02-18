package state

import . "luago/api"

/*
http://www.lua.org/manual/5.3/manual.html#lua_compare

int lua_compare (lua_State *L, int index1, int index2, int op);

Compares two Lua values. Returns 1 if the value at index index1 satisfies op
when compared with the value at index index2, following the semantics of the
corresponding Lua operator (that is, it may call metamethods). Otherwise
returns 0. Also returns 0 if any of the indices is not valid.

The value of op must be one of the following constants:

LUA_OPEQ: compares for equality (==)
LUA_OPLT: compares for less than (<)
LUA_OPLE: compares for less or equal (<=)
*/
func (self *luaState) Compare(idx1, idx2 int, op CompareOp) bool {
	if !self.stack.isValid(idx1) || !self.stack.isValid(idx2) {
		return false
	}

	a := self.stack.get(idx1)
	b := self.stack.get(idx2)
	switch op {
	case LUA_OPEQ:
		return _eq(a, b)
	case LUA_OPLT:
		return _lt(a, b)
	case LUA_OPLE:
		return _le(a, b)
	default:
		panic("invalid compare op!")
	}
}

func _eq(a, b luaValue) bool {
	switch x := a.(type) {
	case nil:
		return b == nil
	case bool:
		y, ok := b.(bool)
		return ok && x == y
	case string:
		y, ok := b.(string)
		return ok && x == y
	case int64:
		switch y := b.(type) {
		case int64:
			return x == y
		case float64:
			return float64(x) == y
		default:
			return false
		}
	case float64:
		switch y := b.(type) {
		case float64:
			return x == y
		case int64:
			return x == float64(y)
		default:
			return false
		}
	default:
		return a == b
	}
}

func _lt(a, b luaValue) bool {
	switch x := a.(type) {
	case string:
		if y, ok := b.(string); ok {
			return x < y
		}
	case int64:
		switch y := b.(type) {
		case int64:
			return x < y
		case float64:
			return float64(x) < y
		}
	case float64:
		switch y := b.(type) {
		case float64:
			return x < y
		case int64:
			return x < float64(y)
		}
	}
	panic("comparison error!")
}

func _le(a, b luaValue) bool {
	switch x := a.(type) {
	case string:
		if y, ok := b.(string); ok {
			return x <= y
		}
	case int64:
		switch y := b.(type) {
		case int64:
			return x <= y
		case float64:
			return float64(x) <= y
		}
	case float64:
		switch y := b.(type) {
		case float64:
			return x <= y
		case int64:
			return x <= float64(y)
		}
	}
	panic("comparison error!")
}
