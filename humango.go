package humango

import (
	"errors"
	"math"
	"strings"
)

// Numbers that are not result of other compounds of numerals
var GLYPHS = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "forteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}

// Numerals for multilpes of 10
var TENS = []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

// Numerals for splitting the numbers to construct compound ones
var QUANTIFIERS = map[int]string{100: "hundred", 1000: "thousand", 1000000: "million"}

// Common points of splitting numbers with quantifiers
// NOTE for example in Japanese localization 10000 man [万] would be here
var QUANTIFY_POINTS = []int{1000000, 1000, 100}

// Sign to append at the beginning when the number is negative
const NEGATIVE_SIGN = "minus"

// Separator between each numeral in sequence
const SEPARATOR = " "

// Word signifing separation of decimal part in word
const DECIMAL_SEPARATOR_WORD = "point"

func Wordify(i interface{}) (string, error) {
	switch i.(type) {
	case int:
		return wordifyInt(i.(int))
	case float64:
		return wordifyFloat(i.(float64))
	default:
		return "", errors.New("The Wordify method currently only handles int and float64 types")
	}
}

// For now only two decimal points are supported (for monetary values usage)
func wordifyFloat(i float64) (string, error) {
	j := math.Abs(i)

	wholeNum := int(math.Floor(j))
	if math.Signbit(i) {
		wholeNum = -wholeNum
	}
	decimalsNum := int(math.Round(math.Mod(j, 1.0) * 100))

	whole, err := wordifyInt(wholeNum)
	if err != nil {
		return "", err
	}
	wordified := []string{whole}
	if decimalsNum != 0 {
		wordified = append(wordified, DECIMAL_SEPARATOR_WORD)
		decimals, err := wordifyInt(decimalsNum)
		if err != nil {
			return "", err
		}
		wordified = append(wordified, decimals)
	}

	return strings.Join(wordified, SEPARATOR), nil
}

func wordifyInt(i int) (string, error) {
	num := int(math.Abs(float64(i)))
	wordified := []string{}

	appendSign(i, &wordified)
	index := 0
	handleQuantifier(&num, &wordified, index)

	return strings.Join(wordified, SEPARATOR), nil
}

func handleQuantifier(num *int, wordified *[]string, quantifierIndex int) {
	if quantifierIndex >= len(QUANTIFY_POINTS) {
		handleBelowHundreds(*num, wordified)
		return
	}

	quantifier := QUANTIFY_POINTS[quantifierIndex]
	units := *num / quantifier

	if units > 0 {
		unit_val := units
		handleQuantifier(&unit_val, wordified, quantifierIndex+1)
		*wordified = append(*wordified, QUANTIFIERS[quantifier])
		*num -= units * quantifier
	}

	handleQuantifier(num, wordified, quantifierIndex+1)
	return
}

func handleBelowHundreds(num int, wordified *[]string) {
	if !handledByGlyphs(num) {
		tens := num / 10
		num = num % 10
		if tens > 1 {
			*wordified = append(*wordified, TENS[tens])
			if num != 0 {
				*wordified = append(*wordified, GLYPHS[num])
			}
		}
	} else {
		if !(len(*wordified) > 1 && num == 0) {
			*wordified = append(*wordified, GLYPHS[num])
		}
	}
}

func appendSign(i int, wordified *[]string) {
	if i < 0 {
		*wordified = append(*wordified, NEGATIVE_SIGN)
	}
}

func handledByGlyphs(num int) bool {
	return num <= len(GLYPHS)-1
}
