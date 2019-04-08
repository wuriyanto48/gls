package gls

import (
	"time"
)

// Asset will describe file or directory
type Asset struct {
	Name         string
	Size         int64
	Type         string
	ModifiedDate time.Time
}

// SizeInKilobByte will convert asset size to kilo byte
func (a *Asset) SizeInKilobByte() float64 {
	return float64(a.Size) / 1000.0
}

// Assets type list of Asset
type Assets []*Asset
