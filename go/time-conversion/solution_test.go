package timeconversion

import "testing"

func TestTimeConversionPM(t *testing.T) {
	converted := timeConversion("12:01:00PM")

	if converted != "12:01:00" {
		t.Errorf("Expected 12:01:00, got %s", converted)
	}
}

func TestTimeConversionAM(t *testing.T) {
	converted := timeConversion("12:01:00AM")

	if converted != "00:01:00" {
		t.Errorf("Expected 00:01:00, got %s", converted)
	}
}

func TestTimeConversionPM2(t *testing.T) {
	converted := timeConversion("01:01:00PM")

	if converted != "13:01:00" {
		t.Errorf("Expected 13:01:00, got %s", converted)
	}
}
