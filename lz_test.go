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
	}{
		{
			"",
			nil,                                     // TODO matters whether nil or empty slice?
			"empty string should return empty run?", // TODO confirm
		},
		{
			"a",
			[]Run{
				{0, 0, 'a'},
			},
			"single character results in single run",
		},
		{
			"aacaacabcaba",
			[]Run{
				{0, 0, 'a'},
				{0, 1, 'c'}, {0, 3, 'a'}, {0, 0, 'b'}, {2, 2, 'b'}, {0, 0, 'a'},
			},
			"the test case Sean & I ran through together",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {

			decoded := Decode(testCase.Runs)
			require.Equal(t,
				testCase.String, string(decoded),
				"the decoded runs should match the input string")

			encoded := Encode([]byte(testCase.String))
			require.Equal(t,
				testCase.Runs, encoded,
				"the encoded String should match the testCase Runs")

			str := Decode(Encode([]byte(testCase.String)))
			require.Equal(t,
				testCase.String, string(str),
				"Strings should come back out the same after (Encode -> Decode) operation")

			runs := Encode(Decode(testCase.Runs))
			require.Equal(t,
				testCase.Runs, runs,
				"Runs should come back out the same after (Decode -> Encode) operation")
		})
	}
}
