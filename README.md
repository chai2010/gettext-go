# goi18n

## Usage

1. Set up extraction and package build. See [this guide](https://wiki.wish.site/display/ENG/New+Project+Onboarding+Guide) for more information on how to set this up.
2. Save the generated translation package somewhere that can be accessed by your application at runtime (for example, in the repo itself or in a separate golang package). 
3. Create a new `TranslationRepository`. Pass the path to the `locale` directory within the translation package. 
    ```
    import (
      i "github.com/ContextLogic/goi18n/i18n"
    )

    golangPackageLocaleDir := "path/to/package/locale"
    tr, err := i.NewTranslatorRepository(i.TranslatorRepositoryOptions{LocaleDir: golangPackageLocaleDir})
    ```
    **:warning: Warning**: Depending on the size of your app, the `TranslationRepository` can get be expensive to create. For this reason, in most cases, you will want to initialize only one instance of TranslationRepository when you start your app/server and keep it available in memory. 
    
4. Now you can use the TranslationRepository to translate strings.
    ```
    err = tr.SetLocale("zh-CN")

    tr.I18n("Learn more"). 
    tr.Ci18n("LEGAL_CONSTANTS", "Terms of Use")
    tr.Ni18n(5, "1 linked product", "{%1=number of products} linked products", 5)
    tr.Cni18n("MY CONTEXT", 1, "1 string", "{%1=number of products} strings", 1)
    ```
    or alternatively in multi-threaded situations, you will want to use the following pattern instead:
    ```
    zhTranslator, err := tr.GetTranslator("zh-CN")
    frTranslator, err := tr.GetTranslator("fr-FR")
     
    var wg sync.WaitGroup
    wg.Add(2)
    go func() {
      zhTranslator.I18n("Learn more"). 
      zhTranslator.Ci18n("LEGAL_CONSTANTS", "Terms of Use")
    }
    go func() {
      frTranslator.Ni18n(5, "1 linked product", "{%1=number of products} linked products", 5)
      frTranslator.Cni18n("MY CONTEXT", 1, "1 string", "{%1=number of products} strings", 1)
    }
    wg.Wait()
    ```
    **:warning: Warning**: Translators are lazy loaded. They are only initialized when `SetLocale` or `GetTranslator` is called, NOT when the `TranslationRepository` is created.

## Full Example

```
import (
  i "github.com/ContextLogic/goi18n/i18n"
)

golangPackageLocaleDir := "path/to/package/locale"
tr, err := i.NewTranslatorRepository(i.TranslatorRepositoryOptions{LocaleDir: golangPackageLocaleDir})
err = tr.SetLocale("zh-TW")
  
tr.I18n("Learn more"). // 
tr.Ci18n("LEGAL_CONSTANTS", "Terms of Use")
```
