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
	if result != "seven thousand eighty two point fifteen pounds" {
		t.Errorf("WordifyWithLocaleAndUnit(7082.15, \"en\", \"GBP\") did not return \"seven thousand eighty two point fifteen pounds\"")
	}
}

func TestUnitsFloatOnePointFortyNinePound(t *testing.T) {
	result, _ := WordifyWithLocaleAndUnit(1.49, "en", "GBP")
	if result != "one point forty nine pound" {
		t.Errorf("WordifyWithLocaleAndUnit(1.49, \"en\", \"GBP\") did not return \"one point forty nine pound\"")
	}
}


func TestUnitsDEFloatEigthThousandEightHunndredEightyEightPointEightyEightEuro(t *testing.T) {
	result, _ := WordifyWithLocaleAndUnit(8888.88, "de", "EUR")
	if result != "achttausendachthundertachtundachtzig Komma achtundachtzig Euro" {
		t.Errorf("WordifyWithLocaleAndUnit(8888.88, \"de\", \"EUR\") did not return \"achttausendachthundertachtundachtzig Komma achtundachtzig Euro\"")
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
	if result != "二百.二十三ドル" {
		t.Errorf("WordifyWithLocaleAndUnit(200.23, \"jp\", \"USD\") did not return \"二百.二十三ドル\"" + result)
	}
}

func TestUnitsJPIntThreeHundredThousandSevenHundredYen(t *testing.T) {
	result, _ := WordifyWithLocaleAndUnit(300700, "jp", "JPY")
	if result != "三十万七百円" {
		t.Errorf("WordifyWithLocaleAndUnit(300700, \"jp\", \"JPY\") did not return \"三十万七百円\"")
	}
}