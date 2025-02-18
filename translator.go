package goi18n

import "golang.org/x/text/language"

// Translator is an interface for handling translations and localization.
// It provides methods to add locales, messages, and perform translations
// with support for pluralization.
type Translator interface {
	// AddLocale adds a new locale with the given key and formatter.
	AddLocale(key string, formatter *language.Tag)

	// LoadBytes loads translation JSON content from byte slices for the specified locale.
	// If an empty locale is provided, it loads the content for the default locale.
	LoadBytes(locale string, content ...[]byte)

	// LoadFiles loads translation content from JSON files for the specified locale.
	// If an empty locale is provided, it loads the file for the default locale.
	LoadFiles(locale string, files ...string) error

	// AddMessage adds a new message with the given key and message string.
	// If an empty locale is provided, it is added to the default locale.
	AddMessage(locale, key, message string, options ...PluralOption)

	// Translate translates a message identified by the key using the provided values.
	// Returns the translated string. If an empty locale is provided or translation is not
	// found in the locale, it translates the message for the default locale. Returns
	// an empty string if the translation is not found.
	Translate(locale, key string, values map[string]any) string

	// Plural translates a plural message identified by the key based on the count
	// and provided values. If an empty locale is provided or translation is not
	// found in the locale, it translates the message for the default locale.
	// Returns an empty string if the translation is not found.
	Plural(locale, key string, count int, values map[string]any) string
}

// NewTranslator create new translator with default locale
func NewTranslator(locale string, formatter language.Tag) Translator {
	return &translator{
		locale: locale,
		tag:    formatter,
		locals: map[string]*localization{
			locale: newLocalization(formatter),
		},
	}
}
