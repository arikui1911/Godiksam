package dkc

type baseType int

const (
	ttBOOLEAN baseType = iota
	ttINT
	ttDOUBLE
	ttSTRING
)

type TypeSpec struct {
	base     baseType
	subtypes []subtype_t
}

func newTypeSpec(base baseType) *TypeSpec {
	return &TypeSpec{base: base}
}

type subtype_t interface{}

/*
type funcSubType struct {
	params []param_t
}
*/
