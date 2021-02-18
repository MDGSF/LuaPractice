package state

/*
    +-----------------------+
    |                       |  6 (isvalid)
    +-----------------------+
-1  |          e            |  5 (top)
    +-----------------------+
-2  |          d            |  4
    +-----------------------+
-3  |          c            |  3
    +-----------------------+
-4  |          b            |  2
    +-----------------------+
-5  |          a            |  1
    +-----------------------+
*/
type luaStack struct {
	slots []luaValue
	top   int
}

// newLuaStack 创建一个虚拟栈
func newLuaStack(size int) *luaStack {
	return &luaStack{
		slots: make([]luaValue, size),
		top:   0,
	}
}

// check 检查虚拟栈中是否还可以容纳至少 n 个值，不满足则扩容。
func (self *luaStack) check(n int) {
	free := len(self.slots) - self.top
	for i := free; i < n; i++ {
		self.slots = append(self.slots, nil)
	}
}

func (self *luaStack) push(val luaValue) {
	if self.top == len(self.slots) {
		panic("stack overflow!")
	}
	self.slots[self.top] = val
	self.top++
}

func (self *luaStack) pop() luaValue {
	if self.top < 1 {
		panic("stack underflow!")
	}
	self.top--
	val := self.slots[self.top]
	self.slots[self.top] = nil
	return val
}

// absIndex 把索引转换成绝对索引
func (self *luaStack) absIndex(idx int) int {
	if idx >= 0 {
		return idx
	}
	return idx + self.top + 1
}

// isValid 判断索引是否有效
func (self *luaStack) isValid(idx int) bool {
	absIdx := self.absIndex(idx)
	return absIdx > 0 && absIdx <= self.top
}

// get 根据索引从栈里取值，如果索引无效则返回 nil 值
func (self *luaStack) get(idx int) luaValue {
	absIdx := self.absIndex(idx)
	if absIdx > 0 && absIdx <= self.top {
		return self.slots[absIdx-1]
	}
	return nil
}

// set 根据索引往栈里写值
func (self *luaStack) set(idx int, val luaValue) {
	absIdx := self.absIndex(idx)
	if absIdx > 0 && absIdx <= self.top {
		self.slots[absIdx-1] = val
		return
	}
	panic("invalid index!")
}

func (self *luaStack) reverse(from, to int) {
	slots := self.slots
	for from < to {
		slots[from], slots[to] = slots[to], slots[from]
		from++
		to--
	}
}
