package binchunk

const (
	LUA_SIGNATURE    = "\x1bLua"
	LUAC_VERSION     = 0x53
	LUAC_FORMAT      = 0
	LUAC_DATA        = "\x19\x93\r\n\x1a\n"
	CINT_SIZE        = 4
	CSIZET_SIZE      = 8
	INSTRUCTION_SIZE = 4
	LUA_INTEGER_SIZE = 8
	LUA_NUMBER_SIZE  = 8
	LUAC_INT         = 0x5678
	LUAC_NUM         = 370.5
)

const (
	TAG_NIL       = 0x00
	TAG_BOOLEAN   = 0x01
	TAG_NUMBER    = 0x03
	TAG_INTEGER   = 0x13
	TAG_SHORT_STR = 0x04
	TAG_LONG_STR  = 0x14
)

// binaryChunk 二进制 chunk 总体结构
type binaryChunk struct {
	header                  // 头部
	sizeUpvalues byte       // 主函数 upvalue 数量
	mainFunc     *Prototype // 主函数原型
}

// header 头部
// xxd -u -g 1 luac.out 可以查看头部格式
type header struct {
	// 固定为 Esc L u a 这 4 个字节
	// 0x1B4C7561
	signature [4]byte

	// (Major Version).(Minor Vesion).(Release Version)
	// (Major Version) * 16 + (Minor Vesion)
	// 5.3.4 => 5 * 16 + 3 = 83
	// 83 转换为 16 进制就是 0x53
	version byte

	// 格式号，如果和虚拟机的格式号不匹配，就无法加载，默认是 0
	format byte

	// LUAC_DATA
	// 前两个字节是 0x1993，是 Lua 1.0 发布的年份
	// 后四个字节是回车符（0x0D），换行符（0x0A），替换符（0x1A），换行符（0x0A）。
	luacData [6]byte

	// cint 在二进制 chunk 中占用的字节数，本机电脑是 0x04
	cintSize byte

	// size_t 在二进制 chunk 中占用的字节数，本机电脑是 0x08
	sizetSize byte

	// Lua虚拟机指令在二进制 chunk 中占用的字节数，本机电脑是 0x04
	instructionSize byte

	// Lua整数在二进制 chunk 中占用的字节数，本机电脑是 0x08
	luaIntegerSize byte

	// Lua浮点数在二进制 chunk 中占用的字节数，本机电脑是 0x08
	luaNumberSize byte

	// LUAC_INT 保存数字 0x5678，用于检查大小端
	luacInt int64

	// LUAC_NUM 保存浮点数 370.5，用于检查浮点数格式
	luacNum float64
}

// Prototype 函数原型
type Prototype struct {
	Source string // 源文件名

	// 起止行号，主函数的这两个数字都是 0
	LineDefined     uint32
	LastLineDefined uint32

	// 固定参数个数
	NumParams byte

	// 是否有变长参数。1 代表是，0 代表否。
	IsVararg byte

	// 寄存器数量
	MaxStackSize byte

	// 指令表
	Code []uint32

	// 常量表，常量包括 nil，布尔值，整数，浮点数，字符串。
	Constants []interface{}

	// Upvalue 表
	Upvalues []Upvalue

	// 子函数原型表
	Protos []*Prototype

	// 行号表
	LineInfo []uint32

	// 局部变量表
	LocVars []LocVar

	// Upvalue名列表
	UpvalueNames []string
}

type Upvalue struct {
	Instack byte
	Idx     byte
}

// LocVar 局部变量
type LocVar struct {
	VarName string
	StartPC uint32
	EndPC   uint32
}

// Undump 用于解析二进制 chunk
func Undump(data []byte) *Prototype {
	reader := &reader{data}
	reader.checkHeader()        // 检验头部
	reader.readByte()           // 跳过 sizeUpvalues
	return reader.readProto("") // 读取函数原型
}
