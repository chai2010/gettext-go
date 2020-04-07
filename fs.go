// Copyright 2013 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gettext

import (
	"archive/zip"
	"fmt"
)

type FileSystem interface {
	LocaleList() []string
	LoadMessagesFile(domain, local, ext string) ([]byte, error)
	LoadResourceFile(domain, local, name string) ([]byte, error)
	String() string
}

func OS(root string) FileSystem {
	return newOsFS(root)
}

func ZipFS(r *zip.Reader, name string) FileSystem {
	return newZipFS(r, name)
}

func NilFS(name string) FileSystem {
	return &nilFS{name}
}

type nilFS struct {
	name string
}

func (p *nilFS) LocaleList() []string {
	return nil
}
func (p *nilFS) LoadMessagesFile(domain, local, ext string) ([]byte, error) {
	return nil, fmt.Errorf("not found")
}
func (p *nilFS) LoadResourceFile(domain, local, name string) ([]byte, error) {
	return nil, fmt.Errorf("not found")
}
func (p *nilFS) String() string {
	return "gettext.nilfs(" + p.name + ")"
}
