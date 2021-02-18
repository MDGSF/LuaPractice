package state

import (
	. "luago/api"
	"luago/number"
	"math"
)

type operator struct {
	integerFunc func(int64, int64) int64
	floatFunc   func(float64, float64) float64
}

var (
	iadd  = func(a, b int64) int64 { return a + b }
	fadd  = func(a, b float64) float64 { return a + b }
	isub  = func(a, b int64) int64 { return a - b }
	fsub  = func(a, b float64) float64 { return a - b }
	imul  = func(a, b int64) int64 { return a * b }
	fmul  = func(a, b float64) float64 { return a * b }
	imod  = number.IMod
	fmod  = number.FMod
	pow   = math.Pow
	div   = func(a, b float64) float64 { return a / b }
	iidiv = number.IFloorDiv
	fidiv = number.FFloorDiv
	band  = func(a, b int64) int64 { return a & b }
	bor   = func(a, b int64) int64 { return a | b }
	bxor  = func(a, b int64) int64 { return a ^ b }
	shl   = number.ShiftLeft
	shr   = number.ShiftRight
	iunm  = func(a, _ int64) int64 { return -a }
	funm  = func(a, _ float64) float64 { return -a }
	bnot  = func(a, _ int64) int64 { return ^a }
)

var operators = []operator{
	operator{iadd, fadd},
	operator{isub, fsub},
	operator{imul, fmul},
	operator{imod, fmod},
	operator{nil, pow},
	operator{nil, div},
	operator{iidiv, fidiv},
	operator{band, nil},
	operator{bor, nil},
	operator{bxor, nil},
	operator{shl, nil},
	operator{shr, nil},
	operator{iunm, funm},
	operator{bnot, nil},
}

/*
http://www.lua.org/manual/5.3/manual.html#lua_arith

void lua_arith (lua_State *L, int op);

Performs an arithmetic or bitwise operation over the two values (or one, in the
case of negations) at the top of the stack, with the value at the top being the
second operand, pops these values, and pushes the result of the operation. The
function follows the semantics of the corresponding Lua operator (that is, it
may call metamethods).

The value of op must be one of the following constants:

LUA_OPADD: performs addition (+)
LUA_OPSUB: performs subtraction (-)
LUA_OPMUL: performs multiplication (*)
LUA_OPDIV: performs float division (/)
LUA_OPIDIV: performs floor division (//)
LUA_OPMOD: performs modulo (%)
LUA_OPPOW: performs exponentiation (^)
LUA_OPUNM: performs mathematical negation (unary -)
LUA_OPBNOT: performs bitwise NOT (~)
LUA_OPBAND: performs bitwise AND (&)
LUA_OPBOR: performs bitwise OR (|)
LUA_OPBXOR: performs bitwise exclusive OR (~)
LUA_OPSHL: performs left shift (<<)
LUA_OPSHR: performs right shift (>>)
*/
func (self *luaState) Arith(op ArithOp) {
	var a, b luaValue // operands
	b = self.stack.pop()
	if op != LUA_OPUNM && op != LUA_OPBNOT {
		a = self.stack.pop()
	} else {
		a = b
	}

	operator := operators[op]
	if result := _arith(a, b, operator); result != nil {
		self.stack.push(result)
	} else {
		panic("arithmetic error!")
	}
}

func _arith(a, b luaValue, op operator) luaValue {
	if op.floatFunc == nil { // bitwise
		if x, ok := convertToInteger(a); ok {
			if y, ok := convertToInteger(b); ok {
				return op.integerFunc(x, y)
			}
		}
	} else { // arith
		if op.integerFunc != nil { // add, sub, mul, mod, idiv, unm
			if x, ok := a.(int64); ok {
				if y, ok := b.(int64); ok {
					return op.integerFunc(x, y)
				}
			}
		}
		if x, ok := convertToFloat(a); ok {
			if y, ok := convertToFloat(b); ok {
				return op.floatFunc(x, y)
			}
		}
	}
	return nil
}
