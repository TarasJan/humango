package humango

import "testing"

func TestUnitsIntZero(t *testing.T) {
	result, _ := WordifyWithLocaleAndUnit(0, "en", "USD")
	if result != "zero dollars" {
		t.Errorf(" WordifyWithLocaleAndUnit(0, \"en\", \"USD\") did not return \"zero dollars\"")
	}
}

func TestUnitsIntFiveHundredFiftyNineEuro(t *testing.T) {
	result, _ := WordifyWithLocaleAndUnit(559, "en", "JPY")
	if result != "five hundred fifty nine yen" {
		t.Errorf("WordifyWithLocaleAndUnit(559, \"en\", \"JPY\") did not return \"five hundred fifty nine yen\"")
	}
}

func TestUnitsIntNineBillionSixtyTwoEuro(t *testing.T) {
	result, _ := WordifyWithLocaleAndUnit(9000000062, "en", "EUR")
	if result != "nine billion sixty two euro" {
		t.Errorf("WordifyWithLocaleAndUnit(9000000062, \"en\", \"EUR\") did not return \"nine billion sixty two euro\"")
	}
}

func TestUnitsFloatSevenThousandEightyTwoPointFifteenPounds(t *testing.T) {
	result, _ := WordifyWithLocaleAndUnit(7082.15, "en", "GBP")
	if result != "seven thousand eighty two pounds fifteen pennies" {
		t.Errorf("WordifyWithLocaleAndUnit(7082.15, \"en\", \"GBP\") did not return \"seven thousand eighty two point fifteen pounds\"")
	}
}

func TestUnitsFloatOnePointFortyNinePound(t *testing.T) {
	result, _ := WordifyWithLocaleAndUnit(1.49, "en", "GBP")
	if result != "one pound forty nine pennies" {
		t.Errorf("WordifyWithLocaleAndUnit(1.49, \"en\", \"GBP\") did not return \"one pound forty nine pennies\" : " + result)
	}
}

func TestUnitsDEFloatEigthThousandEightHunndredEightyEightPointEightyEightEuro(t *testing.T) {
	result, _ := WordifyWithLocaleAndUnit(8888.88, "de", "EUR")
	if result != "achttausendachthundertachtundachtzig Euro achtundachtzig Eurocent" {
		t.Errorf("WordifyWithLocaleAndUnit(8888.88, \"de\", \"EUR\") did not return \"achttausendachthundertachtundachtzig Euro achtundachtzig Eurocent\" : " + result)
	}
}

func TestUnitsDESeventyThousandDollars(t *testing.T) {
	result, _ := WordifyWithLocaleAndUnit(70000, "de", "USD")
	if result != "siebzigtausend Dollar" {
		t.Errorf("WordifyWithLocaleAndUnit(70000, \"de\", \"USD\") did not return \"siebzigtausend Dollar\"")
	}
}

func TestUnitsJPFloatTwoHundredPointTwentyThreeDollars(t *testing.T) {
	result, _ := WordifyWithLocaleAndUnit(200.23, "jp", "USD")
	if result != "二百ドル二十三セント" {
		t.Errorf("WordifyWithLocaleAndUnit(200.23, \"jp\", \"USD\") did not return \"二百ドル二十三セント\"" + result)
	}
}

func TestUnitsPLFloatThirteenThousandFiveHundredThirtyZlotyFiftyGroszys(t *testing.T) {
	result, _ := WordifyWithLocaleAndUnit(13530.50, "pl", "PLN")
	if result != "trzynaście tysięcy pięćset trzydzieści złotych pięćdziesiąt groszy" {
		t.Errorf("WordifyWithLocaleAndUnit(13530.50, \"pl\", \"PLN\") did not return \"trzynaście tysięcy pięćset trzydzieści złotych pięćdziesiąt groszy\"" + result)
	}
}

func TestUnitsPLFloatTwoZlotyTwentyThreeGroszys(t *testing.T) {
	result, _ := WordifyWithLocaleAndUnit(2.23, "pl", "PLN")
	if result != "dwa złote dwadzieścia trzy grosze" {
		t.Errorf("WordifyWithLocaleAndUnit(2.23, \"pl\", \"PLN\") did not return \"dwa złote dwadzieścia trzy grosze\"" + result)
	}
}

func TestUnitsPLFloatFourPoundsThreePennies(t *testing.T) {
	result, _ := WordifyWithLocaleAndUnit(4.03, "pl", "GBP")
	if result != "cztery funty trzy pensy" {
		t.Errorf("WordifyWithLocaleAndUnit(2.23, \"pl\", \"GBP\") did not return \"dwa złote dwadzieścia trzy grosze\"" + result)
	}
}

func TestUnitsJPIntThreeHundredThousandSevenHundredYen(t *testing.T) {
	result, _ := WordifyWithLocaleAndUnit(300700, "jp", "JPY")
	if result != "三十万七百円" {
		t.Errorf("WordifyWithLocaleAndUnit(300700, \"jp\", \"JPY\") did not return \"三十万七百円\"")
	}
}
