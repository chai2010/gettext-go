// Copyright 2020 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gettext

import (
	"testing"
)

func TestLocale(t *testing.T) {
	l := newLocale("hello", "./examples/locale", nil).SetLanguage("zh_CN")
	testLocal_zh_CN(t, l)

	l = newLocale("hello", "./examples/locale.zip", nil).SetLanguage("zh_CN")
	testLocal_zh_CN(t, l)
}

func testLocal_zh_CN(t *testing.T, l Gettexter) {
	lang := l.GetLanguage()
	tAssert(t, lang == "zh_CN", lang)

	expect := "你好, 世界!"
	got := l.Gettext("Hello, world!")
	tAssert(t, got == expect, got, expect)

	expect = "你好, 世界!(ctx:main.main)"
	got = l.PGettext("main.main", "Hello, world!")
	tAssert(t, got == expect, got, expect)
}
