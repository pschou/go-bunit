package bunit_test

import (
	"fmt"

	"github.com/pschou/go-bunit"
)

func ExampleParse() {
	val, _ := bunit.Parse("1kb")
	fmt.Println("1kb =", val)

	val, _ = bunit.Parse("1Kb")
	fmt.Println("1Kb =", val)

	val, _ = bunit.Parse("1kB")
	fmt.Println("1kB =", val)

	val, _ = bunit.Parse("1KB")
	fmt.Println("1KB =", val)

	val, _ = bunit.Parse("1KiB")
	fmt.Println("1KiB =", val)
	// Output:
	// 1kb = 1000
	// 1Kb = 1024
	// 1kB = 8000
	// 1KB = 8192
	// 1KiB = 8192
}

func ExampleFormatRate() {
	val, _ := bunit.FormatRate(8192000, "Mbps")
	fmt.Println("8192000 =", val)
	// Output:
	// 8192000 = 8.192Mbps
}
func ExampleFormat() {
	val, _ := bunit.Format(1000, "%v kb")
	fmt.Println("1kb =", val)

	val, _ = bunit.Format(1024, "Kb")
	fmt.Println("1Kb =", val)

	val, _ = bunit.Format(8192, "%0.02f KiB")
	fmt.Println("1KiB =", val)
	// Output:
	// 1kb = 1 kb
	// 1Kb = 1Kb
	// 1KiB = 1.00 KiB
}

func ExampleParseRate() {
	// Rates can end with bps, b/s, or b/S and have the same meaning
	val, _ := bunit.ParseRate("1kbps")
	fmt.Println("kbps =", val)

	val, _ = bunit.ParseRate("1kb/s")
	fmt.Println("1kb/s =", val)

	val, _ = bunit.ParseRate("1Kb/s")
	fmt.Println("1Kb/s =", val)

	val, _ = bunit.ParseRate("1kB/s")
	fmt.Println("1kB/s =", val)

	val, _ = bunit.ParseRate("1KB/s")
	fmt.Println("1KB/s =", val)

	val, _ = bunit.ParseRate("1KiB/s")
	fmt.Println("1KiB/s =", val)

	val, _ = bunit.ParseRate("1.544 Mbps")
	fmt.Println("1.544 Mbps =", val)

	val, _ = bunit.ParseRate("6.312 megabit/s")
	fmt.Println("6.312 megabit/s =", val)

	val, _ = bunit.ParseRate("44.736 MegaBits/s")
	fmt.Println("44.736 MegaBits/s =", val)
	// Output:
	// kbps = 1000
	// 1kb/s = 1000
	// 1Kb/s = 1024
	// 1kB/s = 8000
	// 1KB/s = 8192
	// 1KiB/s = 8192
	// 1.544 Mbps = 1.544e+06
	// 6.312 megabit/s = 6.312e+06
	// 44.736 MegaBits/s = 4.4736e+07
}
