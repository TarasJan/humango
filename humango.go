package humango

import (
	"errors"
)

func Wordify(i interface{}) (string, error) {
	return WordifyWithLocale(i, "en")
}

func WordifyWithLocale(i interface{}, lang string) (string, error) {
	return WordifyWithLocaleAndUnit(i, lang, "")
}

func WordifyWithLocaleAndUnit(i interface{}, lang string, unit string) (string, error) {
	switch i.(type) {
	case int:
		converter := NewConverter(lang, unit, false)
		return converter.wordifyInt(i.(int))
	case float64:
		converter := NewConverter(lang, unit, true)
		return converter.wordifyFloat(i.(float64))
	default:
		return "", errors.New("The Wordify method currently only handles int and float64 types")
	}
}