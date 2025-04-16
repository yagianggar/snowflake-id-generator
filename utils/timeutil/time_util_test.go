package timeutil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	testutil "github.com/Electrum-id/electrum-transaction-service/pkg/utils/testutils"
)

func TestTimeToISO8601(t *testing.T) {
	var nilTime time.Time
	tests := []struct {
		name     string
		date     time.Time
		expected string
	}{
		{
			name:     "happy case",
			date:     testutil.MustTime("2020-11-24T16:07:21Z"),
			expected: "2020-11-24T16:07:21Z",
		},
		{
			name:     "zero time",
			date:     time.Time{},
			expected: "",
		},
		{
			name:     "zero time",
			date:     nilTime,
			expected: "",
		},
	}

	for _, test := range tests {
		require.Equal(t, test.expected, TimeToISO8601(test.date))
	}
}

func TestParseISO8601(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected time.Time
		err      bool
	}{
		{
			name:     "happy case",
			value:    "2020-11-24T16:07:21Z",
			expected: testutil.MustTime("2020-11-24T16:07:21Z"),
			err:      false,
		},
		{
			name:     "empty",
			value:    "",
			expected: time.Time{},
			err:      true,
		},
	}

	for _, test := range tests {
		date, err := ParseISO8601(test.value)
		if test.err {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
			require.Equal(t, test.expected, date)
		}
	}
}
