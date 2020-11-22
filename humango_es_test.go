package humango

import "testing"

func TestIntESThirteenThousandFiveHundredThirty(t *testing.T) {
	result, _ := WordifyWithLocale(13530, "es")
	if result != "trece mil quinientos trenta" {
		t.Errorf("WordifyWithLocale(13530, \"es\") did not return \"trece mil quinientos trenta\" : " + result)
	}
}

func TestIntESZero(t *testing.T) {
	result, _ := WordifyWithLocale(0, "es")
	if result != "cero" {
		t.Errorf("WordifyWithLocale(0, \"es\") did not return \"cero\"" + result)
	}
}

func TestIntESMinusNineteenPointFortyFour(t *testing.T) {
	result, _ := WordifyWithLocale(-19.44, "es")
	if result != "minus diecinueve , quarenta y cuarto" {
		t.Errorf("WordifyWithLocale(-19.44, \"es\") did not return \"minus diecinueve , quarenta y cuarto\"" + result)
	}
}

func TestIntESTenMillionZlotys(t *testing.T) {
	result, _ := WordifyWithLocaleAndUnit(10000000, "es", "MXN")
	if result != "diez millónes pesos" {
		t.Errorf("WordifyWithLocaleAndUnit(10000000, \"es\") did not return \"diez millónes pesos\"" + result)
	}
}

func TestIntESEightHundredSeventySevenThousandSixHundredTen(t *testing.T) {
	result, _ := WordifyWithLocale(877610, "es")
	if result != "ochocientos setenta y siete mil seiscientos diez" {
		t.Errorf("WordifyWithLocale(877610, \"es\") did not return \"ochocientos setenta y siete mil seiscientos diez\" :" + result)
	}
}

func TestIntESThreeThousandThreeHundredThirtyThree(t *testing.T) {
	result, _ := WordifyWithLocale(3333, "es")
	if result != "tres mil trescientos trenta y tres" {
		t.Errorf("WordifyWithLocale(3333, \"es\") did not return \"tres mil trescientos trenta y tres\" :" + result)
	}
}

func TestIntESTwoMillionZlotys(t *testing.T) {
	result, _ := WordifyWithLocaleAndUnit(2000000, "es", "MXN")
	if result != "dos millónes pesos" {
		t.Errorf("WordifyWithLocale(2000000, \"es\", \"MXN\") did not return \"dos millónes pesos\" :" + result)
	}
}

func TestIntESFourBillionZlotys(t *testing.T) {
	result, _ := WordifyWithLocaleAndUnit(4000000000, "es", "MXN")
	if result != "cuarto millardos pesos" {
		t.Errorf("WordifyWithLocale(4000000000, \"es\", \"MXN\") did not return \":cuarto millardos pesos\" :" + result)
	}
}

func TestIntESTwentyOneThousand(t *testing.T) {
	result, _ := WordifyWithLocale(21000, "es")
	if result != "veintiuno mil" {
		t.Errorf("WordifyWithLocale(21000, \"es\") did not return \"veintiuno mil\" :" + result)
	}
}

func TestIntESTwentyThreeThousand(t *testing.T) {
	result, _ := WordifyWithLocale(23000, "es")
	if result != "veintitrés mil" {
		t.Errorf("WordifyWithLocale(23000, \"es\") did not return \"veintitrés mil\" :" + result)
	}
}
