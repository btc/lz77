package lz

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLZ77(t *testing.T) {

	testCases := []struct {
		String      string
		Runs        []Run
		Description string
		Skip        bool
	}{
		{
			String:      "",
			Runs:        nil,
			Description: "empty string should return empty run",
		},
		{
			String: "a",
			Runs: []Run{
				{0, 0, 'a'},
			},
			Description: "single character results in single run",
		},
		{
			String: "aa",
			Runs: []Run{
				{0, 0, 'a'}, {0, 0, 'a'},
			},
			Description: "repeat 2 chars",
		},
		{
			String: "aaa",
			Runs: []Run{
				{0, 0, 'a'}, {0, 1, 'a'},
			},
			Description: "repeat 3 chars",
		},
		{
			String: "aacaacabcaba",
			Runs: []Run{
				{0, 0, 'a'}, {0, 1, 'c'}, {0, 3, 'a'}, {0, 0, 'b'}, {2, 2, 'b'}, {0, 0, 'a'},
			},
			Description: "the test case Sean & I ran through together",
		},
		{
			String:      "aXaaXaaaXY",
			Runs:        []Run{},
			Description: "TODO: fill in the runs for this greedy case",
			Skip:        true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Description+":"+testCase.String, func(t *testing.T) {
			if testCase.Skip {
				t.SkipNow()
			}

			t.Logf(`String: "%s"`, testCase.String)

			t.Run("Decode", func(t *testing.T) {
				decoded := Decode(testCase.Runs)
				require.Equal(t,
					testCase.String, string(decoded),
					"Decode: the decoded value should match the input string")
			})

			t.Run("Encode", func(t *testing.T) {
				encoded := Encode([]byte(testCase.String))
				require.Len(t,
					encoded, len(testCase.Runs),
					"Encode: should result in Runs of length equal to the testCase.Runs")
				require.Equal(t,
					testCase.Runs, encoded,
					"Encode: the encoded value should match the testCase Runs")
			})

			t.Run("Roundtrip", func(t *testing.T) {
				str := Decode(Encode([]byte(testCase.String)))
				require.Equal(t,
					testCase.String, string(str),
					"Roundtrip: Strings should come back out the same after (Encode -> Decode) operation")

				runs := Encode(Decode(testCase.Runs))
				require.Equal(t,
					testCase.Runs, runs,
					"Roundtrip: Runs should come back out the same after (Decode -> Encode) operation")
			})
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
