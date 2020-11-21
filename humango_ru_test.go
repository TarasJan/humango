package humango

import "testing"

func TestIntRUThirteenThousandFiveHundredThirty(t *testing.T) {
	result, _ := WordifyWithLocale(13530, "ru")
	if result != "тринадцать тысяч пятьсот тридцать" {
		t.Errorf("WordifyWithLocale(13530, \"ru\") did not return \"тринадцать тысяч пятьсот тридцать\" : " + result)
	}
}

func TestIntRUZero(t *testing.T) {
	result, _ := WordifyWithLocale(0, "ru")
	if result != "нуль" {
		t.Errorf("WordifyWithLocale(0, \"ru\") did not return \"нуль\"" + result)
	}
}

func TestIntRUMinusNineteenPointFortyFour(t *testing.T) {
	result, _ := WordifyWithLocale(-19.44, "ru")
	if result != "минус девятнадцать , сорок четыре" {
		t.Errorf("WordifyWithLocale(-19.44, \"ru\") did not return \"минус девятнадцать , сорок четыре\"" + result)
	}
}

func TestIntRUTenMillionZlotys(t *testing.T) {
	result, _ := WordifyWithLocaleAndUnit(10000000, "ru", "RUB")
	if result != "десять миллионов рублей" {
		t.Errorf("WordifyWithLocaleAndUnit(10000000, \"ru\") did not return \"десять миллионов рублей\"" + result)
	}
}

func TestIntRUEightHundredSeventySevenThousandSixHundredTen(t *testing.T) {
	result, _ := WordifyWithLocale(877610, "ru")
	if result != "восемьсот семьдесят семь тысяч шестьсот десять" {
		t.Errorf("WordifyWithLocale(877610, \"ru\") did not return \"восемьсот семьдесят семь тысяч шестьсот десять\" :" + result)
	}
}

func TestIntRUThreeThousandThreeHundredThirtyThree(t *testing.T) {
	result, _ := WordifyWithLocale(3333, "ru")
	if result != "три тысячи триста тридцать три" {
		t.Errorf("WordifyWithLocale(3333, \"ru\") did not return \"три тысячи триста тридцать три\" :" + result)
	}
}

func TestIntRUTwoMillionZlotys(t *testing.T) {
	result, _ := WordifyWithLocaleAndUnit(2000000, "ru", "RUB")
	if result != "два миллиона рублей" {
		t.Errorf("WordifyWithLocale(2000000, \"ru\", \"RUB\") did not return \"два миллиона рублей\" :" + result)
	}
}

func TestIntRUFourBillionZlotys(t *testing.T) {
	result, _ := WordifyWithLocaleAndUnit(4000000000, "ru", "RUB")
	if result != "четыре миллиарда рублей" {
		t.Errorf("WordifyWithLocale(4000000000, \"ru\", \"RUB\") did not return \"четыре миллиарда рублей\" :" + result)
	}
}

func TestIntRUTwentyOneThousand(t *testing.T) {
	result, _ := WordifyWithLocale(21000, "ru")
	if result != "двадцать один тысяч" {
		t.Errorf("WordifyWithLocale(21000, \"ru\") did not return \"двадцать один тысяч\" :" + result)
	}
}

func TestIntRUTwentyThreeThousand(t *testing.T) {
	result, _ := WordifyWithLocale(23000, "ru")
	if result != "двадцать три тысячи" {
		t.Errorf("WordifyWithLocale(23000, \"ru\") did not return \"двадцать три тысячи\" :" + result)
	}
}
