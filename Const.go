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
	"fmt"
	"time"
)

type Bytes []byte
type Bits []byte
type ByteRate struct {
	n Bytes
	d time.Duration
}
type BitRate struct {
	n Bits
	d time.Duration
}

var byteZero Bytes
var bitZero Bits
var byteRateZero ByteRate
var bitRateZero BitRate
var _ fmt.Formatter = byteZero     // Bytes must implement fmt.Formatter
var _ fmt.Formatter = bitZero      // Bits must implement fmt.Formatter
var _ fmt.Formatter = byteRateZero // ByteRate must implement fmt.Formatter
var _ fmt.Formatter = bitRateZero  // BitRate must implement fmt.Formatter

const (
	// Decimal binary values
	Byte       = float64(1)
	KiloByte   = float64(1e3)
	MegaByte   = float64(1e6)
	GigaByte   = float64(1e9)
	TeraByte   = float64(1e12)
	PetaByte   = float64(1e15)
	ExaByte    = float64(1e18)
	ZettaByte  = float64(1e21)
	YottaByte  = float64(1e24)
	RonnaByte  = float64(1e27)
	QuettaByte = float64(1e30)

	// IEC Binary Values
	KibiByte = float64(8 << (iota * 3))
	MebiByte
	GibiByte
	TebiByte
	PebiByte
	ExbiByte
	ZebiByte
	YobiByte
	RobiByte
	QubiByte
)

var unitMap = map[string]float64{
	"":       1,
	"K":      1 << 10,
	"Kilo":   1 << 10,
	"kilo":   1 << 10,
	"M":      1 << 20,
	"Mega":   1 << 20,
	"mega":   1 << 20,
	"G":      1 << 30,
	"Giga":   1 << 30,
	"giga":   1 << 30,
	"T":      1 << 40,
	"Tera":   1 << 40,
	"tera":   1 << 40,
	"P":      1 << 50,
	"Peta":   1 << 50,
	"peta":   1 << 50,
	"E":      1 << 60,
	"Exa":    1 << 60,
	"exa":    1 << 60,
	"Z":      1 << 70,
	"Zetta":  1 << 70,
	"zetta":  1 << 70,
	"Y":      1 << 80,
	"Yotta":  1 << 80,
	"yotta":  1 << 80,
	"R":      1 << 90,
	"Ronna":  1 << 90,
	"ronna":  1 << 90,
	"Q":      1 << 100,
	"Quetta": 1 << 100,
	"quetta": 1 << 100,

	"k":    1e3,
	"Ki":   1e3,
	"Kibi": 1e3,
	"kibi": 1e3,
	"m":    1e6,
	"Mi":   1e6,
	"Mebi": 1e6,
	"mebi": 1e6,
	"g":    1e9,
	"Gi":   1e9,
	"Gibi": 1e9,
	"gibi": 1e9,
	"t":    1e12,
	"Ti":   1e12,
	"Tebi": 1e12,
	"tebi": 1e12,
	"p":    1e15,
	"Pi":   1e15,
	"Pebi": 1e15,
	"pebi": 1e15,
	"e":    1e18,
	"Ei":   1e18,
	"Exbi": 1e18,
	"exbi": 1e18,
	"z":    1e21,
	"Zi":   1e21,
	"Zebi": 1e21,
	"zebi": 1e21,
	"y":    1e24,
	"Yi":   1e24,
	"Yobi": 1e24,
	"yobi": 1e24,
	"r":    1e27,
	"Ri":   1e27,
	"Robi": 1e27,
	"robi": 1e27,
	"q":    1e30,
	"Qi":   1e30,
	"Qubi": 1e30,
	"qubi": 1e30,
}

var thousandVerb = "KMGTPEZYRQkmgtpezyrq"
var thousand = [][]byte{
	[]byte{0x4, 0x0},
	[]byte{0x10, 0x0, 0x0},
	[]byte{0x40, 0x0, 0x0, 0x0},
	[]byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0},
	[]byte{0x4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
	[]byte{0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
	[]byte{0x40, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
	[]byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
	[]byte{0x4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
	[]byte{0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
	[]byte{0x3, 0xe8},
	[]byte{0xf, 0x42, 0x40},
	[]byte{0x3b, 0x9a, 0xca, 0x0},
	[]byte{0xe8, 0xd4, 0xa5, 0x10, 0x0},
	[]byte{0x3, 0x8d, 0x7e, 0xa4, 0xc6, 0x80, 0x0},
	[]byte{0xd, 0xe0, 0xb6, 0xb3, 0xa7, 0x64, 0x0, 0x0},
	[]byte{0x36, 0x35, 0xc9, 0xad, 0xc5, 0xde, 0xa0, 0x0, 0x0},
	[]byte{0xd3, 0xc2, 0x1b, 0xce, 0xcc, 0xed, 0xa1, 0x0, 0x0, 0x0},
	[]byte{0x3, 0x3b, 0x2e, 0x3c, 0x9f, 0xd0, 0x80, 0x3c, 0xe8, 0x0, 0x0, 0x0},
	[]byte{0xc, 0x9f, 0x2c, 0x9c, 0xd0, 0x46, 0x74, 0xed, 0xea, 0x40, 0x0, 0x0, 0x0},
}
