package goi18n

import (
	"strings"

	"golang.org/x/text/language"
)

// pluralMessage represents a message with different plural forms.
type pluralMessage struct {
	zero  string // Zero is the message when the count is zero.
	one   string // One is the message when the count is one.
	two   string // Two is the message when the count is two.
	few   string // Few is the message when the count is a few.
	many  string // Many is the message when the count is many.
	other string // Default is the default message for other counts or fallback message.
}

func (m pluralMessage) resolve(count int) string {
	var message string
	switch {
	case count == 0:
		message = m.zero
	case count == 1:
		message = m.one
	case count == 2:
		message = m.two
	case count > 2 && count <= 10:
		message = m.few
	case count > 10:
		message = m.many
	}

	if message == "" {
		message = m.other
	}

	return message
}

// translate returns the appropriate plural form of the message based on the count and locale.
func (m pluralMessage) translate(locale language.Tag, count int, values map[string]any) string {
	message := m.resolve(count)

	if message != "" {
		for k, v := range values {
			message = strings.ReplaceAll(message, "{"+k+"}", toString(locale, v))
		}
	}

	return message
}

// newMessage creates a new pluralMessage with the given default message and options.
func newMessage(message string, options ...PluralOption) pluralMessage {
	m := &pluralMessage{other: message}
	for _, opt := range options {
		opt(m)
	}

	return *m
}
