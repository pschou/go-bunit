package bunit

import (
	"errors"
	"fmt"
)

// Format the number of bits into a given unit for display
func Format(val float64, formatUnit string) (string, error) {
	format, s := doPrintfSplit(formatUnit)
	f := float64(1)

	// Consume unit.
	i, b := 0, 0
	for ; i < len(s); i++ {
		c := s[i]
		if c == 'b' || c == 'B' {
			b = i
		}
	}
	if i == 0 {
		return "", errors.New("binary unit: missing unit in format " + quote(formatUnit))
	}
	switch s[b:] {
	case "b", "bit", "Bit", "bits", "Bits":
	case "B", "byte", "Byte", "bytes", "Bytes":
		f *= 8
	default:
		return "", errors.New("binary unit: missing byte or bit unit in format " + quote(formatUnit))
	}
	u := s[:b]
	unit, ok := unitMap[u]
	if !ok {
		return "", errors.New("binary unit: unknown unit " + quote(u) + " in format " + quote(formatUnit))
	}

	return fmt.Sprintf(format, val/(f*unit)) + s, nil
}
