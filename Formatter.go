package bunit

import (
	"fmt"
	"time"

	"github.com/cymertek/go-big"
)

func (b Bits) Format(f fmt.State, verb rune) {
	formatByte(b, 1, f, verb, 'b', "b")
}
func (b BitRate) Format(f fmt.State, verb rune) {
	formatByte(b.n, float64(time.Second)/float64(b.d), f, verb, 'b', "bps")
}

func (b Bytes) Format(f fmt.State, verb rune) {
	formatByte(b, 1, f, verb, 'B', "B")
}
func (b ByteRate) Format(f fmt.State, verb rune) {
	formatByte(b.n, float64(time.Second)/float64(b.d), f, verb, 'B', "B/s")
}

func formatByte(b []byte, scale float64, f fmt.State, verb, def rune, suf string) {
	v := (&big.Float{}).SetBytes(b, []byte{})
	if scale != 1 {
		v = v.Mul(v, big.NewFloat(scale))
	}
	switch verb {
	case def:
	case 'a':
		// Auto with 1000 multiples
		n, _ := v.Int(nil)
		n = n.Lsh(n, 2)
		if (&big.Int{}).SetBytes(thousand[10]).Cmp(n) <= 0 {
			n = n.Rsh(n, 10)
			for i, c := range thousandVerb[:10] {
				if (&big.Int{}).SetBytes(thousand[10+i]).Cmp(n) > 0 {
					v.Quo(v, (&big.Float{}).SetBytes(thousand[10+i], []byte{}))
					suf = string(c) + "i" + suf
					break
				}
			}
		}
	case 'A':
		// Auto with 1024 multiples
		n, _ := v.Int(nil)
		n = n.Lsh(n, 2)
		if (&big.Int{}).SetBytes(thousand[0]).Cmp(n) <= 0 {
			n = n.Rsh(n, 10)
			for i, c := range thousandVerb[:10] {
				if (&big.Int{}).SetBytes(thousand[i]).Cmp(n) > 0 {
					v.Quo(v, (&big.Float{}).SetBytes(thousand[i], []byte{}))
					suf = string(c) + suf
					break
				}
			}
		}
	default:
		// All the SI Byte units
		for i, c := range thousandVerb {
			//fmt.Printf("comparing %q %q %v\n", verb, c, thousand[i])
			if c == verb {
				//fmt.Printf("%v / %v\n", v, (&big.Float{}).SetBytes(thousand[i+1], []byte{}))
				v.Quo(v, (&big.Float{}).SetBytes(thousand[i], []byte{}))
				if i < 10 {
					suf = string(verb) + suf
				} else {
					suf = string(thousandVerb[i%10]) + "i" + suf
				}
				break
			}
		}
		if suf == "" {
			f.Write([]byte("%!(INVALID " + quote(string(verb)) + ")"))
			return
		}
	}
	v.Format(f, 'g')
	f.Write([]byte(suf))
}
