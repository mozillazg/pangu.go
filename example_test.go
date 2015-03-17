package pangu_test

import (
	"fmt"

	"github.com/mozillazg/pangu.go"
)

func ExampleSpacing() {
	fmt.Println(pangu.Spacing("請問Jackie的鼻子有幾個？123個！"))
	fmt.Println(pangu.Spacing("主要成份：眼鏡95%、水3%、垃圾2%。"))
	fmt.Println(pangu.Spacing("新阿姆斯特朗炫風噴射阿姆斯特朗砲"))
	// Output:
	//請問 Jackie 的鼻子有幾個？123 個！
	//主要成份：眼鏡 95%、水 3%、垃圾 2%。
	//新阿姆斯特朗炫風噴射阿姆斯特朗砲
}
