package ftag

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParser_Parse(t *testing.T) {
	cases := []struct {
		in     string
		expect Tag
		err    error
	}{
		{
			in:     ``,
			expect: Tag{},
		},
		{
			in: `example`,
			expect: Tag{
				Name: "example",
			},
		},
		{
			in: `,aliases=foo bar`,
			expect: Tag{
				Aliases: []string{"foo", "bar"},
			},
		},
		{
			in: `,aliases= foo bar`,
			expect: Tag{
				Aliases: []string{"foo", "bar"},
			},
		},
		{
			in: `,short=e,desc=desc`,
			expect: Tag{
				Shorthand:   "e",
				Description: "desc",
			},
		},
		{
			in:  `,short==,desc=desc`,
			err: errors.New(`got invalid tag value: "short=="`),
		},
		{
			in:  `,short=ee`,
			err: errors.New(`got invalid tag value: key 'short' must have only 1 character: "short=ee"`),
		},
		{
			in:  `,foo=bar`,
			err: errors.New(`got invalid tag key: "foo=bar"`),
		},
	}

	for _, tc := range cases {
		tag, err := Parse(tc.in)
		require.Equal(t, tc.err, err, tc.in)
		if err == nil {
			require.EqualValues(t, tc.expect, tag, tc.in)
		}
	}
}
