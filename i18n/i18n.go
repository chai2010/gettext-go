package i18n

import (
	"errors"
	"github.com/ContextLogic/goi18n"
	"io/ioutil"
	"sync"
)

type TranslatorRepository struct {
	mutex     				sync.Mutex
	localeDir				string
	domain					string
	translatorMap    		map[string]*Translator
	currentLocale			string
}

type TranslatorRepositoryOptions struct {
	LocaleDir	string
	Domain		*string
}

// NewTranslatorRepository - Create a new TranslatorRepository. This struct
// will contain all translators for your available locales.
func NewTranslatorRepository(options TranslatorRepositoryOptions) (*TranslatorRepository, error) {
	domain := "wish"
	if options.Domain != nil {
		domain = *options.Domain
	}

	translatorRepo := &TranslatorRepository{
		domain: domain,
		localeDir: options.LocaleDir,
	}

	err := translatorRepo.initTranslatorMap(translatorRepo.localeDir)
	if err != nil {
		return nil, err
	}
	if len(translatorRepo.GetAvailableLocales()) == 1 {
		return nil, errors.New("no valid translation files found")
	}

	return translatorRepo, nil
}

func (t *TranslatorRepository) initTranslatorMap(localeDir string) error {
	t.translatorMap = make(map[string]*Translator)

	files, err := ioutil.ReadDir(localeDir)
	if err != nil {
		return errors.New("error reading locale directory")
	}
	for _, f := range files {
		if !f.IsDir() || len(f.Name()) == 0 {
			continue
		}
		t.translatorMap[f.Name()] = nil
	}
	defaultGettexter := gettext.New(t.domain, t.localeDir).SetLanguage(t.getDefaultLocale())
	defaultTranslator := Translator{&defaultGettexter}
	t.translatorMap[t.getDefaultLocale()] = &defaultTranslator
	return nil
}

// GetAvailableLocales - All the locales for which valid translation files could be found
func (t *TranslatorRepository) GetAvailableLocales() []string {
	keys := make([]string, 0, len(t.translatorMap))
	for k := range t.translatorMap {
		keys = append(keys, k)
	}
	return keys
}

// SetLocale - Switch the locale of the TranslatorRepository. All I18n calls to the
// TranslatorRepository will use this locale. If multiple goroutines are using the same
// TranslatorRepository at the same time, use GetTranslator pattern instead.
func (t *TranslatorRepository) SetLocale(locale string) error {
	_, err := t.findOrInitTranslator(locale)
	if err != nil {
		return err
	}

	t.currentLocale = locale
	return nil
}

func (t *TranslatorRepository) findOrInitTranslator(locale string) (Translator, error) {
	if _, ok := t.translatorMap[locale]; !ok {
		return Translator{}, errors.New("locale does not exist. Use GetAvailableLocales to list valid locales")
	}

	if t.translatorMap[locale] == nil {
		newTranslator := t.createNewTranslator(locale)
		t.translatorMap[locale] = &newTranslator
	}
	return *t.translatorMap[locale], nil
}

func (t *TranslatorRepository) createNewTranslator(locale string) Translator {
	gettexter := gettext.New(t.domain, t.localeDir).SetLanguage(locale)
	return Translator{&gettexter}
}

// I18n - Translate a string into the current TranslatorRepository locale
func (t *TranslatorRepository) I18n(str string, args ...interface{}) string {
	tr := t.getCurrentTranslator()
	return tr.I18n(str, args...)
}

// Ci18n - Translate a string with context into the current TranslatorRepository locale
func (t *TranslatorRepository) Ci18n(ctx, str string, args ...interface{}) string {
	tr := t.getCurrentTranslator()
	return tr.Ci18n(ctx, str, args...)
}

// Ni18n - Translate a string with plural form into the current TranslatorRepository locale
func (t *TranslatorRepository) Ni18n(n int, str, pluralStr string, args ...interface{}) string {
	tr := t.getCurrentTranslator()
	return tr.Ni18n(n, str, pluralStr, args...)
}

// Cni18n - Translate a string with plural form and context into the current TranslatorRepository locale
func (t *TranslatorRepository) Cni18n(ctx string, n int, str, pluralStr string, args ...interface{}) string {
	tr := t.getCurrentTranslator()
	return tr.Cni18n(ctx, n, str, pluralStr, args...)
}

// GetTranslator - Get the translator for a specific locale. Can subsequently call I18n with
// the returned Translator to translate into that locale. This is useful when using one
// TranslatorRepository with multiple parallel goroutines
func (t *TranslatorRepository) GetTranslator(locale string) (Translator, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	return t.findOrInitTranslator(locale)
}

func (t *TranslatorRepository) getCurrentTranslator() Translator {
	translator, err := t.GetTranslator(t.currentLocale)
	if err != nil {
		return *t.translatorMap[t.getDefaultLocale()]
	}
	return translator
}

func (t *TranslatorRepository) getDefaultLocale() string {
	return "en"
}

// GetLocaleDir - Get the localeDir
func (t *TranslatorRepository) GetLocaleDir() string {
	return t.localeDir
}

// GetLocaleDir - Get the domain
func (t *TranslatorRepository) GetDomain() string {
	return t.domain
}

// GetLocale - Get the current locale
func (t *TranslatorRepository) GetLocale() string {
	return t.currentLocale
}