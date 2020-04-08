// Copyright 2020 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gettext

import (
	"testing"
)

func TestLocale(t *testing.T) {
	l := NewLocale("hello", "zh_CN", NewFS("./examples/local", nil))
	testLocal_zh_CN(t, l)

	l = NewLocale("hello", "zh_CN", NewFS("./examples/local.zip", nil))
	testLocal_zh_CN(t, l)
}

func testLocal_zh_CN(t *testing.T, l *Locale) {
	lang := l.GetLang()
	tAssert(t, lang == "zh_CN", lang)

	expect := "你好, 世界!"
	got := l.Gettext("Hello, world!")
	tAssert(t, got == expect, got, expect)

	expect = "你好, 世界!(ctx:main.main)"
	got = l.PGettext("main.main", "Hello, world!")
	tAssert(t, got == expect, got, expect)
}
