// Copyright 2013 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gettext

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type osFS struct {
	root string
}

func newOsFS(root string) FileSystem {
	// local zip file
	if fi, err := os.Stat(root); err == nil && !fi.IsDir() {
		if strings.HasSuffix(strings.ToLower(root), ".zip") {
			if x, err := ioutil.ReadFile(root); err == nil {
				if r, err := zip.NewReader(bytes.NewReader(x), int64(len(x))); err == nil {
					return ZipFS(r, root)
				}
			}
		}
	}

	// local dir
	return &osFS{root: root}
}

func (p *osFS) LocaleList() []string {
	list, err := ioutil.ReadDir(p.root)
	if err != nil {
		return nil
	}
	ssMap := make(map[string]bool)
	for _, dir := range list {
		if dir.IsDir() {
			ssMap[dir.Name()] = true
		}
	}
	var locals = make([]string, 0, len(ssMap))
	for s := range ssMap {
		locals = append(locals, s)
	}
	sort.Strings(locals)
	return locals
}

func (p *osFS) LoadMessagesFile(domain, local, ext string) ([]byte, error) {
	trName := p.makeMessagesFileName(domain, local, ext)
	rcData, err := ioutil.ReadFile(trName)
	if err != nil {
		return nil, err
	}
	return rcData, nil
}

func (p *osFS) LoadResourceFile(domain, local, name string) ([]byte, error) {
	rcName := p.makeResourceFileName(domain, local, name)
	rcData, err := ioutil.ReadFile(rcName)
	if err != nil {
		return nil, err
	}
	return rcData, nil
}

func (p *osFS) String() string {
	return "gettext.localfs(" + p.root + ")"
}

func (p *osFS) makeMessagesFileName(domain, local, ext string) string {
	return fmt.Sprintf("%s/%s/LC_MESSAGES/%s%s", p.root, local, domain, ext)
}

func (p *osFS) makeResourceFileName(domain, local, name string) string {
	return fmt.Sprintf("%s/%s/LC_RESOURCE/%s/%s", p.root, local, domain, name)
}
