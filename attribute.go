package cpe

import (
	"regexp"
	"strings"
)

// Attribute groups PartAttr and StringAttr.
type Attribute interface {
	WFNEncoded() string
	UrlEncoded() string
	FmtString() string
	String() string
	IsEmpty() bool
	IsValid() bool
}

// PartAttr reprecents part attribute of cpe item.
type PartAttr rune

// StringAttr reprecents other than part attribute of cpe item.
type StringAttr struct {
	raw  string
	isNa bool
}

var (
	Application      = PartAttr('a')
	OperationgSystem = PartAttr('o')
	Hardware         = PartAttr('h')
	PartNotSet       = PartAttr(0x00)
	Any              = StringAttr{}
	Na               = StringAttr{isNa: true}
)

func NewPartAttrFromWfnEncoded(str string) PartAttr {
	if len(str) != 3 {
		return PartNotSet
	}

	switch PartAttr(str[1]) {
	case Application:
		return Application
	case OperationgSystem:
		return OperationgSystem
	case Hardware:
		return Hardware
	}
	return PartNotSet
}

func NewPartAttrFromUriEncoded(str string) PartAttr {
	if len(str) != 1 {
		return PartNotSet
	}

	switch PartAttr(str[0]) {
	case Application:
		return Application
	case OperationgSystem:
		return OperationgSystem
	case Hardware:
		return Hardware
	}
	return PartNotSet
}

func NewPartAttrFromFmtEncoded(str string) PartAttr {
	if len(str) != 1 {
		return PartNotSet
	}

	switch PartAttr(str[0]) {
	case Application:
		return Application
	case OperationgSystem:
		return OperationgSystem
	case Hardware:
		return Hardware
	}
	return PartNotSet
}

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

func(m PartAttr) FmtString() string {
	return m.String()
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

func NewStringAttrFromWfnEncoded(str string) StringAttr {
	if str == "NA" {
		return Na
	} else if str == "ANY" {
		return Any
	}
	return StringAttr{
		raw: wfn_encoder.Decode(strings.TrimPrefix(strings.TrimSuffix(str, "\""), "\"")),
	}
}

func NewStringAttrFromUriEncoded(str string) StringAttr {
	if str == "-" {
		return Na
	} else if str == "" || str == "*" {
		return Any
	}
	return StringAttr{
		raw: url_encoder.Decode(str),
	}
}

func NewStringAttrFromFmtEncoded(str string) StringAttr {
	if str == "-" {
		return Na
	} else if str == "*" {
		return Any
	}
	return StringAttr{
		raw: fmt_encoder.Decode(str),
	}
}

func (s StringAttr) String() string {
	if s.isNa {
		return "-"
	} else if len(s.raw) == 0 {
		return "*"
	}

	return s.raw
}

func (s StringAttr) WFNEncoded() string {
	if s.isNa {
		return "NA"
	} else if len(s.raw) == 0 {
		return "ANY"
	}

	return "\"" + wfn_encoder.Encode(s.raw) + "\""
}

func (s StringAttr) FmtString() string {
	if s.isNa {
		return "-"
	} else if len(s.raw) == 0 {
		return "*"
	}

	return fmt_encoder.Encode(s.raw)
}

func (s StringAttr) UrlEncoded() string {
	if s.IsEmpty() {
		return "" // *
	} else if s.isNa {
		return "-"
	}
	return url_encoder.Encode(s.raw)
}

// Empty StringAttr means ANY.
func (s StringAttr) IsEmpty() bool {
	return s.raw == "" && !s.isNa
}

func (s StringAttr) IsValid() bool {
	if s.isNa && len(s.raw) != 0 {
		return false
	}

	if regexp.MustCompile("\\A(\\*|\\?+){0,1}[a-zA-Z0-9\\-_!\"#$%&'()+,./:;<=>@\\[\\]^`{}\\|~\\\\]+(\\*|\\?+){0,1}$").FindString(s.raw) != s.raw {
		return false
	}

	return true
}
