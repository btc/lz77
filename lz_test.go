package lz

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimpleCases(t *testing.T) {

	testCases := []struct {
		String      string
		Runs        []Run
		Description string
		Skip        bool
	}{
		{
			String:      "",
			Runs:        nil,
			Description: "empty string should return empty run?", // TODO confirm
		},
		{
			String:      "",
			Runs:        []Run{},
			Description: "nil Runs should be treated the same as empty []Run value",
		},
		{
			String: "a",
			Runs: []Run{
				{0, 0, 'a'},
			},
			Description: "single character results in single run",
		},
		{
			String: "aacaacabcaba",
			Runs: []Run{
				{0, 0, 'a'}, {0, 1, 'c'}, {0, 3, 'a'}, {0, 0, 'b'}, {2, 2, 'b'}, {0, 0, 'a'},
			},
			Description: "the test case Sean & I ran through together",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {

			t.Logf(`String: "%s"`, testCase.String)

			decoded := Decode(testCase.Runs)
			require.Equal(t,
				testCase.String, string(decoded),
				"Decode: the decoded runs should match the input string")

			return

			encoded := Encode([]byte(testCase.String))
			require.Equal(t,
				testCase.Runs, encoded,
				"Encode: the encoded String should match the testCase Runs")

			str := Decode(Encode([]byte(testCase.String)))
			require.Equal(t,
				testCase.String, string(str),
				"Roundtrip: Strings should come back out the same after (Encode -> Decode) operation")

			runs := Encode(Decode(testCase.Runs))
			require.Equal(t,
				testCase.Runs, runs,
				"Roundtrip: Runs should come back out the same after (Decode -> Encode) operation")
		})
	}
}

func TestSliceIndexing(t *testing.T) {
	require.Equal(t,
		""[0:0], "")
	/*
		NB: rules for indexing
		negative numbers aren’t allowed... OK
		low ≤ high... OK
		Offset + Len ≤ len(input)... yes because defined as index into input

	*/
}
