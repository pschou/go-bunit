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
)

func ParseRate(s string) (float64, error) {
	orig := s
	var d float64
	neg := false

	// Consume [-+]?
	if s != "" {
		c := s[0]
		if c == '-' || c == '+' {
			neg = c == '-'
			s = s[1:]
		}
	}
	// Special case: if all that is left is "0", this is zero.
	if s == "0" {
		return 0, nil
	}
	if s == "" {
		return 0, errors.New("binary unit: invalid value " + quote(orig))
	}
	for s != "" {
		var (
			v, f  uint64      // integers before, after decimal point
			scale float64 = 1 // value = v + f/scale
		)

		var err error

		// The next character must be [0-9.]
		if !(s[0] == '.' || '0' <= s[0] && s[0] <= '9') {
			return 0, errors.New("binary unit: invalid value " + quote(orig))
		}
		// Consume [0-9]*
		pl := len(s)
		v, s, err = leadingInt(s)
		if err != nil {
			return 0, errors.New("binary unit: invalid value " + quote(orig))
		}
		pre := pl != len(s) // whether we consumed anything before a period

		// Consume (\.[0-9]*)?
		post := false
		if s != "" && s[0] == '.' {
			s = s[1:]
			pl := len(s)
			f, scale, s = leadingFraction(s)
			post = pl != len(s)
		}
		if !pre && !post {
			// no digits (e.g. ".s" or "-.s")
			return 0, errors.New("binary unit: invalid value " + quote(orig))
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
		if i < 3 {
			return 0, errors.New("binary unit: missing rate unit in value " + quote(orig))
		}
		switch s[b:] {
		case "bit/s", "bps", "b/s", "b/S", "Bit/s", "bits/s", "Bits/s":
		case "Bps", "B/s", "B/S", "bytes/s", "Bytes/s":
			v = v << 3
			f *= 8
		default:
			return 0, errors.New("binary unit: missing byte or bit unit in value " + quote(orig))
		}
		u := s[:b]
		s = s[i:]
		unit, ok := unitMap[u]
		if !ok {
			return 0, errors.New("binary unit: unknown unit " + quote(u) + " in value " + quote(orig))
		}
		d += float64(v) * unit
		if f > 0 {
			d += float64(f) * (unit / scale)
			if v > 1<<63 {
				// overflow
				return 0, errors.New("binary unit: invalid value " + quote(orig))
			}
		}
	}
	if neg {
		return -d, nil
	}
	if d > 1<<63-1 {
		return 0, errors.New("binary unit: invalid value " + quote(orig))
	}
	return d, nil
}
