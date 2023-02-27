package bunit

import (
	"time"

	"github.com/cymertek/go-big"
)

func (b Bytes) Int64() int64 {
	return (&big.Int{}).SetBytes(b).Int64()
}
func (b Bits) Int64() int64 {
	return (&big.Int{}).SetBytes(b).Int64()
}
func (b BitRate) Int64() int64 {
	v := (&big.Float{}).SetBytes(b.n, []byte{})
	v.Mul(v, big.NewFloat(float64(time.Second)/float64(b.d)))
	r, _ := v.Int64()
	return r
}
func (b ByteRate) Int64() int64 {
	v := (&big.Float{}).SetBytes(b.n, []byte{})
	v.Mul(v, big.NewFloat(float64(time.Second)/float64(b.d)))
	r, _ := v.Int64()
	return r
}
