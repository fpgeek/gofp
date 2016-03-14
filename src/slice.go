package gofp

import "github.com/cheekybits/genny/generic"

type ItemType generic.Type

type ItemTypeSlice []ItemType

func (sl ItemTypeSlice) Map(fn func(item ItemType) interface{}) []interface{} {
	newSl := []interface{}{}
	for _, sItem := range sl {
		newSl = append(newSl, fn(sItem))
	}
	return newSl
}

func (sl ItemTypeSlice) Filter(fn func(item ItemType) bool) []ItemType {
	newSl := []ItemType{}
	for _, sItem := range sl {
		if fn(sItem) {
			newSl = append(newSl, sItem)
		}
	}
	return newSl
}

func (sl ItemTypeSlice) Reduce(acc interface{}, fn func(sum interface{}, item ItemType) interface{}) interface{} {
	if len(sl) == 0 {
		return acc
	}

	for _, sItem := range sl[1:] {
		acc = fn(acc, sItem)
	}
	return acc
}

func (sl ItemTypeSlice) ForEach(fn func(item *ItemType)) {
	for _, sItem := range sl {
		fn(&sItem)
	}
}

func (sl ItemTypeSlice) Find(fn func(item ItemType) bool) (ItemType, bool) {
	for _, sItem := range sl {
		if fn(sItem) {
			return sItem, true
		}
	}
	var empty ItemType
	return empty, false
}

func (sl ItemTypeSlice) Every(fn func(item ItemType) bool) bool {
	for _, sItem := range sl {
		if !fn(sItem) {
			return false
		}
	}
	return true
}

func (sl ItemTypeSlice) Some(fn func(item ItemType) bool) bool {
	for _, sItem := range sl {
		if fn(sItem) {
			return true
		}
	}
	return false
}

func (sl ItemTypeSlice) GroupBy(fn func(item ItemType) string) map[string][]ItemType {
	m := map[string][]ItemType{}
	for _, sItem := range sl {
		key := fn(sItem)
		if _, ok := m[key]; !ok {
			m[key] = []ItemType{sItem}
		} else {
			m[key] = append(m[key], sItem)
		}
	}
	return m
}

func (sl ItemTypeSlice) Includes(items ...ItemType) bool {
	sItemSet := map[ItemType]bool{}
	for _, sItem := range sl {
		sItemSet[sItem] = true
	}

	for _, item := range items {
		if !sItemSet[item] {
			return false
		}
	}
	return true
}
