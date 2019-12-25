// Copyright 2013 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gettext

import (
	"runtime"
	"strings"
	"testing"
)

var (
	testInitCallerName0 string = callerName(1)
	testInitCallerName1 string
	testInitCallerName2 string
)

func init() {
	testInitCallerName1 = callerName(1)
}

func init() {
	testInitCallerName2 = callerName(1)
}

var tCaller = func(skip int) string {
	return callerName(skip + 1)
}

func TestCallerName(t *testing.T) {
	var name string

	// init
	name = `github.com/chai2010/gettext-go/gettext.init`
	if s := testInitCallerName0; s != name {
		t.Fatalf("expect = %s, got = %s", name, s)
	}
	name = `github.com/chai2010/gettext-go/gettext.init`
	if s := testInitCallerName1; s != name {
		t.Fatalf("expect = %s, got = %s", name, s)
	}
	name = `github.com/chai2010/gettext-go/gettext.init`
	if s := testInitCallerName2; s != name {
		t.Fatalf("expect = %s, got = %s", name, s)
	}

	// tInit -> gettext.func
	name = `github.com/chai2010/gettext-go/gettext.func`
	if s := tCaller(0); s != name {
		t.Fatalf("expect = %s, got = %s", name, s)
	}

	// caller stack
	name = `github.com/chai2010/gettext-go/gettext.callerName`
	if s := callerName(0); s != name {
		t.Fatalf("expect = %s, got = %s", name, s)
	}
	name = `github.com/chai2010/gettext-go/gettext.TestCallerName`
	if s := callerName(1); s != name {
		t.Fatalf("expect = %s, got = %s", name, s)
	}
	name = `testing.tRunner`
	if s := callerName(2); s != name {
		t.Fatalf("expect = %s, got = %s", name, s)
	}
	name = `runtime.goexit`
	if s := callerName(3); s != name {
		t.Fatalf("expect = %s, got = %s", name, s)
	}
	name = ""
	if s := callerName(4); s != name {
		t.Fatalf("expect = %s, got = %s", name, s)
	}

	// closure
	func() {
		switch s := runtime.Version(); true {
		case s == "go1.3" || strings.HasPrefix(s, "go1.3."),
			s == "go1.4" || strings.HasPrefix(s, "go1.4."),
			s == "go1.5" || strings.HasPrefix(s, "go1.5."):
			name = `github.com/chai2010/gettext-go/gettext.func`

		case s == "go1.10" || strings.HasPrefix(s, "go1.10."),
			s == "go1.11" || strings.HasPrefix(s, "go1.11."),
			s == "go1.12" || strings.HasPrefix(s, "go1.12."),
			s == "go1.13" || strings.HasPrefix(s, "go1.13."):
			name = `github.com/chai2010/gettext-go/gettext.TestCallerName.func`

		default: // tip
			name = `github.com/chai2010/gettext-go/gettext.TestCallerName.func`
		}
		if s := callerName(1); s != name {
			t.Fatalf("expect = %s, got = %s", name, s)
		}
	}()
	func() {
		func() {
			switch s := runtime.Version(); true {
			case s == "go1.3" || strings.HasPrefix(s, "go1.3."),
				s == "go1.4" || strings.HasPrefix(s, "go1.4."),
				s == "go1.5" || strings.HasPrefix(s, "go1.5."):
				name = `github.com/chai2010/gettext-go/gettext.func`

			case s == "go1.10" || strings.HasPrefix(s, "go1.10."),
				s == "go1.11" || strings.HasPrefix(s, "go1.11."),
				s == "go1.12" || strings.HasPrefix(s, "go1.12."),
				s == "go1.13" || strings.HasPrefix(s, "go1.13."):
				name = `github.com/chai2010/gettext-go/gettext.TestCallerName.func`

			default: // tip
				name = `github.com/chai2010/gettext-go/gettext.TestCallerName.func`
			}
			if s := callerName(1); s != name {
				t.Fatalf("expect = %s, got = %s", name, s)
			}
		}()
	}()
}
