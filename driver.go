package goi18n

import (
	"os"
	"strings"
	"sync"

	"golang.org/x/text/language"
)

type translator struct {
	locale string
	tag    language.Tag
	locals map[string]*localization
	mutex  sync.RWMutex
}

func (t *translator) resolveLocale(l string) (string, bool) {
	if l == "" {
		l = t.locale
	}

	locale, ok := t.locals[l]
	return l, ok && locale != nil
}

func (t *translator) AddLocale(k string, f *language.Tag) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if f == nil {
		t.locals[k] = newLocalization(t.tag)
	} else {
		t.locals[k] = newLocalization(*f)
	}
}

func (t *translator) LoadBytes(l string, contents ...[]byte) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	locale, ok := t.resolveLocale(l)
	if !ok {
		return
	}

	for _, c := range contents {
		t.locals[locale].load(c)
	}
}

func (t *translator) LoadFiles(l string, files ...string) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	locale, ok := t.resolveLocale(l)
	if !ok {
		return nil
	}

	for _, f := range files {
		content, err := os.ReadFile(f)
		if err != nil {
			return err
		}
		t.locals[locale].load(content)
	}
	return nil
}
func (t *translator) AddMessage(l, k, m string, options ...PluralOption) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	m = strings.TrimSpace(m)
	locale, ok := t.resolveLocale(l)
	if !ok || m == "" {
		return
	}

	t.locals[locale].addMessage(k, newMessage(m, options...))
}

func (t *translator) Translate(l, k string, values map[string]any) string {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	locale, ok := t.resolveLocale(l)
	if !ok {
		locale, ok = t.resolveLocale("")
		if !ok {
			return ""
		}
	}

	res := t.locals[locale].translate(k, -1, values)
	if res == "" && locale != t.locale {
		res = t.locals[t.locale].translate(k, -1, values)
	}

	return res
}

func (t *translator) Plural(l, k string, c int, values map[string]any) string {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	locale, ok := t.resolveLocale(l)
	if !ok {
		locale, ok = t.resolveLocale("")
		if !ok {
			return ""
		}
	}

	res := t.locals[locale].translate(k, c, values)
	if res == "" && locale != t.locale {
		res = t.locals[t.locale].translate(k, c, values)
	}

	return res
}
