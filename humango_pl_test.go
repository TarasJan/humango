package humango

import "testing"

func TestIntPLThirteenThousandFiveHundredThirty(t *testing.T) {
	result, _ := WordifyWithLocale(13530, "pl")
	if result != "trzynaście tysięcy pięćset trzydzieści" {
		t.Errorf("WordifyWithLocale(13530, \"pl\") did not return \"trzynaście tysięcy pięćset trzydzieści\" : " + result)
	}
}

func TestIntPLZero(t *testing.T) {
	result, _ := WordifyWithLocale(0, "pl")
	if result != "zero" {
		t.Errorf("WordifyWithLocale(0, \"pl\") did not return \"zero\"" + result)
	}
}

func TestIntPLMinusNineteenPointFortyFour(t *testing.T) {
	result, _ := WordifyWithLocale(-19.44, "pl")
	if result != "minus dziewiętnaście , czterdzieści cztery" {
		t.Errorf("WordifyWithLocale(-19.44, \"pl\") did not return \"minus dziewiętnaście , czterdzieści cztery\"" + result)
	}
}

func TestIntPLTenMillionZlotys(t *testing.T) {
	result, _ := WordifyWithLocaleAndUnit(10000000, "pl", "PLN")
	if result != "dziesięć milionów złotych" {
		t.Errorf("WordifyWithLocaleAndUnit(10000000, \"pl\") did not return \"dziesięć milionów złotych\"" + result)
	}
}

func TestIntPLEightHundredSeventySevenThousandSixHundredTen(t *testing.T) {
	result, _ := WordifyWithLocale(877610, "pl")
	if result != "osiemset siedemdziesiąt siedem tysięcy sześćset dziesięć" {
		t.Errorf("WordifyWithLocale(877610, \"pl\") did not return \"osiemset siedemdziesiąt siedem tysięcy sześćset dziesięć\" :" + result)
	}
}

func TestIntPLThreeThousandThreeHundredThirtyThree(t *testing.T) {
	result, _ := WordifyWithLocale(3333, "pl")
	if result != "trzy tysiące trzysta trzydzieści trzy" {
		t.Errorf("WordifyWithLocale(3333, \"pl\") did not return \"trzy tysiące trzysta trzydzieści trzy\" :" + result)
	}
}

func TestIntPLTwoMillionZlotys(t *testing.T) {
	result, _ := WordifyWithLocaleAndUnit(2000000, "pl", "PLN")
	if result != "dwa miliony złotych" {
		t.Errorf("WordifyWithLocale(2000000, \"pl\", \"PLN\") did not return \"dwa miliony złotych\" :" + result)
	}
}

func TestIntPLFourBillionZlotys(t *testing.T) {
	result, _ := WordifyWithLocaleAndUnit(4000000000, "pl", "PLN")
	if result != "cztery miliardy złotych" {
		t.Errorf("WordifyWithLocale(4000000000, \"pl\", \"PLN\") did not return \"cztery miliardy złotych\" :" + result)
	}
}

func TestIntPLTwentyOneThousand(t *testing.T) {
	result, _ := WordifyWithLocale(21000, "pl")
	if result != "dwadzieścia jeden tysięcy" {
		t.Errorf("WordifyWithLocale(21000, \"pl\") did not return \"dwadzieścia jeden tysięcy\" :" + result)
	}
}

func TestIntPLTwentyThreeThousand(t *testing.T) {
	result, _ := WordifyWithLocale(23000, "pl")
	if result != "dwadzieścia trzy tysiące" {
		t.Errorf("WordifyWithLocale(23000, \"pl\") did not return \"dwadzieścia trzy tysiące\" :" + result)
	}
}
