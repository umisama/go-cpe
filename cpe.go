package cpe

import (
	"fmt"
	"strings"
)

type Item struct {
	part       PartAttr
	vendor     StringAttr
	product    StringAttr
	version    StringAttr
	update     StringAttr
	edition    StringAttr
	language   StringAttr
	sw_edition StringAttr
	target_sw  StringAttr
	target_hw  StringAttr
	other      StringAttr
}

func NewItem() *Item {
	return &Item{
		part:       PartNotSet,
		vendor:     Any,
		product:    Any,
		version:    Any,
		update:     Any,
		edition:    Any,
		language:   Any,
		sw_edition: Any,
		target_sw:  Any,
		target_hw:  Any,
		other:      Any,
	}
}

func NewItemFromWfn(wfn string) (*Item, error) {
	if strings.HasPrefix(wfn, "wfn:[") {
		wfn = strings.TrimPrefix(wfn, "wfn:[")
	} else {
		return nil, cpeerr{reason: err_invalid_wfn}
	}

	if strings.HasSuffix(wfn, "]") {
		wfn = strings.TrimSuffix(wfn, "]")
	} else {
		return nil, cpeerr{reason: err_invalid_wfn}
	}

	item := NewItem()
	for _, attr := range strings.Split(wfn, ",") {
		sepattr := strings.Split(attr, "=")
		if len(sepattr) != 2 {
			return nil, cpeerr{reason: err_invalid_wfn}
		}

		n, v := sepattr[0], sepattr[1]
		switch n {
		case "part":
			item.part = NewPartAttrFromWfnEncoded(v)
		case "vendor":
			item.vendor = NewStringAttrFromWfnEncoded(v)
		case "product":
			item.product = NewStringAttrFromWfnEncoded(v)
		case "version":
			item.version = NewStringAttrFromWfnEncoded(v)
		case "update":
			item.update = NewStringAttrFromWfnEncoded(v)
		case "edition":
			item.edition = NewStringAttrFromWfnEncoded(v)
		case "language":
			item.language = NewStringAttrFromWfnEncoded(v)
		case "sw_edition":
			item.sw_edition = NewStringAttrFromWfnEncoded(v)
		case "target_sw":
			item.target_sw = NewStringAttrFromWfnEncoded(v)
		case "target_hw":
			item.target_hw = NewStringAttrFromWfnEncoded(v)
		case "other":
			item.other = NewStringAttrFromWfnEncoded(v)
		}
	}

	return item, nil
}

func NewItemFromUri(uri string) (*Item, error) {
	if strings.HasPrefix(uri, "cpe:/") {
		uri = strings.TrimPrefix(uri, "cpe:/")
	} else {
		return nil, cpeerr{reason: err_invalid_wfn}
	}

	item := NewItem()
	for i, attr := range strings.Split(uri, ":") {
		switch i {
		case 0:
			item.part = NewPartAttrFromUriEncoded(attr)
		case 1:
			item.vendor = NewStringAttrFromUriEncoded(attr)
		case 2:
			item.product = NewStringAttrFromUriEncoded(attr)
		case 3:
			item.version = NewStringAttrFromUriEncoded(attr)
		case 4:
			item.update = NewStringAttrFromUriEncoded(attr)
		case 5:
			editions := strings.Split(attr, "~")
			if len(editions) == 1 {
				item.edition = NewStringAttrFromUriEncoded(editions[0])
			} else if len(editions) == 6 {
				item.edition = NewStringAttrFromUriEncoded(editions[1])
				item.sw_edition = NewStringAttrFromUriEncoded(editions[2])
				item.target_sw = NewStringAttrFromUriEncoded(editions[3])
				item.target_hw = NewStringAttrFromUriEncoded(editions[4])
				item.other = NewStringAttrFromUriEncoded(editions[5])
			} else {
				return nil, cpeerr{reason: err_invalid_wfn}
			}
		}
	}
	return item, nil
}

func NewItemFromFormattedString(str string) (*Item, error) {
	return &Item{
		part:       PartNotSet,
		vendor:     Any,
		product:    Any,
		version:    Any,
		update:     Any,
		edition:    Any,
		language:   Any,
		sw_edition: Any,
		target_sw:  Any,
		target_hw:  Any,
		other:      Any,
	}, nil
}

func (m *Item) Wfn() string {
	wfn := "wfn:["
	first := true

	for _, it := range []struct {
		name string
		attr Attribute
	}{
		{"part", m.part},
		{"vendor", m.vendor},
		{"product", m.product},
		{"version", m.version},
		{"update", m.update},
		{"edition", m.edition},
		{"language", m.language},
		{"sw_edition", m.sw_edition},
		{"target_sw", m.target_sw},
		{"target_hw", m.target_hw},
		{"other", m.other},
	} {
		if !it.attr.IsEmpty() {
			if first {
				first = false
			} else {
				wfn += ","
			}
			wfn += it.name + "=" + it.attr.WFNEncoded()
		}
	}
	wfn += "]"

	return wfn
}

