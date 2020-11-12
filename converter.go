package humango

import (
	"encoding/json"
)

type Locale struct {
	// Separator between each numeral in sequence
	Separator string `json:"separator"`
	// Sign to append at the beginning when the number is negative
	NegativeSign string `json:"negative_sign"`
	// Common points of splitting numbers with quantifiers
	// NOTE for example in Japanese localization 10000 man [万] would be here
	QuantifyPoints []int `json:"quantify_points"`
	// Numerals for splitting the numbers to construct compound ones
	Quantifiers map[string]string `json:"quantifiers"`
	// Numerals for multilpes of 10
	Tens []string `json:"tens"`
	// Numbers that are not result of other compounds of numerals
	Glyphs []string `json:"glyphs"`
	// Word signifing separation of decimal part in number
	DecimalSeparatorWord string `json:"decimal_separator_word"`
}

type Converter struct {
	Locale *Locale
}

func NewConverter(language string) *Converter {
	Locale := localeFor(language)
	return &Converter{
		Locale: Locale,
	}
}

func localeFor(language string) *Locale {
	config := readLocaleConfig()
	return config[language]
}

func readLocaleConfig() map[string]*Locale {
	var localeMap map[string]*Locale
	// config, err := ioutil.ReadFile("./config.json")

	// if err != nil {
	// }

	json.Unmarshal([]byte(CONFIG), &localeMap)
	return localeMap
}
