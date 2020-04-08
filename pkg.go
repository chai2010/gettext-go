// Copyright 2020 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gettext

import (
	"sort"
	"sync"
)

type _DonmainInfo struct {
	path     string
	fs       FileSystem
	localMap map[string]*Locale
}

var pkg struct {
	sync.Mutex

	domain string  // current domain
	lang   string  // current language
	locale *Locale // current local nor nil

	// binded domain map
	domainMap map[string]_DonmainInfo
}

func init() {
	pkg.lang = DefaultLang
	pkg.domainMap = make(map[string]_DonmainInfo)
}

func pkg_SetLocale(lang string) string {
	pkg.Lock()
	defer pkg.Unlock()

	if lang != "" && pkg.lang != lang {
		pkg.lang = lang
		pkg.locale = nil
		if x, ok := pkg.domainMap[pkg.domain]; ok {
			if x, ok := x.localMap[pkg.lang]; ok {
				pkg.locale = x
			}
		}
	}

	return pkg.lang
}

func pkg_SetDomain(domain string) string {
	pkg.Lock()
	defer pkg.Unlock()

	if domain != "" && pkg.domain != domain {
		pkg.domain = domain
		pkg.locale = nil
		if x, ok := pkg.domainMap[pkg.domain]; ok {
			if x, ok := x.localMap[pkg.lang]; ok {
				pkg.locale = x
			}
		}
	}
	return pkg.domain
}

func pkg_BindTextdomain(domain, path string, data interface{}) (domains, paths []string) {
	pkg.Lock()
	defer pkg.Unlock()

	// If the domain and the path are all empty string, don't change anything.
	// Returns is the all bind domains.

	if domain == "" && path == "" {
		for k := range pkg.domainMap {
			domains = append(domains, k)
		}
		sort.Strings(domains)
		for _, k := range domains {
			paths = append(paths, pkg.domainMap[k].path)
		}
		return
	}

	// If the domain is not empty string, but the path is the empty string,
	// delete the domain.
	// If the domain don't exists, return error.

	if domain != "" && path == "" {
		delete(pkg.domainMap, domain)
		if domain == pkg.domain {
			pkg.locale = nil
		}
		return
	}

	// If the domain and path are all not empty string, bind the new domain.
	// If the domain already exists, return error.

	var info = _DonmainInfo{
		path:     path,
		fs:       NewFS(path, data),
		localMap: make(map[string]*Locale),
	}
	for _, lang := range info.fs.LocaleList() {
		info.localMap[lang] = NewLocale(domain, path, data).SetLang(lang)
	}

	if domain != "" && pkg.domain == domain {
		pkg.locale = nil
		if x, ok := pkg.domainMap[pkg.domain]; ok {
			if x, ok := x.localMap[pkg.lang]; ok {
				pkg.locale = x
			}
		}
	}

	return
}

func pkg_Gettext(msgid string) string {
	pkg.Lock()
	defer pkg.Unlock()

	if l := pkg.locale; l != nil {
		return l.Gettext(msgid)
	}
	return msgid
}

func pkg_PGettext(msgctxt, msgid string) string {
	pkg.Lock()
	defer pkg.Unlock()

	if l := pkg.locale; l != nil {
		return l.PGettext(msgctxt, msgid)
	}
	return msgid
}

func pkg_NGettext(msgid, msgidPlural string, n int) string {
	pkg.Lock()
	defer pkg.Unlock()

	if l := pkg.locale; l != nil {
		return l.NGettext(msgid, msgidPlural, n)
	}

	return nilTranslator.PNGettext("", msgid, msgidPlural, n)
}

func pkg_PNGettext(msgctxt, msgid, msgidPlural string, n int) string {
	pkg.Lock()
	defer pkg.Unlock()

	if l := pkg.locale; l != nil {
		return l.PNGettext(msgctxt, msgid, msgidPlural, n)
	}
	return nilTranslator.PNGettext(msgctxt, msgid, msgidPlural, n)
}

func pkg_DGettext(domain, msgid string) string {
	pkg.Lock()
	defer pkg.Unlock()

	if info, ok := pkg.domainMap[domain]; ok {
		if l, ok := info.localMap[pkg.lang]; ok {
			return l.DGettext(domain, msgid)
		}
	}
	return msgid
}
func pkg_DNGettext(domain, msgid, msgidPlural string, n int) string {
	pkg.Lock()
	defer pkg.Unlock()

	if info, ok := pkg.domainMap[domain]; ok {
		if l, ok := info.localMap[pkg.lang]; ok {
			return l.DNGettext(domain, msgid, msgidPlural, n)
		}
	}
	return nilTranslator.PNGettext("", msgid, msgidPlural, n)
}

func pkg_DPGettext(domain, msgctxt, msgid string) string {
	pkg.Lock()
	defer pkg.Unlock()

	if info, ok := pkg.domainMap[domain]; ok {
		if l, ok := info.localMap[pkg.lang]; ok {
			return l.DPGettext(domain, msgctxt, msgid)
		}
	}
	return msgid
}

func pkg_DPNGettext(domain, msgctxt, msgid, msgidPlural string, n int) string {
	pkg.Lock()
	defer pkg.Unlock()

	if info, ok := pkg.domainMap[domain]; ok {
		if l, ok := info.localMap[pkg.lang]; ok {
			return l.DPNGettext(domain, msgctxt, msgid, msgidPlural, n)
		}
	}

	return nilTranslator.PNGettext(msgctxt, msgid, msgidPlural, n)
}

func pkg_Getdata(name string) []byte {
	pkg.Lock()
	defer pkg.Unlock()

	if l := pkg.locale; l != nil {
		return l.Getdata(name)
	}
	return nil
}

func pkg_DGetdata(domain, name string) []byte {
	pkg.Lock()
	defer pkg.Unlock()

	if info, ok := pkg.domainMap[domain]; ok {
		if l, ok := info.localMap[pkg.lang]; ok {
			return l.Getdata(name)
		}
	}
	return nil
}
