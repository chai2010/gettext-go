// Copyright 2013 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gettext

import (
	"os"
	"strings"
)

func getDefaultLocale() string {
	if v := os.Getenv("LC_MESSAGES"); v != "" {
		return simplifiedLocale(v)
	}
	if v := os.Getenv("LANG"); v != "" {
		return simplifiedLocale(v)
	}
	return "default"
}

func simplifiedLocale(lang string) string {
	// en_US/en_US.UTF-8/zh_CN/zh_TW/el_GR@euro/...
	if idx := strings.Index(lang, ":"); idx != -1 {
		lang = lang[:idx]
	}
	if idx := strings.Index(lang, "@"); idx != -1 {
		lang = lang[:idx]
	}
	if idx := strings.Index(lang, "."); idx != -1 {
		lang = lang[:idx]
	}
	return strings.TrimSpace(lang)
}

type Locale struct {
	fs     FileSystem
	locale string
	domain string
}

func NewLocale(fs FileSystem, locale string) *Locale {
	return &Locale{
		fs:     fs,
		locale: locale,
	}
}

func (l *Locale) GetDomain() string {
	return ""
}
func (l *Locale) SetDomain(domain string) {
	return
}

func (l *Locale) Getdata(name string) []byte {
	return nil
}
func (l *Locale) DGetdata(domain, name string) []byte {
	return nil
}

func (l *Locale) Gettext(msgid string) string {
	return msgid
}
func (l *Locale) DGettext(domain, msgid string) string {
	return msgid
}

func (l *Locale) NGettext(msgid, msgidPlural string, n int) string {
	return msgid
}
func (l *Locale) DNGettext(domain, msgid, msgidPlural string, n int) string {
	return msgid
}

func (l *Locale) PGettext(msgctxt, msgid string) string {
	return msgid
}
func (l *Locale) DPGettext(domain, msgctxt, msgid string) string {
	return msgid
}
func (l *Locale) DPNGettext(domain, msgctxt, msgid, msgidPlural string, n int) string {
	return msgid
}
func (l *Locale) PNGettext(msgctxt, msgid, msgidPlural string, n int) string {
	return msgid
}

func (l *Locale) String() string {
	return ""
}
