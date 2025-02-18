package goi18n

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/tidwall/gjson"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// toString formats the given value `v` according to the specified language `l`.
// It handles different types of values:
// - If `v` implements the fmt.Stringer interface, it returns the result of v.String().
// - If `v` is an integer type, it formats it as a decimal number.
// - If `v` is a floating-point type, it formats it with two decimal places.
// - For all other types, it uses fmt.Sprintf to format the value.
func toString(l language.Tag, v any) string {
	// Dereference pointer if val is a pointer
	if val := reflect.ValueOf(v); val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return ""
		} else {
			v = val.Elem().Interface()
		}
	}

	// Stringify value
	switch v := v.(type) {
	case string:
		return v
	case fmt.Stringer:
		return v.String()
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return message.NewPrinter(l).Sprintf("%d", v)
	case float32, float64:
		return message.NewPrinter(l).Sprintf("%.2f", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// translateJson translates a JSON result based on the given locale and count.
// It selects the appropriate translation string based on the count and replaces placeholders with values.
func translateJson(locale language.Tag, json gjson.Result, count int, values map[string]any) string {
	// Get message based on count
	var message string
	switch {
	case count == 0:
		message = json.Get("zero").Str
	case count == 1:
		message = json.Get("one").Str
	case count == 2:
		message = json.Get("two").Str
	case count > 2 && count <= 10:
		message = json.Get("few").Str
	default:
		message = json.Get("many").Str
	}

	// Fallback to default or whole json string if message is empty
	if message == "" {
		message = json.Get("other").Str
		if message == "" {
			message = json.Str
		}
	}

	// Replace placeholders in message with values
	if message != "" {
		for k, v := range values {
			message = strings.ReplaceAll(message, "{"+k+"}", toString(locale, v))
		}
	}

	return message
}
