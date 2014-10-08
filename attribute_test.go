package cpe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStringAttr(t *testing.T) {
	type testcase struct {
		input string
		valid bool
	}
	var cases = []testcase{
		{"microsoft", true},
		{"microsoft&google", true},
		{"??crosoft", true},
		{"microso??", true},
		{"マイクロソフト", false},
		{"microsoft&グーグル", false},
		{"**crosoft", false},
		{"microso**", false},
		{"mic**roso", false},
		// {"-microsoft", false}, // FIXME:this case must check to invalid
	}

	for i, c := range cases {
		sa, err := newStringAttr(c.input)
		if c.valid {
			assert.Nil(t, err, "%d", i)
			assert.Equal(t, c.input, sa.raw)
		} else {
			assert.Error(t, err, "%d", i)
		}
	}
}

func TestWFNEncoded(t *testing.T) {
	type testcase struct {
		input  string
		expect string
	}
	var cases = []testcase{
		{"foo-bar", "foo\\-bar"},
		{"Acrobat_Reader", "Acrobat_Reader"},
		{"\"oh_my!\"", "\\\"oh_my\\!\\\""},
		{"g++", "g\\+\\+"},
		{"g.?", "g\\.?"},
		{"sr*", "sr*"},
		{"big$money", "big\\$money"},
		{"foo:bar", "foo\\:bar"},
		{"with_quoted~tilde", "with_quoted\\~tilde"},
		{"*SOFT*", "*SOFT*"},
		{"8.??", "8\\.??"},
		{"*8.??", "*8\\.??"},
	}

	for i, c := range cases {
		sa, err := newStringAttr(c.input)
		assert.Nil(t, err, "%d", i)
		assert.Equal(t, c.expect, sa.WFNEncode(), "%d", i)
	}
}
