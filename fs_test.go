// Copyright 2020 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gettext

import (
	"testing"
)

func TestFileSystem_os(t *testing.T) {
	fs := OS("./examples/locale")
	tAssert(t, fs.String() == "gettext.localfs(./examples/locale)", fs.String())

	testExamplesLocal(t, fs)
}

func TestFileSystem_zip(t *testing.T) {
	fs := NewFS("./examples/locale.zip", nil)
	tAssert(t, fs.String() == "gettext.zipfs(./examples/locale.zip)", fs.String())

	testExamplesLocal(t, fs)
}

func testExamplesLocal(t *testing.T, fs FileSystem) {
	localeList := fs.LocaleList()
	tAssert(t, len(localeList) == 3, localeList)
	tAssert(t, localeList[0] == "default")
	tAssert(t, localeList[1] == "zh_CN")
	tAssert(t, localeList[2] == "zh_TW")
}
