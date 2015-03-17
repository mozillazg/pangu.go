# pangu.go

[![Build Status](https://travis-ci.org/mozillazg/pangu.go.svg?branch=master)](https://travis-ci.org/mozillazg/pangu.go)
[![Coverage Status](https://coveralls.io/repos/mozillazg/pangu.go/badge.svg?branch=master)](https://coveralls.io/r/mozillazg/pangu.go?branch=master)

Paranoid text spacing for good readability, to automatically insert whitespace between CJK (Chinese, Japanese, Korean), half-width English, digit and symbol characters.

* JavaScript version: [pangu.js](https://github.com/vinta/paranoid-auto-spacing>)
* Node.js version: [pangu.node](https://github.com/huei90/pangu.node)
* Python version: [pangu.py](https://github.com/vinta/pangu.py)
* Java version: [pangu.java](https://github.com/vinta/pangu.java)
* Go version: [pangu.go](https://github.com/mozillazg/pangu.go)


## Installation

```
go get github.com/mozillazg/pangu.go
```


## Usage

```go
import (
	"fmt"

	"github.com/mozillazg/pangu.go"
)

func main() {
	fmt.Println(pangu.Spacing("請問Jackie的鼻子有幾個？123個！"))
	fmt.Println(pangu.Spacing("主要成份：眼鏡95%、水3%、垃圾2%。"))
	fmt.Println(pangu.Spacing("新阿姆斯特朗炫風噴射阿姆斯特朗砲"))
	// Output:
	//請問 Jackie 的鼻子有幾個？123 個！
	//主要成份：眼鏡 95%、水 3%、垃圾 2%。
	//新阿姆斯特朗炫風噴射阿姆斯特朗砲
}
```

## Thanks

This is based on [pangu.py](https://github.com/vinta/pangu.py)
