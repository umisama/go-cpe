package cpe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// from 5.5 WFN Example @ NISTIR-7695-CPE-Naming
func TestWfnExamples(t *testing.T) {
	// Example 1
	item := NewItem()
	item.SetPart(Application)
	item.SetVendor(NewStringAttr("microsoft"))
	item.SetProduct(NewStringAttr("internet_explorer"))
	item.SetVersion(NewStringAttr("8.0.6001"))
	item.SetUpdate(NewStringAttr("beta"))
	item.SetEdition(Na)
	assert.Equal(t, `wfn:[part="a",vendor="microsoft",product="internet_explorer",version="8\.0\.6001",update="beta",edition=NA]`, item.Wfn())

	// Example 2
	item = NewItem()
	item.SetPart(Application)
	item.SetVendor(NewStringAttr("microsoft"))
	item.SetProduct(NewStringAttr("internet_explorer"))
	item.SetVersion(NewStringAttr("8.*"))
	item.SetUpdate(NewStringAttr("sp?"))
	item.SetEdition(Na)
	item.SetLanguage(Any)
	assert.Equal(t, `wfn:[part="a",vendor="microsoft",product="internet_explorer",version="8\.*",update="sp?",edition=NA]`, item.Wfn())

	// Example 3
	item = NewItem()
	item.SetPart(Application)
	item.SetVendor(NewStringAttr("hp"))
	item.SetProduct(NewStringAttr("insight_diagnostics"))
	item.SetVersion(NewStringAttr("7.4.0.1570"))
	item.SetSwEdition(NewStringAttr("online"))
	item.SetTargetSw(NewStringAttr("windows_2003"))
	item.SetTargetHw(NewStringAttr("x64"))
	assert.Equal(t, `wfn:[part="a",vendor="hp",product="insight_diagnostics",version="7\.4\.0\.1570",sw_edition="online",target_sw="windows_2003",target_hw="x64"]`, item.Wfn())

	// Example 4
	item = NewItem()
	item.SetPart(Application)
	item.SetVendor(NewStringAttr("hp"))
	item.SetProduct(NewStringAttr("openview_network_manager"))
	item.SetVersion(NewStringAttr("7.51"))
	item.SetUpdate(Na)
	item.SetTargetSw(NewStringAttr("linux"))
	assert.Equal(t, `wfn:[part="a",vendor="hp",product="openview_network_manager",version="7\.51",update=NA,target_sw="linux"]`, item.Wfn())

	// Example 5
	item = NewItem()
	item.SetPart(Application)
	item.SetVendor(NewStringAttr("foo\\bar"))
	item.SetProduct(NewStringAttr("big$money_2010"))
	item.SetSwEdition(NewStringAttr("special"))
	item.SetTargetSw(NewStringAttr("ipod_touch"))
	assert.Equal(t, `wfn:[part="a",vendor="foo\\bar",product="big\$money_2010",sw_edition="special",target_sw="ipod_touch"]`, item.Wfn())
}

