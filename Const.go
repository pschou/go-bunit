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
	"k":      1e3,
	"Kilo":   1e3,
	"kilo":   1e3,
	"M":      1e6,
	"Mega":   1e6,
	"mega":   1e6,
	"G":      1e9,
	"Giga":   1e9,
	"giga":   1e9,
	"T":      1e12,
	"Tera":   1e12,
	"tera":   1e12,
	"P":      1e15,
	"Peta":   1e15,
	"peta":   1e15,
	"E":      1e18,
	"Exa":    1e18,
	"exa":    1e18,
	"Z":      1e21,
	"Zetta":  1e21,
	"zetta":  1e21,
	"Y":      1e24,
	"Yotta":  1e24,
	"yotta":  1e24,
	"R":      1e27,
	"Ronna":  1e27,
	"ronna":  1e27,
	"Q":      1e30,
	"Quetta": 1e30,
	"quetta": 1e30,

	"K":    1 << 10,
	"Ki":   1 << 10,
	"Kibi": 1 << 10,
	"kibi": 1 << 10,
	"Mi":   1 << 20,
	"Mebi": 1 << 20,
	"mebi": 1 << 20,
	"Gi":   1 << 30,
	"Gibi": 1 << 30,
	"gibi": 1 << 30,
	"Ti":   1 << 40,
	"Tebi": 1 << 40,
	"tebi": 1 << 40,
	"Pi":   1 << 50,
	"Pebi": 1 << 50,
	"pebi": 1 << 50,
	"Ei":   1 << 60,
	"Exbi": 1 << 60,
	"exbi": 1 << 60,
	"Zi":   1 << 70,
	"Zebi": 1 << 70,
	"zebi": 1 << 70,
	"Yi":   1 << 80,
	"Yobi": 1 << 80,
	"yobi": 1 << 80,
	"Ri":   1 << 90,
	"Robi": 1 << 90,
	"robi": 1 << 90,
	"Qi":   1 << 100,
	"Qubi": 1 << 100,
	"qubi": 1 << 100,
}
