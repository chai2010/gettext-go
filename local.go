// Copyright 2020 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gettext

import (
	"fmt"
	"sync"
)

type Locale struct {
	mutex  sync.Mutex
	fs     FileSystem
	lang   string
	domain string
	trMap  map[string]*translator
}

var _ Gettexter = (*Locale)(nil)

func NewLocale(domain, lang string, fs FileSystem) *Locale {
	if lang == "" {
		lang = DefaultLang
	}
	p := &Locale{
		fs:    fs,
		lang:  lang,
		trMap: make(map[string]*translator),
	}
	p.SetDomain(domain)
	return p
}

func (p *Locale) makeTrMapKey(domain, locale string) string {
	return domain + "_$$$_" + locale
}

func (p *Locale) GetLang() string {
	return p.lang
}

func (p *Locale) FileSystem() FileSystem {
	return p.fs
}

func (p *Locale) SetDomain(domain string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if domain == "" || domain == p.domain {
		return
	}

	p.domain = domain
	trMapKey := p.makeTrMapKey(domain, p.lang)

	// try load po file
	if _, ok := p.trMap[trMapKey]; !ok {
		if data, err := p.fs.LoadMessagesFile(domain, p.lang, ".po"); err == nil {
			p.trMap[trMapKey], _ = newPoTranslator(
				fmt.Sprintf("%s_%s.po", domain, p.lang),
				data,
			)
		}
	}

	// try load mo file
	if _, ok := p.trMap[trMapKey]; !ok {
		if data, err := p.fs.LoadMessagesFile(domain, p.lang, ".mo"); err == nil {
			p.trMap[trMapKey], _ = newMoTranslator(
				fmt.Sprintf("%s_%s.mo", domain, p.lang),
				data,
			)
		}
	}

	// no po/mo file
	if _, ok := p.trMap[trMapKey]; !ok {
		p.trMap[trMapKey] = nilTranslator
	}

	return
}

func (p *Locale) GetDomain() string {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.domain
}

func (p *Locale) Gettext(msgid string) string {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.gettext(p.domain, "", msgid, "", 0)
}

func (p *Locale) PGettext(msgctxt, msgid string) string {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.gettext(p.domain, msgctxt, msgid, "", 0)
}

func (p *Locale) NGettext(msgid, msgidPlural string, n int) string {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.gettext(p.domain, "", msgid, msgidPlural, n)
}

func (p *Locale) PNGettext(msgctxt, msgid, msgidPlural string, n int) string {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.gettext(p.domain, msgctxt, msgid, msgidPlural, n)
}

func (p *Locale) DGettext(domain, msgid string) string {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.gettext(domain, "", msgid, "", 0)
}

func (p *Locale) DNGettext(domain, msgid, msgidPlural string, n int) string {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.gettext(domain, "", msgid, msgidPlural, n)
}

func (p *Locale) DPGettext(domain, msgctxt, msgid string) string {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.gettext(domain, msgctxt, msgid, "", 0)
}

func (p *Locale) DPNGettext(domain, msgctxt, msgid, msgidPlural string, n int) string {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.gettext(domain, msgctxt, msgid, msgidPlural, n)
}

func (p *Locale) Getdata(name string) []byte {
	return p.getdata(p.domain, name)
}

func (p *Locale) DGetdata(domain, name string) []byte {
	return p.getdata(domain, name)
}

func (p *Locale) gettext(domain, msgctxt, msgid, msgidPlural string, n int) string {
	if f, ok := p.trMap[p.makeTrMapKey(domain, p.lang)]; ok {
		return f.PNGettext(msgctxt, msgid, msgidPlural, n)
	}
	return msgid
}

func (p *Locale) getdata(domain, name string) []byte {
	if data, err := p.fs.LoadResourceFile(domain, p.lang, name); err == nil {
		return data
	}
	if p.lang != "default" {
		if data, err := p.fs.LoadResourceFile(domain, "default", name); err == nil {
			return data
		}
	}
	return nil
}
