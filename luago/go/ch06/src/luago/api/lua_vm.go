package api

type LuaVM interface {
	LuaState
	PC() int          // 返回当前 PC，仅测试用
	AddPC(n int)      // 修改 PC，用于实现跳转指令
	Fetch() uint32    // 取出当前指令；将 PC 指向下一条指令
	GetConst(idx int) // 用于从常量表里取出指定常量并推入栈顶
	GetRK(rk int)     // 从常量表里提取常量或者从栈里提取值，然后推入栈顶
}
