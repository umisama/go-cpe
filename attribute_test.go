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

func TestMatchWildcard(t *testing.T) {
	type testcase struct {
		input  string
		value  string
		expect bool
	}
	var cases = []testcase{
		{"*123", "11123", true},
		{"*123", "11123a", false},
		{"123*", "12311", true},
		{"123*", "112311", false},
		{"??123", "11123", true},
		{"??123", "1123", false},
		{"??123", "11123a", false},
		{"123??", "12311", true},
		{"123??", "123111", false},
		{"123??", "112311", false},
		{"*123?", "111233", true},
		{"*123*", "11123111", true},
		{"?123?", "11231", true},
		{"?123*", "112333", true},
		{"??123*", "1112333", true},
		{"*123??", "1112335", true},
		{"*123??", "11123355", false},
		{"??123*", "18112333", false},
	}

	for i, c := range cases {
		assert.Equal(t, c.expect, match_wildcard(c.input, c.value), "%d", i)
	}
}

func testComparition(t *testing.T) {
	type testcase struct {
		input  Attribute
		value  Attribute
		expect Relation
	}
	var cases = []testcase{
		{Application, Application, Equal},
		{NewStringAttr("Adobe"), Any, Subset},
		{Any, NewStringAttr("Reader"), Superset},
		{NewStringAttr("9.*"), NewStringAttr("9.3.2"), Superset},
		{Any, Na, Superset},
		{NewStringAttr("PalmOS"), Na, Disjoint},
	}

	for i, c := range cases {
		assert.Equal(t, c.expect, c.input.Comparison(c.value), "%d", i)
	}
}

func BenchmarkStringAttrComparison(b *testing.B) {
	sa1 := NewStringAttr("hellohellohellohello")
	sa2 := NewStringAttr("worldworldworldworld")
	for i := 0; i < b.N; i++ {
		sa1.Comparison(sa2)
	}
}
