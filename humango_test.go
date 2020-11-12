package humango

import "testing"

func TestIntZero(t *testing.T) {
	result, _ := Wordify(0)
	if result != "zero" {
		t.Errorf("Wordify(0) did not return \"zero\"")
	}
}

func TestIntFive(t *testing.T) {
	result, _ := Wordify(5)
	if result != "five" {
		t.Errorf("Wordify(5) did not return \"five\"")
	}
}

func TestIntMinusSeven(t *testing.T) {
	result, _ := Wordify(-7)
	if result != "minus seven" {
		t.Errorf("Wordify(-7) did not return \"minus seven\"")
	}
}

func TestIntEleven(t *testing.T) {
	result, _ := Wordify(11)
	if result != "eleven" {
		t.Errorf("Wordify(11) did not return \"eleven\"")
	}
}

func TestIntNineteen(t *testing.T) {
	result, _ := Wordify(19)
	if result != "nineteen" {
		t.Errorf("Wordify(19) did not return \"nineteen\"")
	}
}

func TestIntMinusTwelve(t *testing.T) {
	result, _ := Wordify(-12)
	if result != "minus twelve" {
		t.Errorf("Wordify(-12) did not return \"minus twelve\"")
	}
}

func TestIntTwentyThree(t *testing.T) {
	result, _ := Wordify(23)
	if result != "twenty three" {
		t.Errorf("Wordify(23) did not return \"twenty three\"")
	}
}

func TestIntEightySeven(t *testing.T) {
	result, _ := Wordify(87)
	if result != "eighty seven" {
		t.Errorf("Wordify(87) did not return \"eighty seven\"")
	}
}

func TestIntMinusThirty(t *testing.T) {
	result, _ := Wordify(-30)
	if result != "minus thirty" {
		t.Errorf("Wordify(-30) did not return \"minus thirty\"" + " \\ " + result)
	}
}

func TestIntNinetyNine(t *testing.T) {
	result, _ := Wordify(99)
	if result != "ninety nine" {
		t.Errorf("Wordify(99) did not return \"ninety nine\"" + " \\ " + result)
	}
}

func TestIntOneHundredOne(t *testing.T) {
	result, _ := Wordify(101)
	if result != "one hundred one" {
		t.Errorf("Wordify(101) did not return \"one hundred one\"" + " \\ " + result)
	}
}

func TestIntMinusNineHundredTwentyThree(t *testing.T) {
	result, _ := Wordify(-923)
	if result != "minus nine hundred twenty three" {
		t.Errorf("Wordify(-923) did not return \"minus nine hundred twenty three\"" + " \\ " + result)
	}
}

func TestIntOneThousand(t *testing.T) {
	result, _ := Wordify(1000)
	if result != "one thousand" {
		t.Errorf("Wordify(2000) did not return \"one thousand\"" + " \\ " + result)
	}
}

func TestIntTwoThousand(t *testing.T) {
	result, _ := Wordify(2000)
	if result != "two thousand" {
		t.Errorf("Wordify(2000) did not return \"two thousand\"" + " \\ " + result)
	}
}

func TestIntThousandNineteen(t *testing.T) {
	result, _ := Wordify(1019)
	if result != "one thousand nineteen" {
		t.Errorf("Wordify(1019) did not return \"one thousand nineteen\"" + " \\ " + result)
	}
}

func TestIntFourHundredSeventeenThousand(t *testing.T) {
	result, _ := Wordify(417000)
	if result != "four hundred seventeen thousand" {
		t.Errorf("Wordify(417000) did not return \"four hundred seventeen thousand\"" + " \\ " + result)
	}
}

func TestIntFortyThreeThousandFourHundredTwelve(t *testing.T) {
	result, _ := Wordify(43412)
	if result != "forty three thousand four hundred twelve" {
		t.Errorf("Wordify(43412) did not return \"forty three thousand four hundred twelve\"" + " \\ " + result)
	}
}

func TestIntMillion(t *testing.T) {
	result, _ := Wordify(1000000)
	if result != "one million" {
		t.Errorf("Wordify(1000000) did not return \"one million\"" + " \\ " + result)
	}
}

