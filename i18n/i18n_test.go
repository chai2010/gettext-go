package i18n

import (
	"path/filepath"
	"sync"
	"testing"
)

func buildTranslatorRepository() (*TranslatorRepository, error) {
	localeDir, err := filepath.Abs("fixtures/locale")
	if err != nil {
		return nil, err
	}

	tr, err := NewTranslatorRepository(TranslatorRepositoryOptions{LocaleDir: localeDir})
	if err != nil {
		return nil, err
	}
	return tr, nil
}

func TestGetAvailableLocales(t *testing.T) {
	tr, err := buildTranslatorRepository()
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
	actualLocales := tr.GetAvailableLocales()
	if len(actualLocales) != 10 {
		t.Fatalf("Expected 10 locales. Got %d", len(actualLocales))
	}
	for _, locale := range actualLocales {
		if _, ok := expectedLocales[locale]; !ok {
			t.Fatalf("Unexpected locale %s", locale)
		}
	}
}

func TestTranslationFuncsSuite(t *testing.T) {
	TestI18n(t)
	TestNi18n(t)
	TestCi18n(t)
	TestCni18n(t)
}

func TestParallelTranslations(t *testing.T) {
	tr, err := buildTranslatorRepository()
	if err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	numWorkers := len(i18nTestData)
	wg.Add(numWorkers)
	for _, v := range i18nTestData {
		go func() {
			testI18nCase(t, tr, v)
			wg.Done()
		}()
	}
	wg.Wait()
}

func testI18nCase(t *testing.T, tr *TranslatorRepository, v i18nCase) {
	localeTranslator, err := tr.GetTranslator(v.locale)
	if err != nil {
		t.Fatalf("Error while setting locale %s: %s", v.locale, err.Error())
	}
	for j := 0; j < 200; j++ {
		var out string
		if v.args == nil {
			out = localeTranslator.I18n(v.msg)
		} else {
			out = localeTranslator.I18n(v.msg, v.args...)
		}
		if out != v.expectedStr {
			t.Fatalf("I18n(%s, %s): expect = %s, got = %s", v.msg, v.args, v.expectedStr, out)
		}
	}
}

func TestI18n(t *testing.T) {
	tr, err := buildTranslatorRepository()
	if err != nil {
		t.Fatal(err)
	}

	tr.currentLocale = ""
	if tr.I18n("View tickets") != "View tickets" {
		t.Fatalf("Expected same string returned when no locale set. Got %s", tr.I18n("View tickets"))
	}
	for _, v := range i18nTestData {
		err := tr.SetLocale(v.locale)
		if err != nil {
			t.Fatalf("Error while setting locale %s: %s", v.locale, err.Error())
		}
		var out string
		if v.args == nil {
			out = tr.I18n(v.msg)
		} else {
			out = tr.I18n(v.msg, v.args...)
		}
		if out != v.expectedStr {
			t.Fatalf("I18n(%s, %s): expect = %s, got = %s", v.msg, v.args, v.expectedStr, out)
		}
	}
}

func TestNi18n(t *testing.T) {
	tr, err := buildTranslatorRepository()
	if err != nil {
		t.Fatal(err)
	}

	tr.currentLocale = ""
	case1 := tr.Ni18n(1, "1 linked product", "{%1=number of products} linked products", 1)
	if case1 != "1 linked product" {
		t.Fatalf("Expected singular string returned when no locale set. Got %s", case1)
	}
	case2 := tr.Ni18n(5, "1 linked product", "{%1=number of products} linked products", 5)
	if case2 != "5 linked products" {
		t.Fatalf("Expected plural string returned when no locale set. Got %s", case2)
	}

	for _, v := range Ni18nTestData {
		err := tr.SetLocale(v.locale)
		if err != nil {
			t.Fatalf("Error while setting locale %s: %s", v.locale, err.Error())
		}
		var out string
		if v.args == nil {
			out = tr.Ni18n(v.n, v.msg, v.msgPlural)
		} else {
			out = tr.Ni18n(v.n, v.msg, v.msgPlural, v.args...)
		}
		if out != v.expectedStr {
			t.Fatalf("Ni18n(%d, %s, %s, %s): expect = %s, got = %s", v.n, v.msg, v.msgPlural, v.args, v.expectedStr, out)
		}
	}
}

