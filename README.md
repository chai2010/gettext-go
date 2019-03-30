gettext-go
==========

### Change 变化 (2019-3-30)：

Add a logger to save keywords not translated.

use `gettext.SaveLog()` to save keywords not translated to 'messages.log'.

- 用函数`gettext.SaveLog()`保存未翻译的关键词.
- 用函数`gettext.T(s string)`翻译没有`msgctxt`信息的关键字。

用法(Usage)：

```
//generate messages.log:
/*
# 2019-04-01 08:09:29

msgid "Friends"
msgstr ""

msgid "World"
msgstr ""

msgid ""
"I love\n"
"World"
msgstr ""

msgid "Hello"
msgstr ""
*/
package main

import (
	"fmt"

	"github.com/chai2010/gettext-go/gettext"
)

var T = gettext.T

func main() {
	gettext.BindTextdomain("exam", "locale", nil)
	gettext.Textdomain("exam")
	defer gettext.SaveLog()

	fmt.Println(T("Hello"))
	fmt.Println(T("Friends"))
	fmt.Println(T("World"))
	fmt.Println(T("I love\nWorld"))
}
```

PkgDoc: [http://godoc.org/github.com/chai2010/gettext-go/gettext](http://godoc.org/github.com/chai2010/gettext-go/gettext)

Install
========

1. `go get github.com/chai2010/gettext-go/gettext`
2. `go run hello.go`

The godoc.org or gowalker.org has more information.

Example
=======

```Go
package main

import (
	"fmt"

	"github.com/chai2010/gettext-go/gettext"
)

func main() {
	gettext.SetLocale("zh_CN")
	gettext.Textdomain("hello")

	gettext.BindTextdomain("hello", "local", nil)

	// gettext.BindTextdomain("hello", "local", nil)         // from local dir
	// gettext.BindTextdomain("hello", "local.zip", nil)     // from local zip file
	// gettext.BindTextdomain("hello", "local.zip", zipData) // from embedded zip data

	// translate source text
	fmt.Println(gettext.Gettext("Hello, world!"))
	// Output: 你好, 世界!

	// if no msgctxt in PO file (only msgid and msgstr),
	// specify context as "" by
	fmt.Println(gettext.PGettext("", "Hello, world!"))
	// Output: 你好, 世界!

	// translate resource
	fmt.Println(string(gettext.Getdata("poems.txt"))))
	// Output: ...
}
```

Go file: [hello.go](https://github.com/chai2010/gettext-go/blob/master/examples/hello.go); PO file: [hello.po](https://github.com/chai2010/gettext-go/blob/master/examples/local/default/LC_MESSAGES/hello.po);

BUGS
====

Please report bugs to <chaishushan@gmail.com>.

Thanks!
