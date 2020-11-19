// Copyright 2020 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
)

func Example_gopkgPath() {
	fmt.Println(gopkgPath("."))
	fmt.Println(gopkgPath("../.."))
	fmt.Println(gopkgPath("../..//examples/hi"))

	// Output:
	// github.com/ContextLogic/goi18n/cmd/xgettext-go
	// github.com/ContextLogic/goi18n
	// github.com/ContextLogic/goi18n/examples/hi
}

func Example_gopkgFiles() {
	fmt.Println(gopkgFiles("."))
	fmt.Println(gopkgFiles("../../examples/hi"))

	// Output:
	// [main.go pkg.go utils.go]
	// [hi.go]
}
