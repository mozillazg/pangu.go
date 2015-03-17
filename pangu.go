package pangu

import "regexp"

// 汉字+引号
var cjk_QUOTE_L_RE = regexp.MustCompile(
	"(?i)([\u3040-\u312f\u3200-\u32ff\u3400-\u4dbf\u4e00-\u9fff\uf900-\ufaff])([\"'])",
)

// 引号+汉字
var cjk_QUOTE_R_RE = regexp.MustCompile(
	"(?i)([\"'])([\u3040-\u312f\u3200-\u32ff\u3400-\u4dbf\u4e00-\u9fff\uf900-\ufaff])",
)

// 被引号包围的字符串
// TODO change . -> \S
var cjk_QUOTE_FIX_RE = regexp.MustCompile(
	`(?i)(["']+)(\s*)(.+?)(\s*)(["']+)`,
)

// [<{( 包围的字符串
var cjk_BRACKET_RE = regexp.MustCompile(
	"(?i)([\u3040-\u312f\u3200-\u32ff\u3400-\u4dbf\u4e00-\u9fff\uf900-\ufaff])" +
		`([<\[\{\(]+(.*?)[>\]\}\)]+)` +
		"([\u3040-\u312f\u3200-\u32ff\u3400-\u4dbf\u4e00-\u9fff\uf900-\ufaff])",
)
var cjk_BRACKETFIX_RE = regexp.MustCompile(
	`(?i)([<\[\{\(]+)(\s*)(.+?)(\s*)([>\]\}\)]+)`,
)
var cjk_BRACKET_L_RE = regexp.MustCompile(
	"(?i)([\u3040-\u312f\u3200-\u32ff\u3400-\u4dbf\u4e00-\u9fff\uf900-\ufaff])" +
		`([<>\[\]\{\}\(\)])`,
)
var cjk_BRACKET_R_RE = regexp.MustCompile(
	`(?i)([<>\[\]\{\}\(\)])` +
		"([\u3040-\u312f\u3200-\u32ff\u3400-\u4dbf\u4e00-\u9fff\uf900-\ufaff])",
)

// #
var cjk_HASH_L_RE = regexp.MustCompile(
	"(?i)([\u3040-\u312f\u3200-\u32ff\u3400-\u4dbf\u4e00-\u9fff\uf900-\ufaff])" +
		`(#(\S+))`,
)
var cjk_HASH_R_RE = regexp.MustCompile(
	`(?i)((\S+)#)` +
		"([\u3040-\u312f\u3200-\u32ff\u3400-\u4dbf\u4e00-\u9fff\uf900-\ufaff])",
)

// 汉字字母混合
var cjk_L_RE = regexp.MustCompile(
	"(?i)([\u3040-\u312f\u3200-\u32ff\u3400-\u4dbf\u4e00-\u9fff\uf900-\ufaff])" +
		"([a-z0-9`" + `@&%=\$\^\*\-\+\|\/\\])`,
)
var cjk_R_RE = regexp.MustCompile(
	"(?i)([a-z0-9`" + `~!%&=;\|\,\.\:\?\$\^\*\-\+\/\\])` +
		"([\u3040-\u312f\u3200-\u32ff\u3400-\u4dbf\u4e00-\u9fff\uf900-\ufaff])",
)

func Spacing(s string) string {
	text := cjk_QUOTE_L_RE.ReplaceAllString(s, "$1 $2")
	text = cjk_QUOTE_R_RE.ReplaceAllString(text, "$1 $2")
	text = cjk_QUOTE_FIX_RE.ReplaceAllString(text, "$1$3$5")

	oldText := text
	newText := cjk_BRACKET_RE.ReplaceAllString(oldText, "$1 $2 $4")
	text = newText
	if oldText == newText {
		text = cjk_BRACKET_L_RE.ReplaceAllString(text, "$1 $2")
		text = cjk_BRACKET_R_RE.ReplaceAllString(text, "$1 $2")
	}
	text = cjk_BRACKETFIX_RE.ReplaceAllString(text, "$1$3$5")

	text = cjk_HASH_L_RE.ReplaceAllString(text, "$1 $2")
	text = cjk_HASH_R_RE.ReplaceAllString(text, "$1 $3")

	text = cjk_L_RE.ReplaceAllString(text, "$1 $2")
	text = cjk_R_RE.ReplaceAllString(text, "$1 $2")

	return text
}
