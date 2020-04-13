// Copyright 2020 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gettext_test

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/chai2010/gettext-go"
)

func Example() {
	gettext := gettext.New("hello", "./examples/locale").SetLanguage("zh_CN")
	fmt.Println(gettext.Gettext("Hello, world!"))

	// Output:
	// 你好, 世界!
}

func Example_zip() {
	gettext := gettext.New("hello", "./examples/locale.zip").SetLanguage("zh_CN")
	fmt.Println(gettext.Gettext("Hello, world!"))

	// Output:
	// 你好, 世界!
}

func Example_zipData() {
	zipData, err := ioutil.ReadFile("./examples/locale.zip")
	if err != nil {
		log.Fatal(err)
	}

	gettext := gettext.New("hello", "???", zipData).SetLanguage("zh_CN")
	fmt.Println(gettext.Gettext("Hello, world!"))

	// Output:
	// 你好, 世界!
}

func Example_bind() {
	gettext.BindLocale(gettext.New("hello", "./examples/locale.zip"))
	gettext.SetLanguage("zh_CN")

	fmt.Println(gettext.Gettext("Hello, world!"))

	// Output:
	// 你好, 世界!
}

func Example_multiLang() {
	zh := gettext.New("hello", "./examples/locale").SetLanguage("zh_CN")
	tw := gettext.New("hello", "./examples/locale").SetLanguage("zh_TW")

	fmt.Println(zh.PGettext(
		"code.google.com/p/gettext-go/examples/hi.SayHi",
		"pkg hi: Hello, world!",
	))

	fmt.Println(tw.PGettext(
		"code.google.com/p/gettext-go/examples/hi.SayHi",
		"pkg hi: Hello, world!",
	))

	// Output:
	// 来自"Hi"包的问候: 你好, 世界!(ctx:code.google.com/p/gettext-go/examples/hi.SayHi)
	// 來自"Hi"包的問候: 你好, 世界!(ctx:code.google.com/p/gettext-go/examples/hi.SayHi)
}

func Example_json() {
	const jsonData = `{
		"zh_CN": {
			"LC_MESSAGES": {
				"hello.json": [{
					"msgctxt"     : "",
					"msgid"       : "Hello, world!",
					"msgid_plural": "",
					"msgstr"      : ["你好, 世界!"]
				}]
			}
		}
	}`

	gettext := gettext.New("hello", "???", jsonData).SetLanguage("zh_CN")
	fmt.Println(gettext.Gettext("Hello, world!"))

	// Output:
	// 你好, 世界!
}
