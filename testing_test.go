// Copyright 2020 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gettext

import (
	"fmt"
	"testing"
)

func print(a ...interface{}) {
	fmt.Print(a...)
}
func println(a ...interface{}) {
	fmt.Println(a...)
}
func printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func tAssert(tb testing.TB, ok bool, a ...interface{}) {
	if !ok {
		tb.Helper()
		if msg := fmt.Sprint(a...); msg != "" {
			tb.Fatal("assert failed:", msg)
		} else {
			tb.Fatal("assert failed")
		}
	}
}

func tAssertf(tb testing.TB, ok bool, format string, a ...interface{}) {
	if !ok {
		tb.Helper()
		if msg := fmt.Sprintf(format, a...); msg != "" {
			tb.Fatal("assert failed:", msg)
		} else {
			tb.Fatal("assert failed")
		}
	}
}
