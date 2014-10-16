package cpe

type Relation int

const (
	Disjoint = Relation(iota)
	Equal
	Subset
	Superset
	Undefined
)

// CheckDisjoint implements CPE_DISJOINT.  Returns true if the set-theoretic reration between the names is DISJOINT.
func CheckDisjoint(src, trg *Item) bool {
	type obj struct {
		src Attribute
		trg Attribute
	}
	for _, v := range []obj{
		{src.part, trg.part},
		{src.vendor, trg.vendor},
		{src.product, trg.product},
		{src.version, trg.version},
		{src.update, trg.update},
		{src.edition, trg.edition},
		{src.language, trg.language},
		{src.sw_edition, trg.sw_edition},
		{src.target_sw, trg.target_sw},
		{src.target_hw, trg.target_hw},
		{src.other, trg.other},
	} {
		switch v.src.Comparison(v.trg) {
		case Disjoint:
			return true
		}
	}
	return false
}

// CheckEqual implements CPE_EQUAL.  Returns true if the set-theoretic relation between src and trg is EQUAL.
func CheckEqual(src, trg *Item) bool {
	type obj struct {
		src Attribute
		trg Attribute
	}
	for _, v := range []obj{
		{src.part, trg.part},
		{src.vendor, trg.vendor},
		{src.product, trg.product},
		{src.version, trg.version},
		{src.update, trg.update},
		{src.edition, trg.edition},
		{src.language, trg.language},
		{src.sw_edition, trg.sw_edition},
		{src.target_sw, trg.target_sw},
		{src.target_hw, trg.target_hw},
		{src.other, trg.other},
	} {
		switch v.src.Comparison(v.trg) {
		case Equal:
		default:
			return false
		}
	}
	return true
}

// CheckSubset implements CPE_SUBSET.  Returns true if the set-theoretic relation between src and trg is SUBSET.
func CheckSubset(src, trg *Item) bool {
	type obj struct {
		src Attribute
		trg Attribute
	}
	for _, v := range []obj{
		{src.part, trg.part},
		{src.vendor, trg.vendor},
		{src.product, trg.product},
		{src.version, trg.version},
		{src.update, trg.update},
		{src.edition, trg.edition},
		{src.language, trg.language},
		{src.sw_edition, trg.sw_edition},
		{src.target_sw, trg.target_sw},
		{src.target_hw, trg.target_hw},
		{src.other, trg.other},
	} {
		switch v.src.Comparison(v.trg) {
		case Subset, Equal:
		default:
			return false
		}
	}
	return true
}

// CheckSuperset implements CPE_SUPERSET.  Returns true if the set-theoretic relation between src and trg is SUPERSET.
func CheckSuperset(src, trg *Item) bool {
	type obj struct {
		src Attribute
		trg Attribute
	}
	for _, v := range []obj{
		{src.part, trg.part},
		{src.vendor, trg.vendor},
		{src.product, trg.product},
		{src.version, trg.version},
		{src.update, trg.update},
		{src.edition, trg.edition},
		{src.language, trg.language},
		{src.sw_edition, trg.sw_edition},
		{src.target_sw, trg.target_sw},
		{src.target_hw, trg.target_hw},
		{src.other, trg.other},
	} {
		switch v.src.Comparison(v.trg) {
		case Superset, Equal:
		default:
			return false
		}
	}
	return true
}
