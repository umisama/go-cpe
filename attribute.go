package cpe

import (
	"net/url"
	"regexp"
)

type PartAttr rune

var (
	Application      = PartAttr('a')
	OperationgSystem = PartAttr('o')
	Hardware         = PartAttr('h')
	PartNotSet       = PartAttr(0x00)
)

func (m PartAttr) String() string {
	if m.IsValid() {
		return string(m)
	} else {
		panic("\"%v\" is not valid as part attribute.")
	}
}

func (m PartAttr) IsValid() bool {
	switch m {
	case Application, OperationgSystem, Hardware:
		return true
	default:
		return false
	}
}

func (m PartAttr) IsEmpty() bool {
	return m == PartNotSet
}

type stringAttr struct {
	raw    string
}

func newStringAttr(str string) (stringAttr, error) {
	if !isValidAsStringAttr(str) {
		return stringAttr{}, cpeerr{reason: err_invalid_attribute_str}
	}
	return stringAttr{
		raw: str,
	}, nil
}

func (s stringAttr) Raw() string {
	return s.raw
}

func (s stringAttr) WFNEncode() string {
	encoded := s.raw
	for key, repl := range map[string]string{
		"-":   "\\-",
		"#":   "\\#",
		"\\$": "\\$",
		"%":   "\\%",
		"&":   "\\&",
		"'":   "\\'",
		"\\(": "\\(",
		"\\)": "\\)",
		"\\+": "\\+",
		",":   "\\,",
		"\\.": "\\.",
		"/":   "\\/",
		":":   "\\:",
		";":   "\\;",
		"<":   "\\<",
		"=":   "\\=",
		">":   "\\>",
		"@":   "\\@",
		"!":   "\\!",
		"\\[": "\\[",
		"\\]": "\\]",
		"\\^": "\\^",
		"`":   "\\`",
		"{":   "\\{",
		"}":   "\\}",
		"\\|": "\\|",
		"~":   "\\~",
	} {
		encoded = regexp.MustCompile(key).ReplaceAllString(encoded, repl)
	}

	return encoded
}

func (s stringAttr) UrlEncode() string {
	return url.QueryEscape(string(s.raw))
}

func (s stringAttr) IsEmpty() bool {
	return s.raw == ""
}

func isValidAsStringAttr(str string) bool {
	return regexp.MustCompile("\\A(\\*|\\?+){0,1}[a-zA-Z0-9\\-_!\"#$%&'()+,./:;<=>@\\[\\]^`{}\\|~]+(\\*|\\?+){0,1}$").FindString(str) == str
}
