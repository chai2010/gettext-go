// Copyright 2013 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gettext

type Gettexter interface {
	Gettext(msgid string) string
	PGettext(msgctxt, msgid string) string

	NGettext(msgid, msgidPlural string, n int) string
	PNGettext(msgctxt, msgid, msgidPlural string, n int) string

	DGettext(domain, msgid string) string
	DPGettext(domain, msgctxt, msgid string) string
	DNGettext(domain, msgid, msgidPlural string, n int) string
	DPNGettext(domain, msgctxt, msgid, msgidPlural string, n int) string

	Getdata(name string) []byte
	DGetdata(domain, name string) []byte
}

var debug = false

// SetLocale sets and queries the program's current lang.
//
// If the lang is not empty string, set the new local.
//
// If the lang is empty string, don't change anything.
//
// Returns is the current locale.
//
// Examples:
//	SetLocale("")      // get locale: return DefaultLocale
//	SetLocale("zh_CN") // set locale: return zh_CN
//	SetLocale("")      // get locale: return zh_CN
func SetLocale(lang string) string {
	if !debug {
		return defaultManager.SetLocale(lang)
	}
	return pkg_SetLocale(lang)
}

// BindTextdomain sets and queries program's domains.
//
// If the domain and path are all not empty string, bind the new domain.
// If the domain already exists, return error.
//
// If the domain is not empty string, but the path is the empty string,
// delete the domain.
// If the domain don't exists, return error.
//
// If the domain and the path are all empty string, don't change anything.
// Returns is the all bind domains.
//
// Examples:
//	BindTextdomain("poedit", "local", nil) // bind "poedit" domain
//	BindTextdomain("", "", nil)            // return all domains
//	BindTextdomain("poedit", "", nil)      // delete "poedit" domain
//	BindTextdomain("", "", nil)            // return all domains
//
// Use zip file:
//	BindTextdomain("poedit", "local.zip", nil)     // bind "poedit" domain
//	BindTextdomain("poedit", "local.zip", zipData) // bind "poedit" domain
//
// Use FileSystem:
//	BindTextdomain("poedit", "name", OS("path/to/dir")) // bind "poedit" domain
//	BindTextdomain("poedit", "name", OS("path/to.zip")) // bind "poedit" domain
//
func BindTextdomain(domain, path string, data interface{}) (domains, paths []string) {
	if !debug {
		return defaultManager.Bind(domain, path, data)
	}
	return pkg_BindTextdomain(domain, path, data)
}

// Textdomain sets and retrieves the current message domain.
//
// If the domain is not empty string, set the new domains.
//
// If the domain is empty string, don't change anything.
//
// Returns is the all used domains.
//
// Examples:
//	Textdomain("poedit") // set domain: poedit
//	Textdomain("")       // get domain: return poedit
func Textdomain(domain string) string {
	if !debug {
		return defaultManager.SetDomain(domain)
	}
	return pkg_SetDomain(domain)
}

// Gettext attempt to translate a text string into the user's native language,
// by looking up the translation in a message catalog.
//
// It use the caller's function name as the msgctxt.
//
// Examples:
//	func Foo() {
//		msg := gettext.Gettext("Hello") // msgctxt is "some/package/name.Foo"
//	}
func Gettext(msgid string) string {
	if !debug {
		return PGettext("", msgid)
	}
	return pkg_Gettext(msgid)
}

// Getdata attempt to translate a resource file into the user's native language,
// by looking up the translation in a message catalog.
//
// Examples:
//	func Foo() {
//		Textdomain("hello")
//		BindTextdomain("hello", "local.zip", nilOrZipData)
//		poems := gettext.Getdata("poems.txt")
//	}
func Getdata(name string) []byte {
	if !debug {
		return defaultManager.Getdata(name)
	}
	return pkg_Getdata(name)
}

// NGettext attempt to translate a text string into the user's native language,
// by looking up the appropriate plural form of the translation in a message
// catalog.
//
// It use the caller's function name as the msgctxt.
//
// Examples:
//	func Foo() {
//		msg := gettext.NGettext("%d people", "%d peoples", 2)
//	}
func NGettext(msgid, msgidPlural string, n int) string {
	if !debug {
		return PNGettext("", msgid, msgidPlural, n)
	}
	return pkg_NGettext(msgid, msgidPlural, n)
}

// PGettext attempt to translate a text string into the user's native language,
// by looking up the translation in a message catalog.
//
// Examples:
//	func Foo() {
//		msg := gettext.PGettext("gettext-go.example", "Hello") // msgctxt is "gettext-go.example"
//	}
func PGettext(msgctxt, msgid string) string {
	if !debug {
		return PNGettext(msgctxt, msgid, "", 0)
	}
	return pkg_PGettext(msgctxt, msgid)
}

// PNGettext attempt to translate a text string into the user's native language,
// by looking up the appropriate plural form of the translation in a message
// catalog.
//
// Examples:
//	func Foo() {
//		msg := gettext.PNGettext("gettext-go.example", "%d people", "%d peoples", 2)
//	}
func PNGettext(msgctxt, msgid, msgidPlural string, n int) string {
	if !debug {
		return defaultManager.PNGettext(msgctxt, msgid, msgidPlural, n)
	}
	return pkg_PNGettext(msgctxt, msgid, msgidPlural, n)
}

// DGettext like Gettext(), but looking up the message in the specified domain.
//
// Examples:
//	func Foo() {
//		msg := gettext.DGettext("poedit", "Hello")
//	}
func DGettext(domain, msgid string) string {
	if !debug {
		return DPGettext(domain, "", msgid)
	}
	return pkg_DGettext(domain, msgid)
}

// DNGettext like NGettext(), but looking up the message in the specified domain.
//
// Examples:
//	func Foo() {
//		msg := gettext.PNGettext("poedit", "gettext-go.example", "%d people", "%d peoples", 2)
//	}
func DNGettext(domain, msgid, msgidPlural string, n int) string {
	if !debug {
		return DPNGettext(domain, "", msgid, msgidPlural, n)
	}
	return pkg_DNGettext(domain, msgid, msgidPlural, n)
}

// DPGettext like PGettext(), but looking up the message in the specified domain.
//
// Examples:
//	func Foo() {
//		msg := gettext.DPGettext("poedit", "gettext-go.example", "Hello")
//	}
func DPGettext(domain, msgctxt, msgid string) string {
	if !debug {
		return DPNGettext(domain, msgctxt, msgid, "", 0)
	}
	return pkg_DPGettext(domain, msgctxt, msgid)
}

// DPNGettext like PNGettext(), but looking up the message in the specified domain.
//
// Examples:
//	func Foo() {
//		msg := gettext.DPNGettext("poedit", "gettext-go.example", "%d people", "%d peoples", 2)
//	}
func DPNGettext(domain, msgctxt, msgid, msgidPlural string, n int) string {
	if !debug {
		return defaultManager.DPNGettext(domain, msgctxt, msgid, msgidPlural, n)
	}
	return pkg_DPNGettext(domain, msgctxt, msgid, msgidPlural, n)
}

// DGetdata like Getdata(), but looking up the resource in the specified domain.
//
// Examples:
//	func Foo() {
//		msg := gettext.DGetdata("hello", "poems.txt")
//	}
func DGetdata(domain, name string) []byte {
	if !debug {
		return defaultManager.DGetdata(domain, name)
	}
	return pkg_DGetdata(domain, name)
}
