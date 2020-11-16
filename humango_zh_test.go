package humango

import "testing"

func TestIntZHFiftyThousandSixHundredOne(t *testing.T) {
	result, _ := WordifyWithLocale(50601, "zh")
	if result != "五万六百一" {
		t.Errorf("WordifyWithLocale(50601, \"zh\") did not return \"五万六百一\"" + result)
	}
}

func TestIntZHEightHundredSeventySevenThousandSixHundredTen(t *testing.T) {
	result, _ := WordifyWithLocale(877610, "zh")
	if result != "八十七万七千六百十" {
		t.Errorf("WordifyWithLocale(877610, \"zh\") did not return \"八十七万七千六百十\"" + result)
	}
}

func TestIntZHZero(t *testing.T) {
	result, _ := WordifyWithLocale(0, "zh")
	if result != "零" {
		t.Errorf("WordifyWithLocale(0, \"zh\") did not return \"零\"" + result)
	}
}

func TestIntZHMinusNineteenPointFortyFour(t *testing.T) {
	result, _ := WordifyWithLocale(-19.44, "zh")
	if result != "-十九.四十四" {
		t.Errorf("WordifyWithLocale(-19.44, \"zh\") did not return \"-十九.四十四\"" + result)
	}
}

func TestIntZHTenMillion(t *testing.T) {
	result, _ := WordifyWithLocale(10000000, "zh")
	if result != "一千万" {
		t.Errorf("WordifyWithLocale(10000000, \"zh\") did not return \"一千万\"" + result)
	}
}
