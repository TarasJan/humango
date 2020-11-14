package humango

import "testing"

func TestIntDEEightyFiveThousandSixHundredSeventyEight(t *testing.T) {
	result, _ := WordifyWithLocale(85678, "de")
	if result != "fünfundachtzigtausendsechshundertachtundsiebzig" {
		t.Errorf("WordifyWithLocale(85678, \"de\") did not return \"fünfundachtzigtausendsechshundertachtundsiebzig\" : " + result)
	}
}

func TestIntDEEightHundredSeventySevenThousandSixHundredTen(t *testing.T) {
	result, _ := WordifyWithLocale(877610, "de")
	if result != "achthundertsiebenundsiebzigtausendsechshundertzehn" {
		t.Errorf("WordifyWithLocale(877610, \"de\") did not return \"achthundertsiebenundsiebzigtausendsechshundertzehn\" :" + result)
	}
}

func TestIntDEZero(t *testing.T) {
	result, _ := WordifyWithLocale(0, "de")
	if result != "null" {
		t.Errorf("WordifyWithLocale(0, \"de\") did not return \"null\"" + result)
	}
}

func TestIntDEMinusNineteenPointFortyFour(t *testing.T) {
	result, _ := WordifyWithLocale(-19.44, "de")
	if result != "minus neunzehn Komma vierundvierzig" {
		t.Errorf("WordifyWithLocale(-19.44, \"de\") did not return \"minus neunzehn Komma vierundvierzig\"" + result)
	}
}

func TestIntDETenMillion(t *testing.T) {
	result, _ := WordifyWithLocale(10000000, "de")
	if result != "zehn Millionen" {
		t.Errorf("WordifyWithLocale(10000000, \"de\") did not return \"zehn Millionen\"" + result)
	}
}