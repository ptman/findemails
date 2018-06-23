package findemails

import (
	"reflect"
	"testing"
)

func TestFindEmails(t *testing.T) {
	testCases := []struct {
		src    string
		emails []string
	}{
		{
			"a.single@email.com",
			[]string{
				"a.single@email.com",
			},
		},
		{
			"a@foo.com b@foo.com",
			[]string{
				"a@foo.com",
				"b@foo.com",
			},
		},
		{
			"filler a@foo.com,moreuselessstuff",
			[]string{
				"a@foo.com",
			},
		},
		{
			"a@foo.com,b@foo.com",
			[]string{
				"a@foo.com",
				"b@foo.com",
			},
		},
		{
			"a@foo.com;b@foo.com",
			[]string{
				"a@foo.com",
				"b@foo.com",
			},
		},
		{
			"@example:matrix.org",
			[]string{},
		},
		/* TODO: ?
		{
			"a@foo.com-b@foo.com",
			[]string{},
		},
		*/
	}

	for _, tc := range testCases {
		t.Run(tc.src, func(t *testing.T) {
			emails := FindEmails(tc.src)
			if !((len(emails) == 0 && len(tc.emails) == 0) ||
				reflect.DeepEqual(emails, tc.emails)) {
				t.Errorf("FindEmails(%v) = %v != %v",
					tc.src, emails, tc.emails)
			}
		})
	}
}