func TestIntMillionTwoHundredFiftyThousand(t *testing.T) {
	result, _ := Wordify(1250000)
	if result != "one million two hundred fifty thousand" {
		t.Errorf("Wordify(1250000) did not return \"one million two hundred fifty thousand\"" + " \\ " + result)
	}
}

func TestIntTwentyMillionOne(t *testing.T) {
	result, _ := Wordify(20000001)
	if result != "twenty million one" {
		t.Errorf("Wordify(20000001) did not return \"twenty million one\"" + " \\ " + result)
	}
}

func TestIntThreeHundredSeventySevenMillionFiveHundredTwentyEightThousandSevenHundredSixtyNine(t *testing.T) {
	result, _ := Wordify(377528769)
	if result != "three hundred seventy seven million five hundred twenty eight thousand seven hundred sixty nine" {
		t.Errorf("Wordify(377528769) did not return \"three hundred seventy seven million five hundred twenty eight thousand seven hundred sixty nine\"" + " \\ " + result)
	}
}

func TestIntEightySixBillionThirteen(t *testing.T) {
	result, _ := Wordify(86000000013)
	if result != "eighty six billion thirteen" {
		t.Errorf("Wordify(86000000013) did not return \"eighty six billion thirteen\"" + " \\ " + result)
	}
}

func TestFloatZero(t *testing.T) {
	result, _ := Wordify(0.0)
	if result != "zero" {
		t.Errorf("Wordify(0) did not return \"zero\"")
	}
}

func TestFloatSevenPointTwentyFour(t *testing.T) {
	result, _ := Wordify(7.24)
	if result != "seven point twenty four" {
		t.Errorf("Wordify(7.24) did not return \"seven point twenty four\"")
	}
}

func TestFloatThirtySixZeroZero(t *testing.T) {
	result, _ := Wordify(36.00)
	if result != "thirty six" {
		t.Errorf("Wordify(36.00) did not return \"thirty six\"")
	}
}

func TestFloatFortyTwoPointOne(t *testing.T) {
	result, _ := Wordify(42.01)
	if result != "forty two point one" {
		t.Errorf("Wordify(42.01) did not return \"forty two point one\"" + result)
	}
}

func TestFloatMinusNintyPointTen(t *testing.T) {
	result, _ := Wordify(-90.1)
	if result != "minus ninety point ten" {
		t.Errorf("Wordify(90.1) did not return \"ninety point ten\"" + result)
	}
}

func TestIntJPFiftyThousandSixHundredOne(t *testing.T) {
	result, _ := WordifyWithLocale(50601, "jp")
	if result != "五万六百一" {
		t.Errorf("WordifyWithLocale(50601, \"jp\") did not return \"五万六百一\"" + result)
	}
}

func TestIntJPEightHundredSeventySevenThousandSixHundredTen(t *testing.T) {
	result, _ := WordifyWithLocale(877610, "jp")
	if result != "八十七万七千六百十" {
		t.Errorf("WordifyWithLocale(877610, \"jp\") did not return \"八十七万七千六百十\"" + result)
	}
}

func TestIntJPZero(t *testing.T) {
	result, _ := WordifyWithLocale(0, "jp")
	if result != "零" {
		t.Errorf("WordifyWithLocale(0, \"jp\") did not return \"零\"" + result)
	}
}

func TestIntJPMinusNineteenPointFortyFour(t *testing.T) {
	result, _ := WordifyWithLocale(-19.44, "jp")
	if result != "-十九.四十四" {
		t.Errorf("WordifyWithLocale(-19.44, \"jp\") did not return \"-十九.四十四\"" + result)
	}
}

func TestIntJPTenMillion(t *testing.T) {
	result, _ := WordifyWithLocale(10000000, "jp")
	if result != "一千万" {
		t.Errorf("WordifyWithLocale(10000000, \"jp\") did not return \"一千万\"" + result)
	}
}

func TestOtherTypesNotSupported(t *testing.T) {
	result, err := Wordify(true)
	if err == nil || result != "" {
		t.Errorf("Wordify(true), did not return an error with empty response")
	}
}
