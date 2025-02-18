package goi18n_test

import (
	"testing"

	"github.com/mekramy/goi18n"
	"golang.org/x/text/language"
)

func TestTranslator(t *testing.T) {
	// Raw data
	en := `{
		"welcome": "Hello {name}, Welcome!",
		"notification": {
			"zero": "No new message",
			"one": "You have a new message",
			"other": "You have {count} new messages"
		}
	}`
	fa := `{
		"welcome": "سلام {name}, خوش آمدید!",
		"notification": {
			"zero": "پیام جدیدی ندارید",
			"one": "یک پیام جدید دارید",
			"other": "{count} پیام جدید دارید"
		}
	}`

	// Create translator
	tr := goi18n.NewTranslator("en", language.English)
	tr.AddLocale("fa", &language.Persian)
	tr.LoadBytes("en", []byte(en))
	tr.LoadBytes("fa", []byte(fa))

	// Run tests
	t.Run("Simple", func(t *testing.T) {
		res := tr.Translate("", "welcome", map[string]any{"name": "John doe"})
		if res != "Hello John doe, Welcome!" {
			t.Fatalf("fails with %s\n", res)
		}
	})

	t.Run("Plural", func(t *testing.T) {
		res := tr.Plural("fa", "notification", 0, map[string]any{"name": "John doe"})
		if res != "پیام جدیدی ندارید" {
			t.Fatalf("fails with %s\n", res)
		}

		res = tr.Plural("ar", "notification", 10, map[string]any{"count": 10})
		if res != "You have 10 new messages" {
			t.Fatalf("fails with %s\n", res)
		}
	})
}
