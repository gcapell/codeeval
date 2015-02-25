package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

type ip uint32

func (a ip) String() string {
	var reply [4]string
	for j := 3; j >= 0; j-- {
		reply[j] = strconv.Itoa(int(a & 0xff))
		a >>= 8
	}
	return strings.Join(reply[:], ".")
}

type converter struct {
	parse func(string) (ip, bool)
	re    *regexp.Regexp
}
type converterDesc struct {
	pattern string
	parse   func(string) (ip, bool)
}

var converterDescs = []converterDesc{
	// dottedDecimal: 192.0.2.235 with no leading zero.
	{`[1-9]\d*\.\d+\.\d+\.\d+`, dottedDecimal},

	// dottedHex: 0xc0.0x0.0x02.0xeb Each octet is individually converted to hexadecimal form.
	{`0x[0-9a-f]{1,2}\.0x[0-9a-f]{1,2}\.0x[0-9a-f]{1,2}\.0x[0-9a-f]{1,2}`, dottedOct},

	// Dotted octal 0300.0000.0002.0353 Each octet is individually converted into octal.
	{`0[0-7]{3}\.0[0-7]{3}\.0[0-7]{3}\.0[0-7]{3}`, dottedHex},

	// Dotted binary 11000000.00000000.00000010.11101011 Each octet is individually converted into binary.
	{`[01]{8}\.[01]{8}\.[01]{8}\.[01]{8}`, dottedBinary},

	// Binary 11000000000000000000001011101011
	{`[01]{32}`, binary},

	// Octal 030000001353
	{`[01]{12}`, octal},

	// Hexadecimal	0xC00002EB	Concatenation of the octets from the dotted hexadecimal.
	{`0x[0-9A-Fa-f]{8}`, hexadecimal},

	// Decimal	3221226219	The 32-bit number expressed in decimal.
	{`[1-9][0-9]+`, decimal},
}

var converters []converter

func init() {
	for _, c := range converterDescs {
		converters = append(converters, converter{c.parse, regexp.MustCompile(c.pattern)})
	}
}

var count = make(map[ip]int)

func scan(line string) {
	for _, c := range converters {
		for _, i := range c.re.FindAllStringIndex(line, -1) {
			s := line[i[0]:i[1]]
			if addr, ok := c.parse(s); ok {
				count[addr]++
			}
		}
	}
}

type ByAddr []ip

func (a ByAddr) Len() int           { return len(a) }
func (a ByAddr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAddr) Less(i, j int) bool { return a[i] < a[j] }

func report() string {
	var maxAddr []ip
	maxCount := 0

	for addr, c := range count {
		switch {
		case c == maxCount:
			maxAddr = append(maxAddr, addr)
		case c > maxCount:
			maxCount = c
			maxAddr = []ip{addr}
		}
	}
	sort.Sort(ByAddr(maxAddr))
	reply := make([]string, 0, len(maxAddr))
	for _, a := range maxAddr {
		reply = append(reply, a.String())
	}
	return strings.Join(reply, " ")
}

func dottedParse(s string, base int) (ip, bool) {
	chunks := strings.Split(s, ".")
	var reply ip
	for _, c := range chunks {
		n, err := strconv.ParseUint(c, base, 8)
		if err != nil {
			// log.Printf("%s parsing %q (base %d) in %q", err, c, base, s)
			return 0, false
		}
		reply = (reply << 8) + ip(n)
	}
	return reply, true
}

func dottedDecimal(s string) (ip, bool) { return dottedParse(s, 10) }
func dottedOct(s string) (ip, bool)     { return dottedParse(s, 0) }
func dottedHex(s string) (ip, bool)     { return dottedParse(s, 0) }
func dottedBinary(s string) (ip, bool)  { return dottedParse(s, 2) }

func parse(s string, base int) (ip, bool) {
	n, err := strconv.ParseUint(s, base, 32)
	if err != nil {
		// log.Printf("%s parsing %q (base %d)", err, s, base)
		return 0, false
	}
	if n > 0xffffffff || n < 0x01000000 {
		return 0, false
	}
	return ip(n), true
}

func binary(s string) (ip, bool)      { return parse(s, 2) }
func octal(s string) (ip, bool)       { return parse(s, 8) }
func hexadecimal(s string) (ip, bool) { return parse(s, 0) }
func decimal(s string) (ip, bool)     { return parse(s, 10) }

func main() {
	if len(os.Args) != 2 {
		log.Fatal("expected filename")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		scan(scanner.Text())
		runtime.GC()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(report())
}
