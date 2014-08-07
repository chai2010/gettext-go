gettext-go
==========

PkgDoc: [http://godoc.org/github.com/chai2010/gettext-go/gettext](http://godoc.org/github.com/chai2010/gettext-go/gettext)

Instasll
========

	go get github.com/chai2010/gettext-go/gettext
	cd gettext-go/gettext && go run hello.go

The godoc.org or gowalker.org has more information.

Example
=======

	// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
	// Use of this source code is governed by a BSD-style
	// license that can be found in the LICENSE file.

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

		// translate resource
		fmt.Println(string(gettext.Getdata("poems.txt"))))
		// Output: ...
	}

Go file: hello.go; PO file: hello.po;

BUGS
====

Please report bugs to <chaishushan@gmail.com>.

Thanks!
