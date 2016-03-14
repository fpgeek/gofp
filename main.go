package gofp

//go:generate genny -in=src/slice/slice.go -out=slice.go gen "ItemType=string,int"
//go:generate genny -in=src/map/map.go -out=map.go gen "KeyType=string,int ValueType=string,int"
