package cpe

import (
	"fmt"
)

type Item struct {
	part       Part
	vendor     stringAttr
	product    stringAttr
	version    stringAttr
	update     stringAttr
	edition    stringAttr
	language   stringAttr
	sw_edition stringAttr
	target_sw  stringAttr
	target_hw  stringAttr
	other      stringAttr
}

func (i *Item) SetPart(p Part) error {
	if !p.IsValid() {
		return cpeerr{reason: err_invalid_type, attr: []interface{}{p, "part"}}
	}

	i.part = p
	return nil
}

func (i *Item) Part() Part {
	return i.part
}

func (i *Item) SetVendor(m string) error {
	var err error
	i.vendor, err = newStringAttr(m)
	if err != nil {
		return err
	}
	return nil
}

func (i *Item) Vendor() string {
	return i.vendor.Raw()
}

func (i *Item) SetProduct(m string) error {
	var err error
	i.product, err = newStringAttr(m)
	if err != nil {
		return err
	}
	return nil
}

func (i *Item) Product() string {
	return i.product.Raw()
}

func (i *Item) SetVersion(m string) error {
	var err error
	i.version, err = newStringAttr(m)
	if err != nil {
		return err
	}
	return nil
}

func (i *Item) Version() string {
	return i.version.Raw()
}

func (i *Item) SetUpdate(m string) error {
	var err error
	i.update, err = newStringAttr(m)
	if err != nil {
		return err
	}
	return nil
}

func (i *Item) Update() string {
	return i.update.Raw()
}

func (i *Item) SetEdition(m string) error {
	var err error
	i.edition, err = newStringAttr(m)
	if err != nil {
		return err
	}
	return nil
}

func (i *Item) Edition() string {
	return i.edition.Raw()
}

func (i *Item) SetLanguage(m string) error {
	var err error
	i.language, err = newStringAttr(m)
	if err != nil {
		return err
	}
	return nil
}

func (i *Item) Language() string {
	return i.language.Raw()
}

func (i *Item) SetSwEdition(m string) error {
	var err error
	i.sw_edition, err = newStringAttr(m)
	if err != nil {
		return err
	}
	return nil
}

func (i *Item) SwEdition() string {
	return i.sw_edition.Raw()
}

func (i *Item) SetTargetSw(m string) error {
	var err error
	i.target_sw, err = newStringAttr(m)
	if err != nil {
		return err
	}
	return nil
}

func (i *Item) TargetSw() string {
	return i.target_sw.Raw()
}

func (i *Item) SetTargetHw(m string) error {
	var err error
	i.target_hw, err = newStringAttr(m)
	if err != nil {
		return err
	}
	return nil
}

func (i *Item) TargetHw() string {
	return i.target_hw.Raw()
}

func (i *Item) SetOther(m string) error {
	var err error
	i.other, err = newStringAttr(m)
	if err != nil {
		return err
	}
	return nil
}

func (i *Item) Other() string {
	return i.other.Raw()
}

type cpeerr struct {
	reason string
	attr   []interface{}
}

var (
	err_invalid_type = "\"%#v\" is not valid as %v attribute."
	err_invalid_attribute_str = "invalid attribute string."
)

func (e cpeerr) Error() string {
	return fmt.Sprintf("cpe:"+e.reason, e.attr...)
}