func TestCi18n(t *testing.T) {
	tr, err := buildTranslatorRepository()
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range Ci18nTestData {
		err := tr.SetLocale(v.locale)
		if err != nil {
			t.Fatalf("Error while setting locale %s: %s", v.locale, err.Error())
		}
		var out string
		if v.args == nil {
			out = tr.Ci18n(v.ctx, v.msg)
		} else {
			out = tr.Ci18n(v.ctx, v.msg, v.args...)
		}
		if out != v.expectedStr {
			t.Fatalf("Ci18n(%s, %s, %s): expect = %s, got = %s", v.ctx, v.msg, v.args, v.expectedStr, out)
		}
	}
}

func TestCni18n(t *testing.T) {
	tr, err := buildTranslatorRepository()
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range Cni18nTestData {
		err := tr.SetLocale(v.locale)
		if err != nil {
			t.Fatalf("Error while setting locale %s: %s", v.locale, err.Error())
		}
		var out string
		if v.args == nil {
			out = tr.Cni18n(v.ctx, v.n, v.msg, v.msgPlural)
		} else {
			out = tr.Cni18n(v.ctx, v.n, v.msg, v.msgPlural, v.args...)
		}
		if out != v.expectedStr {
			t.Fatalf("Cni18n(%s, %d, %s, %s, %s): expect = %s, got = %s", v.ctx, v.n, v.msg, v.msgPlural, v.args, v.expectedStr, out)
		}
	}
}

type i18nCase struct {
	locale      string
	msg         string
	args        []interface{}
	expectedStr string
}

var i18nTestData = []i18nCase{
	{"fr-FR", "View tickets", nil, "Afficher les demandes d'assistance"},
	{"zh-TW", "Learn more", nil, "瞭解更多"},
	{"fr-FR", "{%1=app_name} is delivering to {%2=city}", []interface{}{"Wish", "Waterloo"}, "Wish livre à Waterloo"},
	{"fr-FR", "1 linked product", nil, "1 produit associé"},
	{"es-LA", "View tickets", nil, "View tickets"},
	{"ko-KR", "Non-existent string", nil, "Non-existent string"},
}

var Ni18nTestData = []struct {
	locale      string
	n           int
	msg         string
	msgPlural   string
	args        []interface{}
	expectedStr string
}{
	{"fr-FR", 1, "1 linked product", "{%1=number of products} linked products", []interface{}{1}, "1 produit associé"},
	{"fr-FR", 5, "1 linked product", "{%1=number of products} linked products", []interface{}{5}, "5 produits associés"},
	{"ko-KR", 5, "1 linked product", "{%1=number of products} linked products", []interface{}{5}, "5 linked products"},
	{"zh-CN", 1, "1 non-existent string", "{%1=number of products} non-existent strings", []interface{}{1}, "1 non-existent string"},
	{"zh-CN", 5, "1 non-existent string", "{%1=number of products} non-existent strings", []interface{}{5}, "5 non-existent strings"},
}

var Ci18nTestData = []struct {
	locale      string
	ctx         string
	msg         string
	args        []interface{}
	expectedStr string
}{
	{"zh-TW", "LEGAL_CONSTANTS", "Terms of Use", nil, "使用條款"},
	{"es-LA", "SOME CONTEXT", "View tickets", nil, "View tickets"},
	{"ko-KR", "SOME CONTEXT", "Non-existent string", nil, "Non-existent string"},
}

var Cni18nTestData = []struct {
	locale      string
	ctx         string
	n           int
	msg         string
	msgPlural   string
	args        []interface{}
	expectedStr string
}{
	{"fr-FR", "", 100, "1 linked product", "{%1=number of products} linked products", []interface{}{100}, "100 produits associés"},
}
