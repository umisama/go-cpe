package cpe

import (
	"fmt"
)

type Item struct {
	part       Part
	vendor     string
	product    string
	version    string
	update     string
	edition    string
	language   string
	sw_edition string
	target_sw  string
	target_hw  string
	other      string
}

func (i *Item) SetPart(p Part) error {
	if !p.IsValid() {
		return cpeerr{reason: err_invalid_type, attr: []interface{}{p, "part"}}
	}

	i.part = p
	return nil
}

type cpeerr struct {
	reason string
	attr   []interface{}
}

var (
	err_invalid_type = "\"%#v\" is not valid as %v attribute."
)

func (e cpeerr) Error() string {
	return fmt.Sprintf("cpe:"+e.reason, e.attr...)
}
