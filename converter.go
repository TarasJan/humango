package humango

import (
	"math"
	"strconv"
	"strings"
)

type Converter struct {
	Locale      *Locale
	Unit        string
	WithDecimal bool
}

func NewConverter(language string, unit string, withDecimal bool) *Converter {
	Locale := localeFor(language)
	return &Converter{
		Locale:      Locale,
		Unit:        unit,
		WithDecimal: withDecimal,
	}
}

// For now only two decimal points are supported (for monetary values usage)
func (c *Converter) wordifyFloat(i float64) (string, error) {
	j := math.Abs(i)

	wholeNum := int(math.Floor(j))
	if math.Signbit(i) {
		wholeNum = -wholeNum
	}
	decimalsNum := int(math.Round(math.Mod(j, 1.0) * 100))

	whole, err := c.wordifyInt(wholeNum)
	if err != nil {
		return "", err
	}
	wordified := []string{whole}
	if decimalsNum != 0 {
		wordified = append(wordified, c.Locale.DecimalSeparatorWord)
		decimals, err := c.wordifyInt(decimalsNum)
		if err != nil {
			return "", err
		}
		wordified = append(wordified, decimals)
	}

	if c.WithDecimal {
		c.appendUnit(j, &wordified)
	}

	return strings.Join(wordified, c.Locale.Separator), nil
}

func (c *Converter) wordifyInt(i int) (string, error) {
	num := int(math.Abs(float64(i)))
	wordified := []string{}
	index := 0

	c.appendSign(i, &wordified)
	c.handleQuantifier(&num, &wordified, index)

	if !c.WithDecimal {
		c.appendUnit(math.Abs(float64(i)), &wordified)
	}

	result := strings.Join(wordified, c.Locale.Separator)
	return strings.TrimSpace(result), nil
}

func (c *Converter) handleQuantifier(num *int, wordified *[]string, quantifierIndex int) {
	if quantifierIndex >= len(c.Locale.QuantifyPoints) {
		c.handleBelowHundreds(*num, wordified)
		return
	}
	var quantifierName string
	quantifier := c.Locale.QuantifyPoints[quantifierIndex]
	units := *num / quantifier

	if units > 0 {
		unit_val := units
		if c.Locale.Rules["custom_hundreds"] != nil && quantifier == 100 {
			quantifierName = c.Locale.Rules["custom_hundreds"].Context["hundreds"].([]interface{})[units].(string)
		} else {
			quantifierName = c.Locale.Quantifiers[strconv.Itoa(quantifier)].Singular
			if units > 1 {
				quantifierName = c.Locale.Quantifiers[strconv.Itoa(quantifier)].Plural
				if c.Locale.Rules["slavic_plural"] != nil && (units < 5 || (units > 19 && units%10 < 5 && units%10 > 1)) {
					quantifierName = c.Locale.Rules["slavic_plural"].Context["quantifiers"].(map[string]interface{})[strconv.Itoa(quantifier)].(string)
				}
			}
			c.handleQuantifier(&unit_val, wordified, quantifierIndex+1)
		}

		*wordified = append(*wordified, quantifierName)
		*num -= units * quantifier
	}

	c.handleQuantifier(num, wordified, quantifierIndex+1)
	return
}

func (c *Converter) handleBelowHundreds(num int, wordified *[]string) {
	if !c.handledByGlyphs(num) {
		subNumber := []string{}
		tens := num / 10
		num = num % 10
		if tens > 1 {
			subNumber = append(subNumber, c.Locale.Tens[tens])
			if num != 0 {
				subNumber = append(subNumber, c.Locale.Glyphs[num])
			}
			if c.Locale.Rules["agglunative_tens"] != nil {
				joinerWord := c.Locale.Rules["agglunative_tens"].Context["joiner_word"]
				for i := len(subNumber)/2 - 1; i >= 0; i-- {
					opp := len(subNumber) - 1 - i
					subNumber[i], subNumber[opp] = subNumber[opp], subNumber[i]
				}
				subNumber = []string{strings.Join(subNumber, joinerWord.(string))}
			}
			*wordified = append(*wordified, subNumber...)
		}
	} else {
		if !(len(*wordified) > 1 && num == 0) {
			*wordified = append(*wordified, c.Locale.Glyphs[num])
		}
	}
}

func (c *Converter) appendUnit(i float64, wordified *[]string) {
	if c.Locale.Units[c.Unit] != nil {
		unit := c.Locale.Units[c.Unit]
		if i >= 2.0 || i == 0.0 {
			*wordified = append(*wordified, unit.Plural)
		} else {
			*wordified = append(*wordified, unit.Singular)
		}
	}
}

func (c *Converter) appendSign(i int, wordified *[]string) {
	if i < 0 {
		*wordified = append(*wordified, c.Locale.NegativeSign)
	}
}

func (c *Converter) handledByGlyphs(num int) bool {
	return num <= len(c.Locale.Glyphs)-1
}
