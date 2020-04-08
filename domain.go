// Copyright 2013 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gettext

import (
	"sync"
)

var defaultManager = newDomainManager()

type _DomainManager struct {
	mutex         sync.Mutex
	locale        string
	domain        string
	domainMap     map[string]FileSystem
	domainPathMap map[string]string
	trTextMap     map[string]*translator
}

func newDomainManager() *_DomainManager {
	return &_DomainManager{
		locale:        DefaultLang,
		domainMap:     make(map[string]FileSystem),
		domainPathMap: make(map[string]string),
		trTextMap:     make(map[string]*translator),
	}
}

func (p *_DomainManager) makeTrMapKey(domain, locale string) string {
	return domain + "_$$$_" + locale
}

func (p *_DomainManager) Bind(domain, path string, data interface{}) (domains, paths []string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	switch {
	case domain != "" && path != "": // bind new domain
		p.bindDomainTranslators(domain, path, data)
		p.domainPathMap[domain] = path
	case domain != "" && path == "": // delete domain
		p.deleteDomain(domain)
		delete(p.domainPathMap, domain)
	}

	// return all bind domain
	for k := range p.domainMap {
		domains = append(domains, k)
	}
	for _, s := range domains {
		paths = append(paths, p.domainPathMap[s])
	}
	return
}

func (p *_DomainManager) SetLocale(locale string) string {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if locale != "" {
		p.locale = locale
	}
	return p.locale
}

func (p *_DomainManager) SetDomain(domain string) string {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if domain != "" {
		p.domain = domain
	}
	return p.domain
}

func (p *_DomainManager) Getdata(name string) []byte {
	return p.getdata(p.domain, name)
}

func (p *_DomainManager) DGetdata(domain, name string) []byte {
	return p.getdata(domain, name)
}

func (p *_DomainManager) PNGettext(msgctxt, msgid, msgidPlural string, n int) string {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.gettext(p.domain, msgctxt, msgid, msgidPlural, n)
}

func (p *_DomainManager) DPNGettext(domain, msgctxt, msgid, msgidPlural string, n int) string {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.gettext(domain, msgctxt, msgid, msgidPlural, n)
}

func (p *_DomainManager) gettext(domain, msgctxt, msgid, msgidPlural string, n int) string {
	if p.locale == "" || p.domain == "" {
		return msgid
	}
	if _, ok := p.domainMap[domain]; !ok {
		return msgid
	}
	if f, ok := p.trTextMap[p.makeTrMapKey(domain, p.locale)]; ok {
		return f.PNGettext(msgctxt, msgid, msgidPlural, n)
	}
	return msgid
}

func (p *_DomainManager) getdata(domain, name string) []byte {
	if p.locale == "" || p.domain == "" {
		return nil
	}
	if _, ok := p.domainMap[domain]; !ok {
		return nil
	}
	if fs, ok := p.domainMap[domain]; ok {
		if data, err := fs.LoadResourceFile(domain, p.locale, name); err == nil {
			return data
		}
		if p.locale != "default" {
			if data, err := fs.LoadResourceFile(domain, "default", name); err == nil {
				return data
			}
		}
	}
	return nil
}
