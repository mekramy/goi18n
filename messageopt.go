package goi18n

// PluralOption is a type for functions that modify a Message with plural forms.
type PluralOption func(msg *pluralMessage)

// PluralZero sets the zero plural form of the message.
func PluralZero(msg string) PluralOption {
	return func(m *pluralMessage) {
		m.zero = msg
	}
}

// PluralOne sets the one plural form of the message.
func PluralOne(msg string) PluralOption {
	return func(m *pluralMessage) {
		m.one = msg
	}
}

// PluralTwo sets the two plural form of the message.
func PluralTwo(msg string) PluralOption {
	return func(m *pluralMessage) {
		m.two = msg
	}
}

// PluralFew sets the few plural form of the message.
func PluralFew(msg string) PluralOption {
	return func(m *pluralMessage) {
		m.few = msg
	}
}

// PluralMany sets the many plural form of the message.
func PluralMany(msg string) PluralOption {
	return func(m *pluralMessage) {
		m.many = msg
	}
}
