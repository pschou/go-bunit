package bunit_test

import (
	"fmt"
	"time"

	"github.com/pschou/go-bunit"
)

func ExampleParseBits() {
	val, _ := bunit.ParseBits("32Kib")
	fmt.Println("32Kib =", val)

	// Notice that the Kb will used 1024, while Kib will use 1000
	val, _ = bunit.ParseBits("16kb")
	fmt.Println("16kb =", val)

	// Combined SI and base unit
	val, _ = bunit.ParseBits("16k100b")
	fmt.Println("16k100b =", val)
	// Output:
	// 32Kib = 32.768kb
	// 16kb = 16kb
	// 16k100b = 16.1kb
}

func ExampleParseBytes() {
	val, _ := bunit.ParseBytes("1kB")
	fmt.Println("1kB =", val)

	val, _ = bunit.ParseBytes("1KiB")
	fmt.Println("1KiB =", val)

	// Notice the lower case 'b' indicating bits
	val, _ = bunit.ParseBytes("1k200b")
	fmt.Println("1k200b =", val)
	// Output:
	// 1kB = 1kB
	// 1KiB = 1.024kB
	// 1k200b = 150B
}

func ExampleBits_format() {
	// Print out a unit in KiB format (1000):
	fmt.Printf("%%k = %k\n", bunit.NewBits(32768))

	// Print out a unit in kB format (1024):
	fmt.Printf("%%K = %K\n", bunit.NewBits(32768))

	// Print out a unit in MB format (1024^2):
	fmt.Printf("%%M = %M\n", bunit.NewBits(32768))

	// Print out a unit in Auto Byte format:
	fmt.Printf("%%0.5v = %0.5v\n", bunit.NewBits(13529101000))

	// Output:
	// %k = 32.768Kib
	// %K = 32kb
	// %M = 0.03125Mb
	// %0.5v = 13.529Gb
}

func ExampleBytes_format() {
	// Print out a unit in KiB format (1000):
	fmt.Printf("%%k = %k\n", bunit.NewBytes(32768))

	// Print out a unit in kB format (1024):
	fmt.Printf("%%K = %K\n", bunit.NewBytes(32768))

	// Print out a unit in MB format (1024^2):
	fmt.Printf("%%M = %M\n", bunit.NewBytes(32768))

	// Print out a unit in a string format (1000)
	fmt.Printf("%%s = %s\n", bunit.NewBytes(32768))

	// Print out a unit in a string format (1024)
	fmt.Printf("%%S = %S\n", bunit.NewBytes(32768))

	// Print out a unit in Auto Byte format:
	fmt.Printf("%%0.5v = %0.5v\n", bunit.NewBytes(13529000000))

	// Output:
	// %k = 32.768KiB
	// %K = 32kB
	// %M = 0.03125MB
	// %s = 32.768KiloByte
	// %S = 32KibiByte
	// %0.5v = 13.529GB
}

func ExampleByteRate_format() {
	// Print out a unit in KiB format (1000):
	fmt.Printf("%%k = %k\n", bunit.NewByteRate(32768, time.Second))

	// Print out a unit in kB format (1024):
	fmt.Printf("%%K = %K\n", bunit.NewByteRate(32768, time.Second))

	// Print out a unit in MB format (1024^2):
	fmt.Printf("%%M = %M\n", bunit.NewByteRate(3276800, time.Second))

	// Print out a unit in Auto Byte format:
	fmt.Printf("%%0.5v = %0.5v\n", bunit.NewByteRate(13529000000, time.Hour))

	// Output:
	// %k = 32.768KiB/s
	// %K = 32kB/s
	// %M = 3.125MB/s
	// %0.5v = 3.7581MB/s
}

func ExampleBitRate_format() {
	// Print out a unit in KiB format (1000):
	fmt.Printf("%%k = %k\n", bunit.NewBitRate(32768, time.Second))

	// Print out a unit in kB format (1024):
	fmt.Printf("%%K = %K\n", bunit.NewBitRate(32768, time.Second))

	// Print out a unit in MB format (1024^2):
	fmt.Printf("%%M = %M\n", bunit.NewBitRate(3276800, time.Second))

	// Print out a unit in Auto Byte format:
	fmt.Printf("%%0.5v = %0.5v\n", bunit.NewBitRate(13529000000, time.Hour))

	// Output:
	// %k = 32.768Kibps
	// %K = 32kbps
	// %M = 3.125Mbps
	// %0.5v = 3.7581Mbps
}

func ExampleParseBitRate() {
	// Rates can end with bps, b/s, or b/S and have the same meaning
	val, _ := bunit.ParseBitRate("1kbps")
	fmt.Println("1kbps =", val)

	val, _ = bunit.ParseBitRate("1kb/s")
	fmt.Println("1kb/s =", val)

	val, _ = bunit.ParseBitRate("1M400kb/s")
	fmt.Println("1M400kb/s =", val)

	val, _ = bunit.ParseBitRate("60kB/m")
	fmt.Println("60kB/m =", val)

	val, _ = bunit.ParseBitRate("1MB/s")
	fmt.Println("1MB/s =", val)

	val, _ = bunit.ParseBitRate("1KiB/s")
	fmt.Println("1KiB/s =", val)

	val, _ = bunit.ParseBitRate("1.544Mbps")
	fmt.Println("1.544 Mbps =", val)

	val, _ = bunit.ParseBitRate("6.312 Mbit/s")
	fmt.Println("6.312 Mbit/s =", val)

	val, _ = bunit.ParseBitRate("44.736 MBits/s")
	fmt.Println("44.736 MBits/s =", val)
	// Output:
	// 1kbps = 1kbps
	// 1kb/s = 1kbps
	// 1M400kb/s = 1.4Mbps
	// 60kB/m = 1kbps
	// 1MB/s = 1Mbps
	// 1KiB/s = 1.024kbps
	// 1.544 Mbps = 1.544Mbps
	// 6.312 Mbit/s = 6.312Mbps
	// 44.736 MBits/s = 44.736Mbps
}
