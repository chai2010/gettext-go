package i18n

// TranslatorManager -
type TranslatorManager struct {
	translatorRepo *TranslatorRepository
}

// NewTranslatorManager -
func NewTranslatorManager(options TranslatorRepositoryOptions) (*TranslatorManager, error) {
	translatorRepo, err := NewTranslatorRepository(options)
	if err != nil {
		return nil, err
	}
	locales := translatorRepo.GetAvailableLocales()

	// Given that the translator is designed to be initialized during container startup, it'd be better to initialize all
	// locales instead of lazying loading during runtime.
	for _, locale := range locales {
		_, _ = translatorRepo.GetTranslator(locale)
	}

	return &TranslatorManager{
		translatorRepo: translatorRepo,
	}, nil
}

// GetAvailableLocales - All the locales for which valid translation files could be found
func (t *TranslatorManager) GetAvailableLocales() []string {
	keys := make([]string, 0, len(t.translatorRepo.translatorMap))
	for k := range t.translatorRepo.translatorMap {
		keys = append(keys, k)
	}
	return keys
}

// I18n - Translate a string into the current given locale
func (t *TranslatorManager) I18n(locale string, str string, args ...interface{}) string {
	tr := t.getTranslatorByLocale(locale)
	return tr.I18n(str, args...)
}

// Ci18n - Translate a string with context into the current given locale
func (t *TranslatorManager) Ci18n(locale string, ctx, str string, args ...interface{}) string {
	tr := t.getTranslatorByLocale(locale)
	return tr.Ci18n(ctx, str, args...)
}

// Ni18n - Translate a string with plural form into the current given locale
func (t *TranslatorManager) Ni18n(locale string, n int, str, pluralStr string, args ...interface{}) string {
	tr := t.getTranslatorByLocale(locale)
	return tr.Ni18n(n, str, pluralStr, args...)
}

// Cni18n - Translate a string with plural form and context into the current given locale
func (t *TranslatorManager) Cni18n(locale string, ctx string, n int, str, pluralStr string, args ...interface{}) string {
	tr := t.getTranslatorByLocale(locale)
	return tr.Cni18n(ctx, n, str, pluralStr, args...)
}

func (t *TranslatorManager) getTranslatorByLocale(locale string) Translator {
	tr, err := t.translatorRepo.GetTranslator(locale)
	if err != nil {
		tr = t.translatorRepo.GetFallBackTranslator()
	}
	return tr
}
