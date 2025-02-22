# goi18n

`goi18n` is a Go package for handling translations and localization with support for pluralization. It allows you to manage multiple locales, load translation files, and perform translations based on message keys and values.

## Installation

To install `goi18n`, use `go get`:

```sh
go get github.com/mekramy/goi18n
```

## Usage

### Creating a Translator

To create a new translator, use the `NewTranslator` function. You need to provide a default locale and a language tag for formatting numbers.

```go
package main

import (
    "github.com/mekramy/goi18n"
    "golang.org/x/text/language"
)

func main() {
    tr := goi18n.NewTranslator("en", language.English)
}
```

### Adding Locales

You can add additional locales using the `AddLocale` method.

```go
tr.AddLocale("fa", &language.Persian)
```

### Loading Translation Files

You can load translation files from byte slices or from files.

#### Loading from Byte Slices

```go
en := []byte(`{
    "welcome": "Hello {name}, Welcome!",
    "notification": {
        "zero": "No new message",
        "one": "You have a new message",
        "other": "You have {count} new messages"
    }
}`)
tr.LoadBytes("en", en)
```

#### Loading from Files

```go
err := tr.LoadFiles("en", "path/to/en.json")
if err != nil {
    log.Fatal(err)
}
```

### Adding Messages

You can add individual messages using the `AddMessage` method.

```go
tr.AddMessage("en", "greeting", "Hello {name}!", goi18n.PluralOne("Hello {name}!"))
```

### Translating Messages

To translate a message, use the `Translate` method.

```go
res := tr.Translate("en", "welcome", map[string]any{"name": "John"})
fmt.Println(res) // Output: Hello John, Welcome!
```

### Pluralization

To handle pluralization, use the `Plural` method.

```go
res := tr.Plural("en", "notification", 5, map[string]any{"count": 5})
fmt.Println(res) // Output: You have 5 new messages
```

## Example

Here is a complete example demonstrating the usage of `goi18n`.

```go
package main

import (
    "fmt"
    "log"

    "github.com/mekramy/goi18n"
    "golang.org/x/text/language"
)

func main() {
    // Create translator
    tr := goi18n.NewTranslator("en", language.English)
    tr.AddLocale("fa", &language.Persian)

    // Load translations
    en := []byte(`{
        "welcome": "Hello {name}, Welcome!",
        "notification": {
            "zero": "No new message",
            "one": "You have a new message",
            "other": "You have {count} new messages"
        }
    }`)
    fa := []byte(`{
        "welcome": "سلام {name}, خوش آمدید!",
        "notification": {
            "zero": "پیام جدیدی ندارید",
            "one": "یک پیام جدید دارید",
            "other": "{count} پیام جدید دارید"
        }
    }`)
    tr.LoadBytes("en", en)
    tr.LoadBytes("fa", fa)

    // Translate messages
    res := tr.Translate("en", "welcome", map[string]any{"name": "John"})
    fmt.Println(res) // Output: Hello John, Welcome!

    res = tr.Plural("fa", "notification", 0, map[string]any{"count": 0})
    fmt.Println(res) // Output: پیام جدیدی ندارید
}
```

## License

This project is licensed under the ISC License. See the [LICENSE](LICENSE) file for details.
