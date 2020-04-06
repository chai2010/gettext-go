// Copyright 2013 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gettext

import (
	"fmt"
	"io/ioutil"
)

type localFS struct {
	root string
}

func newLocalFS(root string) *localFS {
	return &localFS{root: root}
}

func (p *localFS) LoadMessagesFile(domain, local, ext string) ([]byte, error) {
	trName := p.makeMessagesFileName(domain, local, ext)
	rcData, err := ioutil.ReadFile(trName)
	if err != nil {
		return nil, err
	}
	return rcData, nil
}

func (p *localFS) LoadResourceFile(domain, local, name string) ([]byte, error) {
	rcName := p.makeResourceFileName(domain, local, name)
	rcData, err := ioutil.ReadFile(rcName)
	if err != nil {
		return nil, err
	}
	return rcData, nil
}

func (p *localFS) String() string {
	return "gettext.localfs(" + p.root + ")"
}

func (p *localFS) makeMessagesFileName(domain, local, ext string) string {
	return fmt.Sprintf("%s/%s/LC_MESSAGES/%s%s", p.root, local, domain, ext)
}

func (p *localFS) makeResourceFileName(domain, local, name string) string {
	return fmt.Sprintf("%s/%s/LC_RESOURCE/%s/%s", p.root, local, domain, name)
}