// from 6.1.2.4 Examples of binding a WFN to a URI @ NISTIR-7695-CPE-Naming
func TestUrlExamples(t *testing.T) {
	// Example 1
	item := NewItem()
	item.SetPart(Application)
	item.SetVendor(NewStringAttr("microsoft"))
	item.SetProduct(NewStringAttr("internet_explorer"))
	item.SetVersion(NewStringAttr("8.0.6001"))
	item.SetUpdate(NewStringAttr("beta"))
	item.SetEdition(Any)
	assert.Equal(t, `cpe:/a:microsoft:internet_explorer:8.0.6001:beta`, item.Uri())

	// Example 2
	item = NewItem()
	item.SetPart(Application)
	item.SetVendor(NewStringAttr("microsoft"))
	item.SetProduct(NewStringAttr("internet_explorer"))
	item.SetVersion(NewStringAttr("8.*"))
	item.SetUpdate(NewStringAttr("sp?"))
	assert.Equal(t, `cpe:/a:microsoft:internet_explorer:8.%02:sp%01`, item.Uri())

	// Example3
	item = NewItem()
	item.SetPart(Application)
	item.SetVendor(NewStringAttr("hp"))
	item.SetProduct(NewStringAttr("insight_diagnostics"))
	item.SetVersion(NewStringAttr("7.4.0.1570"))
	item.SetUpdate(Na)
	item.SetSwEdition(NewStringAttr("online"))
	item.SetTargetSw(NewStringAttr("win2003"))
	item.SetTargetHw(NewStringAttr("x64"))
	assert.Equal(t, `cpe:/a:hp:insight_diagnostics:7.4.0.1570:-:~~online~win2003~x64~`, item.Uri())

	// Example 4
	item = NewItem()
	item.SetPart(Application)
	item.SetVendor(NewStringAttr("hp"))
	item.SetProduct(NewStringAttr("openview_network_manager"))
	item.SetVersion(NewStringAttr("7.51"))
	item.SetTargetSw(NewStringAttr("linux"))
	assert.Equal(t, `cpe:/a:hp:openview_network_manager:7.51::~~~linux~~`, item.Uri())

	// Example 5
	item = NewItem()
	item.SetPart(Application)
	item.SetVendor(NewStringAttr("foo\\bar"))
	item.SetProduct(NewStringAttr("big$money_manager_2010"))
	item.SetSwEdition(NewStringAttr("special"))
	item.SetTargetSw(NewStringAttr("ipod_touch"))
	item.SetTargetHw(NewStringAttr("80gb"))
	assert.Equal(t, `cpe:/a:foo%5cbar:big%24money_manager_2010:::~~special~ipod_touch~80gb~`, item.Uri())
}

func TestFormattedString(t *testing.T) {
	// Example 1
	item := NewItem()
	item.SetPart(Application)
	item.SetVendor(NewStringAttr("microsoft"))
	item.SetProduct(NewStringAttr("internet_explorer"))
	item.SetVersion(NewStringAttr("8.0.6001"))
	item.SetUpdate(NewStringAttr("beta"))
	assert.Equal(t, `cpe:2.3:a:microsoft:internet_explorer:8.0.6001:beta:*:*:*:*:*:*`, item.Formatted())

	// Example 2
	item = NewItem()
	item.SetPart(Application)
	item.SetVendor(NewStringAttr("microsoft"))
	item.SetProduct(NewStringAttr("internet_explorer"))
	item.SetVersion(NewStringAttr("8.*"))
	item.SetUpdate(NewStringAttr("sp?"))
	item.SetLanguage(Any)
	assert.Equal(t, `cpe:2.3:a:microsoft:internet_explorer:8.*:sp?:*:*:*:*:*:*`, item.Formatted())

	// Example 3
	item = NewItem()
	item.SetPart(Application)
	item.SetVendor(NewStringAttr("hp"))
	item.SetProduct(NewStringAttr("insight_diagnostics"))
	item.SetVersion(NewStringAttr("7.4.0.1570"))
	item.SetUpdate(Na)
	item.SetSwEdition(NewStringAttr("online"))
	item.SetTargetSw(NewStringAttr("win2003"))
	item.SetTargetHw(NewStringAttr("x64"))
	assert.Equal(t, `cpe:2.3:a:hp:insight_diagnostics:7.4.0.1570:-:*:*:online:win2003:x64:*`, item.Formatted())

	// Example 4
	item = NewItem()
	item.SetPart(Application)
	item.SetVendor(NewStringAttr("hp"))
	item.SetProduct(NewStringAttr("openview_network_manager"))
	item.SetVersion(NewStringAttr("7.51"))
	item.SetTargetSw(NewStringAttr("linux"))
	assert.Equal(t, `cpe:2.3:a:hp:openview_network_manager:7.51:*:*:*:*:linux:*:*`, item.Formatted())

	// Example 5
	item = NewItem()
	item.SetPart(Application)
	item.SetVendor(NewStringAttr("foo\\bar"))
	item.SetProduct(NewStringAttr("big$money_2010"))
	item.SetSwEdition(NewStringAttr("special"))
	item.SetTargetSw(NewStringAttr("ipod_touch"))
	item.SetTargetHw(NewStringAttr("80gb"))
	assert.Equal(t, `cpe:2.3:a:foo\\bar:big\$money_2010:*:*:*:*:special:ipod_touch:80gb:*`, item.Formatted())
}
