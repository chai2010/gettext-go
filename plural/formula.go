// Copyright 2013 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plural

import (
	gotextp "github.com/leonelquinteros/gotext/plurals"

	"strings"
)

// Formula provides the language's standard plural formula.
func Formula(pluralFormHeader string) (func(n uint32) int, error) {
	expr, err := gotextp.Compile(pluralFormHeader)
	if err != nil {
		return nil, err
	}
	return expr.Eval, nil

}

// FormulaByLang provides the language's standard plural formula.
func FormulaByLang(lang string) func(n uint32) int {
	if idx := index(lang); idx != -1 {
		return formulaTable[fmtForms(FormsTable[idx].Value)]
	}
	if idx := index("??"); idx != -1 {
		return formulaTable[fmtForms(FormsTable[idx].Value)]
	}
	return func(n uint32) int {
		return int(n)
	}
}

func index(lang string) int {
	for i := 0; i < len(FormsTable); i++ {
		if strings.HasPrefix(lang, FormsTable[i].Lang) {
			return i
		}
	}
	return -1
}

func fmtForms(forms string) string {
	forms = strings.TrimSpace(forms)
	forms = strings.Replace(forms, " ", "", -1)
	return forms
}

var formulaTable = map[string]func(n uint32) int{
	fmtForms("nplurals=n; plural=n-1;"): func(n uint32) int {
		if n > 0 {
			return int(n - 1)
		}
		return 0
	},
	fmtForms("nplurals=1; plural=0;"): func(n uint32) int {
		return 0
	},
	fmtForms("nplurals=2; plural=(n != 1);"): func(n uint32) int {
		if n <= 1 {
			return 0
		}
		return 1
	},
	fmtForms("nplurals=2; plural=(n > 1);"): func(n uint32) int {
		if n <= 1 {
			return 0
		}
		return 1
	},
	fmtForms("nplurals=3; plural=(n%10==1 && n%100!=11 ? 0 : n != 0 ? 1 : 2);"): func(n uint32) int {
		if n%10 == 1 && n%100 != 11 {
			return 0
		}
		if n != 0 {
			return 1
		}
		return 2
	},
	fmtForms("nplurals=3; plural=n==1 ? 0 : n==2 ? 1 : 2;"): func(n uint32) int {
		if n == 1 {
			return 0
		}
		if n == 2 {
			return 1
		}
		return 2
	},
	fmtForms("nplurals=3; plural=n==1 ? 0 : (n==0 || (n%100 > 0 && n%100 < 20)) ? 1 : 2;"): func(n uint32) int {
		if n == 1 {
			return 0
		}
		if n == 0 || (n%100 > 0 && n%100 < 20) {
			return 1
		}
		return 2
	},
	fmtForms("nplurals=3; plural=(n%10==1 && n%100!=11 ? 0 : n%10>=2 && (n%100<10 || n%100>=20) ? 1 : 2);"): func(n uint32) int {
		if n%10 == 1 && n%100 != 11 {
			return 0
		}
		if n%10 >= 2 && (n%100 < 10 || n%100 >= 20) {
			return 1
		}
		return 2
	},
	fmtForms("nplurals=3; plural=(n%10==1 && n%100!=11 ? 0 : n%10>=2 && n%10<=4 && (n%100<10 || n%100>=20) ? 1 : 2);"): func(n uint32) int {
		if n%10 == 1 && n%100 != 11 {
			return 0
		}
		if n%10 >= 2 && n%10 <= 4 && (n%100 < 10 || n%100 >= 20) {
			return 1
		}
		return 2
	},
	fmtForms("nplurals=3; plural=(n%10==1 && n%100!=11 ? 0 : n%10>=2 && n%10<=4 && (n%100<10 || n%100>=20) ? 1 : 2);"): func(n uint32) int {
		if n%10 == 1 && n%100 != 11 {
			return 0
		}
		if n%10 >= 2 && n%10 <= 4 && (n%100 < 10 || n%100 >= 20) {
			return 1
		}
		return 2
	},
	fmtForms("nplurals=3; plural=(n%10==1 && n%100!=11 ? 0 : n%10>=2 && n%10<=4 && (n%100<10 || n%100>=20) ? 1 : 2);"): func(n uint32) int {
		if n%10 == 1 && n%100 != 11 {
			return 0
		}
		if n%10 >= 2 && n%10 <= 4 && (n%100 < 10 || n%100 >= 20) {
			return 1
		}
		return 2
	},
	fmtForms("nplurals=3; plural=(n%10==1 && n%100!=11 ? 0 : n%10>=2 && n%10<=4 && (n%100<10 || n%100>=20) ? 1 : 2);"): func(n uint32) int {
		if n%10 == 1 && n%100 != 11 {
			return 0
		}
		if n%10 >= 2 && n%10 <= 4 && (n%100 < 10 || n%100 >= 20) {
			return 1
		}
		return 2
	},
	fmtForms("nplurals=3; plural=(n%10==1 && n%100!=11 ? 0 : n%10>=2 && n%10<=4 && (n%100<10 || n%100>=20) ? 1 : 2);"): func(n uint32) int {
		if n%10 == 1 && n%100 != 11 {
			return 0
		}
		if n%10 >= 2 && n%10 <= 4 && (n%100 < 10 || n%100 >= 20) {
			return 1
		}
		return 2
	},
	fmtForms("nplurals=3; plural=(n==1) ? 0 : (n>=2 && n<=4) ? 1 : 2;"): func(n uint32) int {
		if n == 1 {
			return 0
		}
		if n >= 2 && n <= 4 {
			return 1
		}
		return 2
	},
	fmtForms("nplurals=3; plural=(n==1) ? 0 : (n>=2 && n<=4) ? 1 : 2;"): func(n uint32) int {
		if n == 1 {
			return 0
		}
		if n >= 2 && n <= 4 {
			return 1
		}
		return 2
	},
	fmtForms("nplurals=3; plural=(n==1 ? 0 : n%10>=2 && n%10<=4 && (n%100<10 || n%100>=20) ? 1 : 2);"): func(n uint32) int {
		if n == 1 {
			return 0
		}
		if n%10 >= 2 && n%10 <= 4 && (n%100 < 10 || n%100 >= 20) {
			return 1
		}
		return 2
	},
	fmtForms("nplurals=4; plural=(n%100==1 ? 0 : n%100==2 ? 1 : n%100==3 || n%100==4 ? 2 : 3);"): func(n uint32) int {
		if n%100 == 1 {
			return 0
		}
		if n%100 == 2 {
			return 1
		}
		if n%100 == 3 || n%100 == 4 {
			return 2
		}
		return 3
	},
}
