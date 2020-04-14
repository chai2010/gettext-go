// Copyright 2020 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gettext

type MsgMap map[MsgMapKey]MsgMapValue

type MsgMapKey struct {
	MsgContext  string // msgctxt context
	MsgId       string // msgid untranslated-string
	MsgIdPlural string // msgid_plural untranslated-string-plural
}

type MsgMapValue struct {
	MsgStr       string   // msgstr translated-string
	MsgStrPlural []string // msgstr[0] translated-string-case-0
}

// LoadMsgMap form po/mo/json file or data.
func LoadMsgMap(path string, data ...[]byte) (MsgMap, error) {
	return make(MsgMap), nil
}

func (m MsgMap) Gettext(msgid string) string {
	return m.PGettext("", msgid)
}

func (m MsgMap) PGettext(msgctxt, msgid string) string {
	key := MsgMapKey{
		MsgContext: msgctxt,
		MsgId:      msgid,
	}
	if s := m[key].MsgStr; s != "" {
		return s
	}
	return msgid
}

func (m MsgMap) NGettext(msgid, msgidPlural string, n int) string {
	return m.PNGettext("", msgid, msgidPlural, n)
}

func (m MsgMap) PNGettext(msgctxt, msgid, msgidPlural string, n int) string {
	key := MsgMapKey{
		MsgContext:  msgctxt,
		MsgId:       msgid,
		MsgIdPlural: msgidPlural,
	}
	if v, ok := m[key]; ok {
		if n > 0 {
			if n > len(v.MsgStrPlural) {
				n = len(v.MsgStrPlural) - 1
			}
			return v.MsgStrPlural[n]
		} else {
			if v.MsgStr != "" {
				return v.MsgStr
			}
		}
	}

	if n > 0 {
		return msgidPlural
	}
	return msgid
}
