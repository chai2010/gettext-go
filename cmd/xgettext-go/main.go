// Copyright 2020 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

// The xgettext-go program extracts translatable strings from Go packages.
package main

import (
    "flag"
    "fmt"
    "os"
)

var (
    d string
    D string
    f string
    o string
    p string
    V bool
)

func version()  {
    fmt.Fprintf(os.Stdout,"%s %s\n",os.Args[0],"v0.0.1")
}

func main() {
    flag.Usage = func() {
        usageText := `xgettext [] []...

...
  -f, --files-from=       <>
  -D, --directory=        <>
 -

  -d, --default-domain=   <.po>( messages.po)
  -o, --output=
  -p, --output-dir=       <>
 -

  -h, --help
  -V, --version

report to <xxx@xxx.xx>`
        fmt.Fprintf(os.Stderr, "%s\n",usageText)
    }
    flag.StringVar(&d,"d","","")
    flag.StringVar(&D,"D","","")
    flag.StringVar(&f,"f","","")
    flag.StringVar(&o,"o","","")
    flag.StringVar(&p,"p","","")
    flag.BoolVar(&V,"V",false,"v0.0.1")
    flag.Parse()
    if V {
        version()
    }
    // TODO://
}
