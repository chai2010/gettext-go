// Copyright 2013 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gettext

import (
	go_locale "github.com/rocket049/go-locale"
)

func getDefaultLocale() string {
	loc, err := go_locale.DetectLocale()
	if err != nil {
		return "default"
	}
	return loc
}