func (m *Item) Uri() string {
	uri := "cpe:/"

	l := []struct {
		name string
		attr Attribute
	}{
		{"part", m.part},
		{"vendor", m.vendor},
		{"product", m.product},
		{"version", m.version},
		{"update", m.update},
	}

	for c, it := range l {
		if !it.attr.IsEmpty() {
			uri += it.attr.UrlEncoded()
		}
		if c+1 != len(l) {
			uri += ":"
		}
	}

	if m.target_hw.UrlEncoded() != "" ||
		m.target_sw.UrlEncoded() != "" ||
		m.sw_edition.UrlEncoded() != "" ||
		m.other.UrlEncoded() != "" {
		uri += ":~" + m.edition.UrlEncoded()
		uri += "~" + m.sw_edition.UrlEncoded()
		uri += "~" + m.target_sw.UrlEncoded()
		uri += "~" + m.target_hw.UrlEncoded()
		uri += "~" + m.other.UrlEncoded()
	} else {
		uri += ":" + m.edition.UrlEncoded()
	}

	uri += ":" + m.language.UrlEncoded()
	return strings.TrimRight(uri, ":*")
}

func (m *Item) Formatted() string {
	fmted := "cpe:2.3"

	for _, it := range []Attribute{
		m.part, m.vendor, m.product, m.version, m.update, m.edition, m.language, m.sw_edition, m.target_sw, m.target_hw, m.other,
	} {
		if !it.IsEmpty() {
			fmted += ":" + it.FmtString()
		} else {
			fmted += ":*"

		}
	}
	return fmted
}

func (i *Item) SetPart(p PartAttr) error {
	if !p.IsValid() {
		return cpeerr{reason: err_invalid_type, attr: []interface{}{p, "part"}}
	}

	i.part = p
	return nil
}

func (i *Item) Part() PartAttr {
	return i.part
}

func (i *Item) SetVendor(s StringAttr) error {
	if !s.IsValid() {
		return cpeerr{reason: err_invalid_attribute_str}
	}

	i.vendor = s
	return nil
}

func (i *Item) Vendor() StringAttr {
	return i.vendor
}

func (i *Item) SetProduct(s StringAttr) error {
	if !s.IsValid() {
		return cpeerr{reason: err_invalid_attribute_str}
	}

	i.product = s
	return nil
}

func (i *Item) Product() StringAttr {
	return i.product
}

func (i *Item) SetVersion(s StringAttr) error {
	if !s.IsValid() {
		return cpeerr{reason: err_invalid_attribute_str}
	}

	i.version = s
	return nil
}

func (i *Item) Version() StringAttr {
	return i.version
}

func (i *Item) SetUpdate(s StringAttr) error {
	if !s.IsValid() {
		return cpeerr{reason: err_invalid_attribute_str}
	}

	i.update = s
	return nil
}

func (i *Item) Update() StringAttr {
	return i.update
}

func (i *Item) SetEdition(s StringAttr) error {
	if !s.IsValid() {
		return cpeerr{reason: err_invalid_attribute_str}
	}

	i.edition = s
	return nil
}

func (i *Item) Edition() StringAttr {
	return i.edition
}

func (i *Item) SetLanguage(s StringAttr) error {
	if !s.IsValid() {
		return cpeerr{reason: err_invalid_attribute_str}
	}

	i.language = s
	return nil
}

func (i *Item) Language() StringAttr {
	return i.language
}

func (i *Item) SetSwEdition(s StringAttr) error {
	if !s.IsValid() {
		return cpeerr{reason: err_invalid_attribute_str}
	}

	i.sw_edition = s
	return nil
}

func (i *Item) SwEdition() StringAttr {
	return i.sw_edition
}

func (i *Item) SetTargetSw(s StringAttr) error {
	if !s.IsValid() {
		return cpeerr{reason: err_invalid_attribute_str}
	}

	i.target_sw = s
	return nil
}

func (i *Item) TargetSw() StringAttr {
	return i.target_sw
}

func (i *Item) SetTargetHw(s StringAttr) error {
	if !s.IsValid() {
		return cpeerr{reason: err_invalid_attribute_str}
	}

	i.target_hw = s
	return nil
}

func (i *Item) TargetHw() StringAttr {
	return i.target_hw
}

func (i *Item) SetOther(s StringAttr) error {
	if !s.IsValid() {
		return cpeerr{reason: err_invalid_attribute_str}
	}

	i.other = s
	return nil
}

func (i *Item) Other() StringAttr {
	return i.other
}

type cpeerr struct {
	reason string
	attr   []interface{}
}

var (
	err_invalid_type          = "\"%#v\" is not valid as %v attribute."
	err_invalid_attribute_str = "invalid attribute string."
	err_invalid_wfn           = "invalid wfn string."
)

func (e cpeerr) Error() string {
	return fmt.Sprintf("cpe:"+e.reason, e.attr...)
}
