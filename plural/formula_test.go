// Copyright 2013 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plural

import (
	"testing"
)

func TestFormula(t *testing.T) {
	for _, v := range pluralFormHeaders {
		formulaFunc, err := Formula(v.pluralForm)
		if err != nil {
			t.Fatalf("%d - %s: Unexpected error: %s", v.in, v.pluralForm, err.Error())
		}
		if out := formulaFunc(v.in); out != v.out {
			t.Fatalf("%d - %s: expect = %d, got = %d", v.in, v.pluralForm, v.out, out)
		}
	}
}

var arPluralForm = " nplurals=6; plural= n==0 ? 0 : n==1 ? 1 : n==2 ? 2 : n%100>=3 && n%100<=10 ? 3 : n%100>=11 ? 4 : 5;\n"
var ukPluralForm = "nplurals=3; plural=n%10==1 && n%100!=11 ? 0 : n%10>=2 && n%10<=4 && (n%100<10 || n%100>=20) ? 1 : 2;\n"
var frPluralForm = " nplurals=2; plural=n>1;"

var pluralFormHeaders = []struct {
	pluralForm string
	in   uint32
	out  int
}{
	// # Arabic - 6 forms
	{arPluralForm, 0, 0},
	{arPluralForm, 1, 1},
	{arPluralForm, 2, 2},
	{arPluralForm, 5, 3},
	{arPluralForm, 12, 4},
	{arPluralForm, 102, 5},

	// Ukrainian - 3 forms
	{ukPluralForm, 0, 2},
	{ukPluralForm, 1, 0},
	{ukPluralForm, 2, 1},
	{ukPluralForm, 5, 2},
	{ukPluralForm, 100, 2},
	{ukPluralForm, 32, 1},

	// French - 2 forms
	{frPluralForm, 0, 0},
	{frPluralForm, 1, 0},
	{frPluralForm, 5, 1},
	{frPluralForm, 100, 1},
}

func TestFormulaByLang(t *testing.T) {
	for i, v := range testData {
		if out := FormulaByLang(v.lang)(v.in); out != v.out {
			t.Fatalf("%d/%s: expect = %d, got = %d", i, v.lang, v.out, out)
		}
	}
}

var testData = []struct {
	lang string
	in   uint32
	out  int
}{
	{"#@", 0, 0},
	{"#@", 1, 0},
	{"#@", 10, 0},

	{"zh", 0, 0},
	{"zh", 1, 0},
	{"zh", 10, 0},

	{"zh_CN", 0, 0},
	{"zh_CN", 1, 0},
	{"zh_CN", 10, 0},

	{"en", 0, 0},
	{"en", 1, 0},
	{"en", 2, 1},
	{"en", 10, 1},

	{"en_US", 0, 0},
	{"en_US", 1, 0},
	{"en_US", 2, 1},
	{"en_US", 10, 1},
}
