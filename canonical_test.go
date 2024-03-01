package hexpalette

import (
	"testing"
)

func TestCanonical(t *testing.T) {

	tests := []struct{
		HexColorCodes string
		Expected string
	}{
		{
			HexColorCodes: "",
			Expected:       "",
		},



		{
			HexColorCodes: "#12345",
			Expected:       "12345",
		},
		{
			HexColorCodes: "#6789a",
			Expected:       "6789a",
		},
		{
			HexColorCodes: "#6789A",
			Expected:       "6789a",
		},
		{
			HexColorCodes: "#cdef01",
			Expected:       "cdef01",
		},
		{
			HexColorCodes: "#CDEF01",
			Expected:       "cdef01",
		},



		{
			HexColorCodes: "CDEF01",
			Expected:      "cdef01",
		},



		{
			HexColorCodes: " CDEF01",
			Expected:       "cdef01",
		},
		{
			HexColorCodes: "CDEF01 ",
			Expected:      "cdef01",
		},
		{
			HexColorCodes: " CDEF01 ",
			Expected:       "cdef01",
		},
		{
			HexColorCodes: "    CDEF01   ",
			Expected:          "cdef01",
		},



		{
			HexColorCodes: "#e63946 #edae49 #3376BD #00798C #52489c ",
			Expected:       "e63946edae493376bd00798c52489c",
		},



		{
			HexColorCodes: " #e63946 edae49 #3376BD   #00798C     #52489c ",
			Expected:        "e63946edae493376bd00798c52489c",
		},
	}

	for testNumber, test := range tests {
		actual, err := Canonical(test.HexColorCodes)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one..", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("COLORS: %q", test.HexColorCodes)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual 'palette-string' is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("COLORS: %q", test.HexColorCodes)
			continue
		}
	}
}

