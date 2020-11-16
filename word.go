package humango

type Pluralizable interface {
	ForNumber(int, *Locale) string
}

type Word struct {
	Pluralizable
	Singular            string `json:"singular"`
	Plural              string `json:"plural"`
	SupplementaryPlural string `json:"splural"`
}

func (word *Word) ForNumber(units int, locale *Locale) string {
	result := word.Singular
	if units > 1 || units == 0 {
		result = word.Plural
		if locale.Rules["slavic_plural"] != nil && (units < 5 || (units > 19 && units%10 < 5 && units%10 > 1)) {
			result = word.SupplementaryPlural
		}
	}
	return result
}
