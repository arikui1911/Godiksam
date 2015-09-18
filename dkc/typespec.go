package dkc

type baseType int

const (
	ttBOOLEAN baseType = iota
	ttINT
	ttDOUBLE
	ttSTRING
)

type typeSpec struct {
	base     baseType
	subtypes []subtype_t
}

type subtype_t interface{}

/*
type funcSubType struct {
	params []param_t
}
*/
