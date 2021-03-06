package pangu

import (
	"testing"

	"github.com/bmizerany/assert"
)

var testCases = map[string]string{
	`請問Jackie的鼻子有幾個？123個！`:    `請問 Jackie 的鼻子有幾個？123 個！`,
	`請問 Jackie 的鼻子有幾個？123 個！`: `請問 Jackie 的鼻子有幾個？123 個！`,
	`前面_後面`:                   `前面_後面`,
	`前面 _ 後面`:                 `前面 _ 後面`,
	`前面~後面`:                   `前面~ 後面`,
	`前面 ~ 後面`:                 `前面 ~ 後面`,
	`前面!後面`:                   `前面! 後面`,
	`前面?後面`:                   `前面? 後面`,
	`前面:後面`:                   `前面: 後面`,
	`前面;後面`:                   `前面; 後面`,
	`前面,後面`:                   `前面, 後面`,
	`前面.後面`:                   `前面. 後面`,
	`請@vinta吃大便`:              `請 @vinta 吃大便`,
	`請@陳上進 吃大便`:               `請 @陳上進 吃大便`,
	`前面#H2G2後面`:               `前面 #H2G2 後面`,
	`前面#銀河便車指南 後面`:            `前面 #銀河便車指南 後面`,
	`前面#銀河公車指南 #銀河大客車指南 後面`:   `前面 #銀河公車指南 #銀河大客車指南 後面`,
	`前面#銀河閃電霹靂車指南#後面`:         `前面 #銀河閃電霹靂車指南# 後面`,
	`前面#H2G2#後面`:              `前面 #H2G2# 後面`,
	`前面$後面`:                   `前面 $ 後面`,
	`前面%後面`:                   `前面 % 後面`,
	`前面^後面`:                   `前面 ^ 後面`,
	`前面&後面`:                   `前面 & 後面`,
	`前面*後面`:                   `前面 * 後面`,
	"前面`後面":                   "前面 ` 後面",
	`前面+後面`:                   `前面 + 後面`,
	`前面-後面`:                   `前面 - 後面`,
	`前面=後面`:                   `前面 = 後面`,
	`前面|後面`:                   `前面 | 後面`,
	`前面/後面`:                   `前面 / 後面`,
	`前面\\後面`:                  `前面 \\ 後面`,
	`前面(後面`:                   `前面 ( 後面`,
	`前面)後面`:                   `前面 ) 後面`,
	`前面(中文123漢字)後面`:           `前面 (中文 123 漢字) 後面`,
	`前面(中文123)後面`:             `前面 (中文 123) 後面`,
	`前面(123中文)後面`:             `前面 (123 中文) 後面`,
	`前面(中文123) then`:          `前面 (中文 123) then`,
	`前面(123中文) then`:          `前面 (123 中文) then`,
	`前面( ) then`:              `前面 ( ) then`,
	`前面[後面`:                   `前面 [ 後面`,
	`前面]後面`:                   `前面 ] 後面`,
	`前面{後面`:                   `前面 { 後面`,
	`前面}後面`:                   `前面 } 後面`,
	`前面<後面`:                   `前面 < 後面`,
	`前面>後面`:                   `前面 > 後面`,
	`前面'中文123漢字'後面`:           `前面 '中文 123 漢字' 後面`,
	`前面' '後面`:                 `前面 ' ' 後面`,
	`前面"中文123漢字"後面`:           `前面 "中文 123 漢字" 後面`,
	`前面" "後面`:                 `前面 " " 後面`,
}

func TestSpacing(t *testing.T) {
	for old, s := range testCases {
		assert.Equal(t, Spacing(old), s)
	}
}
