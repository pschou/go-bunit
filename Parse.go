// Copyright 2023 github.com/pschou
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bunit

import (
	"errors"
	"log"

	"github.com/cymertek/go-big"
)

func MustParseBytes(s string) Bytes {
	b, err := ParseBytes(s)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func ParseBytes(s string) (Bytes, error) {
	b, err := ParseBits(s)
	if err == nil {
		i := (&big.Int{}).SetBytes(b)
		i.Rsh(i, 3)
		return Bytes(i.Bytes()), nil
	}
	return nil, err
}

func MustParseBits(s string) Bits {
	b, err := ParseBits(s)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

var eight = big.NewFloat(8)

func ParseBits(s string) (Bits, error) {
	d := &big.Float{}
	orig := s

	// Consume [-+]?
	if s != "" {
		switch s[0] {
		case '-':
			return nil, errors.New("binary unit: invalid value " + quote(orig))
		case '+':
			s = s[1:]
		}
	}
	// Special case: if all that is left is "0", this is zero.
	if s == "0" {
		return []byte{0}, nil
	}
	if s == "" {
		return nil, errors.New("binary unit: invalid value " + quote(orig))
	}
	for s != "" {
		var v *big.Float
		var err error

		// The next character must be [0-9.]
		if !(s[0] == '.' || '0' <= s[0] && s[0] <= '9') {
			return nil, errors.New("binary unit: invalid value " + quote(orig))
		}
		// Consume [0-9.]*
		v, s, err = leadingBigFloat(s)
		if err != nil {
			return nil, errors.New("binary unit: invalid value " + quote(orig))
		}

		// Get rid of spaces
		for len(s) > 0 && s[0] == ' ' {
			s = s[1:]
		}

		// Consume unit.
		i, b := 0, 0
		for ; i < len(s); i++ {
			c := s[i]
			if c == 'b' || c == 'B' {
				b = i
			}
			if c == '.' || '0' <= c && c <= '9' {
				break
			}
		}
		if i == 0 {
			return nil, errors.New("binary unit: missing unit in value " + quote(orig))
		}
		switch s[b:i] {
		case "b", "bit", "Bit", "bits", "Bits":
		case "B", "byte", "Byte", "bytes", "Bytes":
			v.Mul(v, eight)
		default:
			return nil, errors.New("binary unit: missing byte or bit unit in value " + quote(orig))
		}
		u := s[:b]
		s = s[i:]
		unit, ok := unitMap[u]
		if !ok {
			return nil, errors.New("binary unit: unknown unit " + quote(u) + " in value " + quote(orig))
		}
		v.Mul(v, big.NewFloat(unit))
		d.Add(d, v)
	}
	b, _ := d.Bytes()
	return Bits(b), nil
}
