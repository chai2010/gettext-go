package i18n

import (
	"path/filepath"
	"sync"
	"testing"
)

func buildTranslatorManager() (*TranslatorManager, error) {
	localeDir, err := filepath.Abs("fixtures/locale")
	if err != nil {
		return nil, err
	}
	return NewTranslatorManager(TranslatorRepositoryOptions{LocaleDir: localeDir})
}
func TestGetAvailableLocalesFromTranslatorManager(t *testing.T) {
	tm, err := buildTranslatorManager()
	if err != nil {
		t.Fatal(err)
	}

	expectedLocales := map[string]bool{
		"de-DE":   true,
		"es-LA":   true,
		"fr-FR":   true,
		"hc":      true,
		"ja-JP":   true,
		"ko-KR":   true,
		"zh-CN":   true,
		"zh-TW":   true,
		"en":      true,
		"default": true,
	}
	actualLocales := tm.GetAvailableLocales()
	if len(actualLocales) != 10 {
		t.Fatalf("Expected 10 locales. Got %d", len(actualLocales))
	}
	for _, locale := range actualLocales {
		if _, ok := expectedLocales[locale]; !ok {
			t.Fatalf("Unexpected locale %s", locale)
		}
	}
}

func TestTranslatorManagerTranslationFuncsSuite(t *testing.T) {
	TestTranslatorManagerI18n(t)
	TestTranslatorManagerNi18n(t)
	TestTranslatorManagerCi18n(t)
	TestTranslatorManagerCni18n(t)
}

func TestParallelTranslatorManagerTranslations(t *testing.T) {
	tm, err := buildTranslatorManager()
	if err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	numWorkers := len(i18nTestData)
	wg.Add(numWorkers)
	for _, v := range i18nTestData {
		go func() {
			testTranslatorManagerI18nCase(t, tm, v)
			wg.Done()
		}()
	}
	wg.Wait()
}

func testTranslatorManagerI18nCase(t *testing.T, tm *TranslatorManager, v i18nCase) {
	for j := 0; j < 200; j++ {
		var out string
		if v.args == nil {
			out = tm.I18n(v.locale, v.msg)
		} else {
			out = tm.I18n(v.locale, v.msg, v.args...)
		}
		if out != v.expectedStr {
			t.Fatalf("I18n(%s, %s, %s): expect = %s, got = %s", v.locale, v.msg, v.args, v.expectedStr, out)
		}
	}
}

func TestTranslatorManagerI18n(t *testing.T) {
	tr, err := buildTranslatorManager()
	if err != nil {
		t.Fatal(err)
	}

	if tr.I18n("Non-existent", "View tickets") != "View tickets" {
		t.Fatalf("Expected same string returned when no locale set. Got %s", tr.I18n("Non-existent", "View tickets"))
	}
	for _, v := range i18nTestData {
		var out string
		if v.args == nil {
			out = tr.I18n(v.locale, v.msg)
		} else {
			out = tr.I18n(v.locale, v.msg, v.args...)
		}
		if out != v.expectedStr {
			t.Fatalf("I18n(%s, %s, %s): expect = %s, got = %s", v.locale, v.msg, v.args, v.expectedStr, out)
		}
	}
}

func TestTranslatorManagerNi18n(t *testing.T) {
	tr, err := buildTranslatorManager()
	if err != nil {
		t.Fatal(err)
	}

	case1 := tr.Ni18n("Non-existent", 1, "1 produit associé", "{%1=number of products} produit associé", 1)
	if case1 != "1 produit associé" {
		t.Fatalf("Expected singular string returned when no locale set. Got %s", case1)
	}
	case2 := tr.Ni18n("Non-existent", 5, "1 produit associé", "{%1=number of products} produits associés", 5)
	if case2 != "5 produits associés" {
		t.Fatalf("Expected plural string returned when no locale set. Got %s", case2)
	}

	for _, v := range Ni18nTestData {
		var out string
		if v.args == nil {
			out = tr.Ni18n(v.locale, v.n, v.msg, v.msgPlural)
		} else {
			out = tr.Ni18n(v.locale, v.n, v.msg, v.msgPlural, v.args...)
		}
		if out != v.expectedStr {
			t.Fatalf("Ni18n(%s, %d, %s, %s, %s): expect = %s, got = %s", v.locale, v.n, v.msg, v.msgPlural, v.args, v.expectedStr, out)
		}
	}
}

func TestTranslatorManagerCi18n(t *testing.T) {
	tr, err := buildTranslatorManager()
	if err != nil {
		t.Fatal(err)
	}

	case1 := tr.Ci18n("Non-existent", "LEGAL_CONSTANTS", "Terms of Use")
	if case1 != "Terms of Use" {
		t.Fatalf("Expected same string returned when no locale set. Got %s", case1)
	}

	for _, v := range Ci18nTestData {
		var out string
		if v.args == nil {
			out = tr.Ci18n(v.locale, v.ctx, v.msg)
		} else {
			out = tr.Ci18n(v.locale, v.ctx, v.msg, v.args...)
		}
		if out != v.expectedStr {
			t.Fatalf("Ci18n(%s, %s, %s, %s): expect = %s, got = %s", v.locale, v.ctx, v.msg, v.args, v.expectedStr, out)
		}
	}
}

func TestTranslatorManagerCni18n(t *testing.T) {
	tr, err := buildTranslatorManager()
	if err != nil {
		t.Fatal(err)
	}

	case1 := tr.Cni18n("Non-existent", "", 1, "1 linked product", "{%1=number of products} linked products", 1)
	if case1 != "1 linked product" {
		t.Fatalf("Expected singular string returned when no locale set. Got %s", case1)
	}
	case2 := tr.Cni18n("Non-existent", "", 5, "1 linked product", "{%1=number of products} linked products", 5)
	if case2 != "5 linked products" {
		t.Fatalf("Expected plural string returned when no locale set. Got %s", case2)
	}

	for _, v := range Cni18nTestData {
		var out string
		if v.args == nil {
			out = tr.Cni18n(v.locale, v.ctx, v.n, v.msg, v.msgPlural)
		} else {
			out = tr.Cni18n(v.locale, v.ctx, v.n, v.msg, v.msgPlural, v.args...)
		}
		if out != v.expectedStr {
			t.Fatalf("Cni18n(%s, %s, %d, %s, %s, %s): expect = %s, got = %s", v.locale, v.ctx, v.n, v.msg, v.msgPlural, v.args, v.expectedStr, out)
		}
	}
}
