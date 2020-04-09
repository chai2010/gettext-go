// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"

	"github.com/chai2010/gettext-go"
)

func main() {
	gettext.SetLanguage("zh_CN")
	gettext.BindLocale(gettext.New("hello", "./examples/locale"))

	fmt.Println(gettext.Gettext("Hello, world!"))

	// Output:
	// 你好, 世界!

	func() {
		gettext := gettext.New("hello", "./examples/locale").SetLanguage("zh_TW")
		fmt.Println(gettext.PGettext("main.func", "Gettext in func."))

		// Output:
		// 閉包函數中的Gettext.(ctx:main.func)
	}()
}
