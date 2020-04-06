// Copyright 2013 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gettext

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"strings"
)

type zipFS struct {
	root string
	name string
	r    *zip.Reader
}

func newZipFS(r *zip.Reader, name string) *zipFS {
	fs := &zipFS{r: r, name: name}
	fs.root = fs.zipName()
	return fs
}

func (p *zipFS) zipName() string {
	name := p.name
	if x := strings.LastIndexAny(name, `\/`); x != -1 {
		name = name[x+1:]
	}
	name = strings.TrimSuffix(name, ".zip")
	return name
}

func (p *zipFS) LoadMessagesFile(domain, local, ext string) ([]byte, error) {

	trName := p.makeMessagesFileName(domain, local, ext)
	for _, f := range p.r.File {
		if f.Name != trName {
			continue
		}
		rc, err := f.Open()
		if err != nil {
			return nil, err
		}
		rcData, err := ioutil.ReadAll(rc)
		rc.Close()
		return rcData, err
	}
	return nil, fmt.Errorf("not found")

}

func (p *zipFS) LoadResourceFile(domain, local, name string) ([]byte, error) {
	rcName := p.makeResourceFileName(domain, local, name)
	for _, f := range p.r.File {
		if f.Name != rcName {
			continue
		}
		rc, err := f.Open()
		if err != nil {
			return nil, err
		}
		rcData, err := ioutil.ReadAll(rc)
		rc.Close()
		return rcData, err
	}
	return nil, fmt.Errorf("not found")
}

func (p *zipFS) String() string {
	return "gettext.zipfs(" + p.name + ")"
}

func (p *zipFS) makeMessagesFileName(domain, local, ext string) string {
	return fmt.Sprintf("%s/%s/LC_MESSAGES/%s%s", p.root, local, domain, ext)
}

func (p *zipFS) makeResourceFileName(domain, local, name string) string {
	return fmt.Sprintf("%s/%s/LC_RESOURCE/%s/%s", p.root, local, domain, name)
}
