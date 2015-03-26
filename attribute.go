package cpe

import (
	"regexp"
	"strings"
)

// Attribute groups PartAttr and StringAttr.
type Attribute interface {
	wfnEncoded() string
	urlEncoded() string
	fmtString() string
	String() string
	IsEmpty() bool
	IsValid() bool
	Comparison(Attribute) Relation
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

func newPartAttrFromWfnEncoded(str string) PartAttr {
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

func newPartAttrFromUriEncoded(str string) PartAttr {
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

func newPartAttrFromFmtEncoded(str string) PartAttr {
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

func (m PartAttr) wfnEncoded() string {
	return "\"" + m.String() + "\""
}

func (m PartAttr) fmtString() string {
	return m.String()
}

func (m PartAttr) urlEncoded() string {
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

func (src PartAttr) Comparison(trg Attribute) Relation {
	trg_part, ok := trg.(PartAttr)
	if !ok {
		return Undefined
	}

	if !src.IsValid() || !trg_part.IsValid() {
		return Undefined
	}

	if src == trg_part {
		return Equal
	}

	return Disjoint
}

// NewStringAttr returns attribute of item with str.
func NewStringAttr(str string) StringAttr {
	return StringAttr{
		raw: str,
	}
}

func newStringAttrFromWfnEncoded(str string) StringAttr {
	if str == "NA" {
		return Na
	} else if str == "ANY" {
		return Any
	}
	return StringAttr{
		raw: wfn_encoder.Decode(strings.TrimPrefix(strings.TrimSuffix(str, "\""), "\"")),
	}
}

func newStringAttrFromUriEncoded(str string) StringAttr {
	if str == "-" {
		return Na
	} else if str == "" || str == "*" {
		return Any
	}
	return StringAttr{
		raw: url_encoder.Decode(str),
	}
}

func newStringAttrFromFmtEncoded(str string) StringAttr {
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

func (s StringAttr) wfnEncoded() string {
	if s.isNa {
		return "NA"
	} else if len(s.raw) == 0 {
		return "ANY"
	}

	return "\"" + wfn_encoder.Encode(s.raw) + "\""
}

func (s StringAttr) fmtString() string {
	if s.isNa {
		return "-"
	} else if len(s.raw) == 0 {
		return "*"
	}

	return fmt_encoder.Encode(s.raw)
}

func (s StringAttr) urlEncoded() string {
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

var stringAttrIsValidRegExp = regexp.MustCompile("\\A(\\*|\\?+){0,1}[a-zA-Z0-9\\-_!\"#$%&'()+,./:;<=>@\\[\\]^`{}\\|~\\\\]+(\\*|\\?+){0,1}$")

func (s StringAttr) IsValid() bool {
	if s.isNa && len(s.raw) != 0 {
		return false
	}

	if stringAttrIsValidRegExp.FindString(s.raw) != s.raw {
		return false
	}

	return true
}

func (src StringAttr) Comparison(trg Attribute) Relation {
	trg_str, ok := trg.(StringAttr)
	if !ok {
		return Undefined
	}

	if !src.IsValid() || !trg_str.IsValid() {
		return Undefined
	}

	if src == Any {
		if trg_str == Any {
			return Equal
		} else if trg_str == Na {
			return Superset
		} else if !trg_str.withWildCard() {
			return Superset
		}
		return Undefined
	}

	if src == Na {
		if trg_str == Any {
			return Subset
		} else if trg_str == Na {
			return Equal
		} else if !trg_str.withWildCard() {
			return Disjoint
		}
		return Undefined
	}

	if src.withWildCard() {
		if trg_str == Any {
			return Subset
		} else if trg_str == Na {
			return Disjoint
		} else if trg_str.withWildCard() {
			return Undefined
		} else if match_wildcard(src.raw, trg_str.raw) {
			return Superset
		}
		return Disjoint
	} else {
		if trg_str == Any {
			return Subset
		} else if trg_str == Na {
			return Disjoint
		} else if trg_str.withWildCard() {
			return Undefined
		} else if trg_str.raw == src.raw {
			return Equal
		}
		return Disjoint
	}

	return Undefined
}

func (m StringAttr) withWildCard() bool {
	prefix, suffix := m.raw[0], m.raw[len(m.raw)-1]
	return prefix == '*' || prefix == '?' || suffix == '*' || suffix == '?'
}

func match_wildcard(src, trg string) bool {
	sufw, sufq, prew, preq := 0, 0, 0, 0
	if strings.HasPrefix(src, "?") {
		before := len(src)
		src = strings.TrimLeft(src, "?")
		preq = before - len(src)
	}
	if strings.HasPrefix(src, "*") {
		src = strings.TrimPrefix(src, "*")
		prew = 1
	}
	if strings.HasSuffix(src, "?") {
		before := len(src)
		src = strings.TrimRight(src, "?")
		sufq = before - len(src)
	}
	if strings.HasSuffix(src, "*") {
		src = strings.TrimSuffix(src, "*")
		sufw = 1
	}

	i := strings.Index(trg, src)
	if prew != 0 {
		if i != len(trg)-len(src)-sufq && sufw == 0 {
			return false
		}
	}
	if sufw != 0 {
		if i != preq && prew == 0 {
			return false
		}
	}
	if preq != 0 {
		if i != preq || (i != len(trg)-len(src)-sufq && sufw == 0) {
			return false
		}
	}
	if sufq != 0 {
		if i != len(trg)-sufq-len(src) || (i != preq && prew == 0) {
			return false
		}
	}

	return true
}
