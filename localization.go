package goi18n

import (
	"github.com/tidwall/gjson"
	"golang.org/x/text/language"
)

// localization represents a localization context with a specific language tag,
// a collection of JSON files, and a map of messages.
type localization struct {
	tag      language.Tag
	files    []gjson.Result
	messages map[string]pluralMessage
}

// load parses the provided JSON content and appends it to the localization files
// if the content is valid JSON.
func (l *localization) load(content []byte) {
	if gjson.ValidBytes(content) {
		l.files = append(l.files, gjson.ParseBytes(content))
	}
}

// addMessage adds a plural message to the localization context.
func (l *localization) addMessage(k string, m pluralMessage) {
	l.messages[k] = m
}

func (l *localization) translate(k string, count int, values map[string]any) string {
	// translate translates a message key with the given count and values.
	if m, ok := l.messages[k]; ok {
		return m.translate(l.tag, count, values)
	}

	for _, f := range l.files {
		if v := f.Get(k); v.Exists() {
			return translateJson(l.tag, v, count, values)
		}
	}

	return ""

}

// newLocalization creates a new localization context with the given language tag.
func newLocalization(t language.Tag) *localization {
	return &localization{
		tag:      t,
		files:    make([]gjson.Result, 0),
		messages: make(map[string]pluralMessage),
	}
}
