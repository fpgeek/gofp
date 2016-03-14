package gofp

import "github.com/cheekybits/genny/generic"

type KeyType generic.Type
type ValueType generic.Type

type KeyTypeValueTypeMap map[KeyType]ValueType

func (m KeyTypeValueTypeMap) Keys() []KeyType {
	list := []KeyType{}
	for k := range m {
		list = append(list, k)
	}
	return list
}

func (m KeyTypeValueTypeMap) Values() []ValueType {
	list := []ValueType{}
	for _, v := range m {
		list = append(list, v)
	}
	return list
}
