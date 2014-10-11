package cpe

import (
	"regexp"
	"strings"
)

type Attribute interface {
	WFNEncoded() string
	UrlEncoded() string
	IsEmpty() bool
	IsValid() bool
}

type PartAttr rune

type StringAttr struct {
	raw   string
	isAny bool
	isNa  bool
}

var (
	Application      = PartAttr('a')
	OperationgSystem = PartAttr('o')
	Hardware         = PartAttr('h')
	PartNotSet       = PartAttr(0x00)
	Any              = StringAttr{isAny: true}
	Na               = StringAttr{isNa: true}
)

func (m PartAttr) String() string {
	if m.IsValid() {
		return string(m)
	} else {
		panic("\"%v\" is not valid as part attribute.")
	}
}

func (m PartAttr) WFNEncoded() string {
	return "\"" + m.String() + "\""
}

func (m PartAttr) UrlEncoded() string {
	return m.String()
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

func NewStringAttr(str string) StringAttr {
	return StringAttr{
		raw: str,
	}
}

func (s StringAttr) Raw() string {
	return s.raw
}

func (s StringAttr) WFNEncoded() string {
	if s.isNa {
		return "NA"
	} else if s.isAny {
		return "ANY"
	}

	encoded := strings.Replace(s.raw, "\\", "\\\\", -1)
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
		"\"":  "\\\"",
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

	return "\"" + encoded + "\""
}

func (s StringAttr) UrlEncoded() string {
	if s.isAny {
		return "" // *
	} else if s.isNa {
		return "-"
	}
	return url_encoder.Encode(s.raw)
}

func (s StringAttr) IsEmpty() bool {
	return s.raw == "" && !s.isNa && !s.isAny
}

func (s StringAttr) IsValid() bool {
	if s.isAny && s.isNa {
		return false
	}

	if (s.isAny || s.isNa) && s.raw != "" {
		return false
	}

	if regexp.MustCompile("\\A(\\*|\\?+){0,1}[a-zA-Z0-9\\-_!\"#$%&'()+,./:;<=>@\\[\\]^`{}\\|~\\\\]+(\\*|\\?+){0,1}$").FindString(s.raw) != s.raw {
		return false
	}

	return true
}
