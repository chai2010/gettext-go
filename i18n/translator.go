package i18n

import (
	"fmt"
	gettext "github.com/ContextLogic/goi18n"
	"regexp"
	"strconv"
)

type Translator struct {
	gettexter *gettext.Gettexter
}

// I18n - Translate a string into this Translator's locale
func (tr *Translator) I18n(str string, args ...interface{}) string {
	gettexter := *tr.gettexter
	translatedStr := gettexter.Gettext(str)
	return formatStr(translatedStr, args...)
}

// Ci18n - Translate a string with context into this Translator's locale
func (tr *Translator) Ci18n(ctx, str string, args ...interface{}) string {
	gettexter := *tr.gettexter
	translatedStr := gettexter.PGettext(ctx, str)
	return formatStr(translatedStr, args...)
}

// Ni18n - Translate a string with plural form into this Translator's locale
func (tr *Translator) Ni18n(n int, str, pluralStr string, args ...interface{}) string {
	gettexter := *tr.gettexter
	translatedStr := gettexter.NGettext(str, pluralStr, n)
	return formatStr(translatedStr, args...)
}

// Cni18n - Translate a string with plural form and context into this Translator's locale
func (tr *Translator) Cni18n(ctx string, n int, str, pluralStr string, args ...interface{}) string {
	gettexter := *tr.gettexter
	translatedStr := gettexter.PNGettext(ctx, str, pluralStr, n)
	return formatStr(translatedStr, args...)
}

func formatStr(str string, args ...interface{}) string {
	if args == nil || len(args) == 0 {
		return str
	}

	return subDescriptivePlaceholders(str, args...)
}

func subDescriptivePlaceholders(str string, args ...interface{}) string {
	re := regexp.MustCompile(`\{%(\d+)=[^{}]+\}`)
	return re.ReplaceAllStringFunc(
		str,
		func(placeholder string) string {
			subMatches := re.FindStringSubmatch(placeholder)
			if subMatches != nil && len(subMatches) == 2 {
				argIndex, err := strconv.Atoi(subMatches[1])
				if err == nil {
					return fmt.Sprintf("%v", args[argIndex-1])
				}
			}
			return placeholder
		},
	)
}

