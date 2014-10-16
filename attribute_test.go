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
		sa := NewStringAttr(c.input)
		assert.Equal(t, c.valid, sa.IsValid(), "%d", i)
	}
}

func TestWFNEncoded(t *testing.T) {
	type testcase struct {
		input  string
		expect string
	}
	var cases = []testcase{
		{"foo-bar", "\"foo\\-bar\""},
		{"Acrobat_Reader", "\"Acrobat_Reader\""},
		{"\"oh_my!\"", "\"\\\"oh_my\\!\\\"\""},
		{"g++", "\"g\\+\\+\""},
		{"g.?", "\"g\\.?\""},
		{"sr*", "\"sr*\""},
		{"big$money", "\"big\\$money\""},
		{"foo:bar", "\"foo\\:bar\""},
		{"with_quoted~tilde", "\"with_quoted\\~tilde\""},
		{"*SOFT*", "\"*SOFT*\""},
		{"8.??", "\"8\\.??\""},
		{"*8.??", "\"*8\\.??\""},
	}

	for i, c := range cases {
		sa := NewStringAttr(c.input)
		assert.Equal(t, c.expect, sa.wfnEncoded(), "%d", i)
	}
}

func TestNewStringAttrFromWcnEncoded(t *testing.T) {
	type testcase struct {
		input  string
		expect string
	}
	var cases = []testcase{
		{"\"foo\\-bar\"", "foo-bar"},
		{"\"Acrobat_Reader\"", "Acrobat_Reader"},
		{"\"\\\"oh_my\\!\\\"\"", "\"oh_my!\""},
		{"\"g\\+\\+\"", "g++"},
		{"\"g\\.?\"", "g.?"},
		{"\"sr*\"", "sr*"},
		{"\"big\\$money\"", "big$money"},
		{"\"foo\\:bar\"", "foo:bar"},
		{"\"with_quoted\\~tilde\"", "with_quoted~tilde"},
		{"\"*SOFT*\"", "*SOFT*"},
		{"\"8\\.??\"", "8.??"},
		{"\"*8\\.??\"", "*8.??"},
	}

	for i, c := range cases {
		sa := newStringAttrFromWfnEncoded(c.input)
		assert.Equal(t, c.expect, sa.raw, "%d", i)
	}
}

func TestNewPartAttrFromWcnEncoded(t *testing.T) {
	type testcase struct {
		input  string
		expect PartAttr
	}
	var cases = []testcase{
		{`"a"`, Application},
		{`"o"`, OperationgSystem},
		{`"h"`, Hardware},
		{`"z"`, PartNotSet},
	}

	for i, c := range cases {
		pa := newPartAttrFromWfnEncoded(c.input)
		assert.Equal(t, c.expect, pa, "%d", i)
	}
}
