# Copyright 2013 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

default:
	msgfmt -o locale/default/LC_MESSAGES/hello.mo locale/default/LC_MESSAGES/hello.po
	msgfmt -o locale/zh_CN/LC_MESSAGES/hello.mo   locale/zh_CN/LC_MESSAGES/hello.po
	msgfmt -o locale/zh_TW/LC_MESSAGES/hello.mo   locale/zh_TW/LC_MESSAGES/hello.po
	7z a -tzip -scsUTF-8 locale.zip locale
	go run hello.go

clean:
	rm locale.zip
