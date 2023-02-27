package bunit_test

import (
	"fmt"
	"time"

	"github.com/pschou/go-bunit"
)

func ExampleParseBits() {
	val, _ := bunit.ParseBits("32Kib")
	fmt.Printf("32Kb = %k\n", val)

	// Notice that the Kb will used 1024, while Kib will use 1000
	val, _ = bunit.ParseBits("16Kb")
	fmt.Printf("16Kb = %k\n", val)
	// Output:
	// 32Kb = 32Kib
	// 16Kb = 16.384Kib
}

func ExampleParseBytes() {
	val, _ := bunit.ParseBytes("1KB")
	fmt.Println("1KB =", val)

	val, _ = bunit.ParseBytes("1KiB")
	fmt.Println("1KiB =", val)

	// Notice the lower case 'b' indicating bits
	val, _ = bunit.ParseBytes("1Kb")
	fmt.Println("1Kb =", val)
	// Output:
	// 1KB = 1024B
	// 1KiB = 1000B
	// 1Kb = 128B
}

func ExampleBits_format() {
	// Print out a unit in KiB format (1000):
	fmt.Printf("%%k = %k\n", bunit.NewBits(32768))

	// Print out a unit in KB format (1024):
	fmt.Printf("%%K = %K\n", bunit.NewBits(32768))

	// Print out a unit in MB format (1024^2):
	fmt.Printf("%%M = %M\n", bunit.NewBits(32768))

	// Print out a unit in Auto Byte format:
	fmt.Printf("%%0.5A = %0.5A\n", bunit.NewBits(13529000000))

	// Output:
	// %k = 32.768Kib
	// %K = 32Kb
	// %M = 0.03125Mb
	// %0.5A = 12.6Gb
}

func ExampleBytes_format() {
	// Print out a unit in KiB format (1000):
	fmt.Printf("%%k = %k\n", bunit.NewBytes(32768))

	// Print out a unit in KB format (1024):
	fmt.Printf("%%K = %K\n", bunit.NewBytes(32768))

	// Print out a unit in MB format (1024^2):
	fmt.Printf("%%M = %M\n", bunit.NewBytes(32768))

	// Print out a unit in Auto Byte format:
	fmt.Printf("%%0.5A = %0.5A\n", bunit.NewBytes(13529000000))

	// Output:
	// %k = 32.768KiB
	// %K = 32KB
	// %M = 0.03125MB
	// %0.5A = 12.6GB
}

func ExampleByteRate_format() {
	// Print out a unit in KiB format (1000):
	fmt.Printf("%%k = %k\n", bunit.NewByteRate(32768, time.Second))

	// Print out a unit in KB format (1024):
	fmt.Printf("%%K = %K\n", bunit.NewByteRate(32768, time.Second))

	// Print out a unit in MB format (1024^2):
	fmt.Printf("%%M = %M\n", bunit.NewByteRate(3276800, time.Second))

	// Print out a unit in Auto Byte format:
	fmt.Printf("%%0.5A = %0.5A\n", bunit.NewByteRate(13529000000, time.Hour))

	// Output:
	// %k = 32.768KiB/s
	// %K = 32KB/s
	// %M = 3.125MB/s
	// %0.5A = 3.584MB/s
}

func ExampleBitRate_format() {
	// Print out a unit in KiB format (1000):
	fmt.Printf("%%k = %k\n", bunit.NewBitRate(32768, time.Second))

	// Print out a unit in KB format (1024):
	fmt.Printf("%%K = %K\n", bunit.NewBitRate(32768, time.Second))

	// Print out a unit in MB format (1024^2):
	fmt.Printf("%%M = %M\n", bunit.NewBitRate(3276800, time.Second))

	// Print out a unit in Auto Byte format:
	fmt.Printf("%%0.5A = %0.5A\n", bunit.NewBitRate(13529000000, time.Hour))

	// Output:
	// %k = 32.768Kibps
	// %K = 32Kbps
	// %M = 3.125Mbps
	// %0.5A = 3.584Mbps
}

/*
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
*/
func ExampleParseBitRate() {
	// Rates can end with bps, b/s, or b/S and have the same meaning
	val, _ := bunit.ParseBitRate("1kbps")
	fmt.Println("1kbps =", val)

	val, _ = bunit.ParseBitRate("1kb/s")
	fmt.Println("1kb/s =", val)

	val, _ = bunit.ParseBitRate("1Kb/s")
	fmt.Println("1Kb/s =", val)

	val, _ = bunit.ParseBitRate("1kB/s")
	fmt.Println("1kB/s =", val)

	val, _ = bunit.ParseBitRate("1KB/s")
	fmt.Println("1KB/s =", val)

	val, _ = bunit.ParseBitRate("1KiB/s")
	fmt.Println("1KiB/s =", val)

	val, _ = bunit.ParseBitRate("1.544mbps")
	fmt.Printf("1.544 mbps = %a\n", val)

	val, _ = bunit.ParseBitRate("6.312 mbit/s")
	fmt.Printf("6.312 mbit/s = %a\n", val)

	val, _ = bunit.ParseBitRate("44.736 MiBits/m")
	fmt.Printf("44.736 MiBits/m = %a\n", val)
	// Output:
	// 1kbps = 1000bps
	// 1kb/s = 1000bps
	// 1Kb/s = 1024bps
	// 1kB/s = 8000bps
	// 1KB/s = 8192bps
	// 1KiB/s = 8000bps
	// 1.544 mbps = 1.544Mibps
	// 6.312 mbit/s = 6.312Mibps
	// 44.736 MiBits/m = 0.7456Mibps
}
