package gettext

//T wrap of PGettext,it will translate keywords without `msgctxt`
func T(s string) string {
	return PGettext("", s)
}
