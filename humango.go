package humango

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

func Wordify(i interface{}) (string, error) {
	return WordifyWithLocale(i, "en")
}

func WordifyWithLocale(i interface{}, lang string) (string, error) {
	converter := NewConverter(lang)
	switch i.(type) {
	case int:
		return converter.wordifyInt(i.(int))
	case float64:
		return converter.wordifyFloat(i.(float64))
	default:
		return "", errors.New("The Wordify method currently only handles int and float64 types")
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

	return strings.Join(wordified, c.Locale.Separator), nil
}

func (c *Converter) wordifyInt(i int) (string, error) {
	num := int(math.Abs(float64(i)))
	wordified := []string{}

	c.appendSign(i, &wordified)
	index := 0
	c.handleQuantifier(&num, &wordified, index)

	return strings.Join(wordified, c.Locale.Separator), nil
}

func (c *Converter) handleQuantifier(num *int, wordified *[]string, quantifierIndex int) {
	if quantifierIndex >= len(c.Locale.QuantifyPoints) {
		c.handleBelowHundreds(*num, wordified)
		return
	}

	quantifier := c.Locale.QuantifyPoints[quantifierIndex]
	units := *num / quantifier

	if units > 0 {
		unit_val := units
		c.handleQuantifier(&unit_val, wordified, quantifierIndex+1)
		*wordified = append(*wordified, c.Locale.Quantifiers[strconv.Itoa(quantifier)])
		*num -= units * quantifier
	}

	c.handleQuantifier(num, wordified, quantifierIndex+1)
	return
}

func (c *Converter) handleBelowHundreds(num int, wordified *[]string) {
	if !c.handledByGlyphs(num) {
		tens := num / 10
		num = num % 10
		if tens > 1 {
			*wordified = append(*wordified, c.Locale.Tens[tens])
			if num != 0 {
				*wordified = append(*wordified, c.Locale.Glyphs[num])
			}
		}
	} else {
		if !(len(*wordified) > 1 && num == 0) {
			*wordified = append(*wordified, c.Locale.Glyphs[num])
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
