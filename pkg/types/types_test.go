package types

import "testing"

func TestStringSet(t *testing.T) {
	cases := []struct {
		expected string
		inputs   []string
	}{
		{
			expected: `["aa"]`,
			inputs:   []string{"aa"},
		},
		{
			expected: `["aa","bb"]`,
			inputs:   []string{"aa", "bb"},
		},
		{
			expected: `["aa","bb"]`,
			inputs:   []string{"aa", "bb", "aa"},
		},
	}

	for _, c := range cases {
		ss := StringSet{}
		for _, input := range c.inputs {
			ss.Add(input)
		}
		if ss.ToJSON() != c.expected {
			t.Errorf("actual %s\nexpected %s\n", ss.ToJSON(), c.expected)
		}
	}
}
