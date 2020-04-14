// Copyright 2020 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// $ go run . ../..
// $ go run . ../../examples/hi

// The xgettext-go program extracts translatable strings from Go packages.
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/chai2010/gettext-go/po"
)

var (
	flagPotFile = flag.String("pot-file", "output.pot", "set output pot file path")
	flagHelp    = flag.Bool("h", false, "show help info")
)

func init() {
	log.SetFlags(log.Lshortfile)

	flag.Usage = func() {
		fmt.Println("usage: xgettext-go [flags] pkgpath")
		fmt.Println("       xgettext-go -pot-file=output.pot pkgpath")
		fmt.Println("       xgettext-go -h")
		fmt.Println()

		flag.PrintDefaults()
		fmt.Println()

		fmt.Println("See https://github.com/chai2010/gettext-go")
	}
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 || *flagHelp {
		flag.Usage()
		return
	}

	pkg := LoadPackage(flag.Arg(0))
	potFile := pkg.GenPotFile()

	if err := potFile.Save(*flagPotFile); err != nil {
		log.Fatal(err)
	}

	if _, err := po.LoadFile(*flagPotFile); err != nil {
		log.Println(err)
	}
}
