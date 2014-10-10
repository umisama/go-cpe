package cpe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// from 5.5 WFN Example @ NISTIR-7695-CPE-Naming
func TestWfnExamples(t *testing.T) {
	item := NewItem()
	item.SetPart(Application)
	item.SetVendor(NewStringAttr("microsoft"))
	item.SetProduct(NewStringAttr("internet_explorer"))
	item.SetVersion(NewStringAttr("8.0.6001"))
	item.SetUpdate(NewStringAttr("beta"))
	item.SetEdition(Na)
	assert.Equal(t, `wfn:[part="a",vendor="microsoft",product="internet_explorer",version="8\.0\.6001",update="beta",edition=NA]`, item.Wfn())

	item = NewItem()
	item.SetPart(Application)
	item.SetVendor(NewStringAttr("microsoft"))
	item.SetProduct(NewStringAttr("internet_explorer"))
	item.SetVersion(NewStringAttr("8.*"))
	item.SetUpdate(NewStringAttr("sp?"))
	item.SetEdition(Na)
	item.SetLanguage(Any)
	assert.Equal(t, `wfn:[part="a",vendor="microsoft",product="internet_explorer",version="8\.*",update="sp?",edition=NA,language=ANY]`, item.Wfn())

	item = NewItem()
	item.SetPart(Application)
	item.SetVendor(NewStringAttr("hp"))
	item.SetProduct(NewStringAttr("insight_diagnostics"))
	item.SetVersion(NewStringAttr("7.4.0.1570"))
	item.SetSwEdition(NewStringAttr("online"))
	item.SetTargetSw(NewStringAttr("windows_2003"))
	item.SetTargetHw(NewStringAttr("x64"))
	assert.Equal(t, `wfn:[part="a",vendor="hp",product="insight_diagnostics",version="7\.4\.0\.1570",sw_edition="online",target_sw="windows_2003",target_hw="x64"]`, item.Wfn())

	item = NewItem()
	item.SetPart(Application)
	item.SetVendor(NewStringAttr("hp"))
	item.SetProduct(NewStringAttr("openview_network_manager"))
	item.SetVersion(NewStringAttr("7.51"))
	item.SetUpdate(Na)
	item.SetTargetSw(NewStringAttr("linux"))
	assert.Equal(t, `wfn:[part="a",vendor="hp",product="openview_network_manager",version="7\.51",update=NA,target_sw="linux"]`, item.Wfn())

	item = NewItem()
	item.SetPart(Application)
	item.SetVendor(NewStringAttr("foo\\bar"))
	item.SetProduct(NewStringAttr("big$money_2010"))
	item.SetSwEdition(NewStringAttr("special"))
	item.SetTargetSw(NewStringAttr("ipod_touch"))
	assert.Equal(t, `wfn:[part="a",vendor="foo\\bar",product="big\$money_2010",sw_edition="special",target_sw="ipod_touch"]`, item.Wfn())
}
