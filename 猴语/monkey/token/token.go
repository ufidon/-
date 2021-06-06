/* 词类定义，Token，词
 */
package token

type C词类 int

const (
	F非法 C词类 = iota
	D档尾

	// 标识、值
	B标识
	Z整型
	C串型

	// 算符
	F赋
	J加
	J减
	F非
	C乘
	C除

	X小于
	D大于

	D等于
	B不等

	// 界附
	D逗号
	F分号
	M冒号

	Z左括
	Y右括
	Z左曲
	Y右曲
	Z左方
	Y右方

	// 键词
	H函
	L令
	Z真
	J假
	R若
	F否
	F返
)

var 词文 = [...][2]string{
	F非法: {"非法", "ILLEGAL"},
	D档尾: {"档尾", "EOF"},

	// 标识、值
	B标识: {"标识", "IDENT"},  // 相加, 身高, x, y, ...
	Z整型: {"整型", "INT"},    // 1343456
	C串型: {"串型", "STRING"}, // "字串"

	// 算符
	F赋: {"=", "="},
	J加: {"+", "+"},
	J减: {"-", "-"},
	F非: {"!", "!"},
	C乘: {"*", "*"},
	C除: {"/", "/"},

	X小于: {"<", "<"},
	D大于: {">", ">"},

	D等于: {"==", "=="},
	B不等: {"!=", "!="},

	// 界附
	D逗号: {",", ","},
	F分号: {";", ";"},
	M冒号: {":", ":"},

	Z左括: {"(", "("},
	Y右括: {")", ")"},
	Z左曲: {"{", "{"},
	Y右曲: {"}", "}"},
	Z左方: {"[", "["},
	Y右方: {"]", "]"},

	// 键词
	H函: {"函", "FUNCTION"},
	L令: {"令", "LET"},
	Z真: {"真", "TRUE"},
	J假: {"假", "FALSE"},
	R若: {"若", "IF"},
	F否: {"否", "ELSE"},
	F返: {"返", "RETURN"},
}

func (词类 C词类) String() string {
	return 词文[词类][0]
}

type C词 struct {
	L类 C词类
	W文 string
}

var 键词集 = map[string]C词类{
	"函":      H函,
	"令":      L令,
	"真":      Z真,
	"假":      J假,
	"若":      R若,
	"否":      F否,
	"返":      F返,
	"fn":     H函,
	"let":    L令,
	"true":   Z真,
	"false":  J假,
	"if":     R若,
	"else":   F否,
	"return": F返,
}

func C查词类(词 string) C词类 {
	if 词类, 有 := 键词集[词]; 有 {
		return 词类
	}
	return B标识
}
