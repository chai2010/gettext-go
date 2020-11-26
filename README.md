# goi18n

This lib is in beta. Please report any bugs or suggestions to @pkuhanandan

## Setup
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
4. Now you can use your `TranslationRepository` to translate strings. See below for detailed usage instructions.

## Translation Functions
This lib features the 4 common translation functions used at wish:
 1. **I18n**
    Translate a string
    ```
    tr.I18n("Learn more"). 
    ```
 2. **Ci18n**
    Translate a string with context. Use this when your string's meaning is ambiguous.
    ```
    tr.Ci18n("Meaning the state of an order", "state")
    tr.Ci18n("Meaning a region, like the State of New York", "state")
    ```
 3. **Ni18n**
    Translate a string with a singular and plural form.  
    ```
    tr.Ni18n(5, "Picked up 1 day ago", "Picked up  {%1=Number of days} days ago", 5)
    ```
 4. **Cni18n**
    Translate a string with context, as well as singular and plural form.
    ```
    tr.Cni18n("State as in a region, like the State of New York", 5, "Available in 1 state", "Available in  {%1=Number of States} states", 5)
    ```
    
## Placeholders
This lib only supports descriptive placeholders. It is important to provide good descriptive placeholders because many other languages have complex forms, which can change depending on what is inserted into placeholders. Providing good descriptions can give the translators more context to make accurate translations. 

**Example:**
`tr.I18n("PayPal payment scheduled for: {%1=user email}.", user.EmailAddress)`

## Usage Patterns
After setup, you can use the `TranslationRepository` to translate strings. To help you get started, consider the following two patterns for usage.

1. Use the singleton `TranslationRepository` to do the translations. In this pattern, we set the desired locale globally using `SetLocale`. Then translate using the i18n functions (similar to how it's done in clroot).
    ```
    tr, err := i.NewTranslatorRepository(i.TranslatorRepositoryOptions{LocaleDir: golangPackageLocaleDir})
    
    err = tr.SetLocale("zh-CN")

    tr.I18n("Learn more")
    tr.Ci18n("LEGAL_CONSTANTS", "Terms of Use")
    tr.Ni18n(5, "1 linked product", "{%1=number of products} linked products", 5)
    tr.Cni18n("MY CONTEXT", 1, "1 string", "{%1=number of products} strings", 1)
    ```

2. Use `GetTranslator` to get `Translator` instances for the locales of your choice from the `TranslationRepository`, and use them to translate strings independently. This is useful in multi-threaded situations, where a global locale will not suffice.
    ```
    tr, err := i.NewTranslatorRepository(i.TranslatorRepositoryOptions{LocaleDir: golangPackageLocaleDir})

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

