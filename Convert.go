package bunit

import "github.com/cymertek/go-big"

func (b Bytes) Int() int64 {
	return (&big.Int{}).SetBytes(b).Int64()
}
func (b Bits) Int() int64 {
	return (&big.Int{}).SetBytes(b).Int64()
}
