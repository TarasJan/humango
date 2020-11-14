package humango

import "testing"

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