// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This is a gettext-go exmaple.
package main

import (
	"fmt"

	"github.com/chai2010/gettext-go"
	"github.com/chai2010/gettext-go/examples/hi"
)

func init() {
	fmt.Println("=== main.init: default ===")

	// bind app domain
	gettext.BindLocale(gettext.New("hello", "locale"))

	// $(LC_MESSAGES) or $(LANG) or empty
	fmt.Println(gettext.Gettext("Gettext in init."))
	fmt.Println(gettext.PGettext("main.init", "Gettext in init."))
	hi.SayHi()

	// Output(depends on locale environment):
	// ?
	// ?
	// ?
	// ?

	fmt.Println("=== main.init: zh_CN ===")

	// set simple chinese
	gettext.SetLanguage("zh_CN")

	// simple chinese
	fmt.Println(gettext.Gettext("Gettext in init."))
	fmt.Println(gettext.PGettext("main.init", "Gettext in init."))
	hi.SayHi()

	// Output:
	// Init函数中的Gettext.
	// Init函数中的Gettext.(ctx:main.init)
	// 来自"Hi"包的问候: 你好, 世界!
	// 来自"Hi"包的问候: 你好, 世界!(ctx:code.google.com/p/gettext-go/examples/hi.SayHi)
}

func main() {
	fmt.Println("=== main.main: zh_CN ===")

	// simple chinese
	fmt.Println(gettext.Gettext("Hello, world!"))
	fmt.Println(gettext.PGettext("main.main", "Hello, world!"))
	hi.SayHi()

	// Output:
	// 你好, 世界!
	// 你好, 世界!(ctx:main.main)
	// 来自"Hi"包的问候: 你好, 世界!
	// 来自"Hi"包的问候: 你好, 世界!(ctx:code.google.com/p/gettext-go/examples/hi.SayHi)

	fmt.Println("=== main.main: zh_TW ===")

	// set traditional chinese
	gettext.SetLanguage("zh_TW")

	// traditional chinese
	func() {
		fmt.Println(gettext.Gettext("Gettext in func."))
		fmt.Println(gettext.PGettext("main.func", "Gettext in func."))
		hi.SayHi()

		// Output:
		// 閉包函數中的Gettext.
		// 閉包函數中的Gettext.(ctx:main.func)
		// 來自"Hi"包的問候: 你好, 世界!
		// 來自"Hi"包的問候: 你好, 世界!(ctx:code.google.com/p/gettext-go/examples/hi.SayHi)
	}()

	fmt.Println()

	// translate resource
	fmt.Println("=== main.main: zh_CN ===")
	gettext.SetLanguage("zh_CN")
	fmt.Println("poems(simple chinese):")
	fmt.Println(string(gettext.Getdata("poems.txt")))
	fmt.Println("=== main.main: zh_TW ===")
	gettext.SetLanguage("zh_TW")
	fmt.Println("poems(traditional chinese):")
	fmt.Println(string(gettext.Getdata("poems.txt")))
	fmt.Println("=== main.main: ?? ===")
	gettext.SetLanguage("??")
	fmt.Println("poems(default is english):")
	fmt.Println(string(gettext.Getdata("poems.txt")))
	// Output: ...
}
