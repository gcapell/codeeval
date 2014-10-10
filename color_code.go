package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type rgb struct{ r, g, b int }

func (r rgb) String() string { return fmt.Sprintf("RGB(%d,%d,%d)", r.r, r.g, r.b) }

type parser func(string) rgb

var parsers = map[string]parser{
	"HSL": hsl,
	"HSV": hsv,
	"(":   cmyk,
	"#":   hex,
}

func convert(line string) rgb {
	for prefix, p := range parsers {
		if strings.HasPrefix(line, prefix) {
			return p(line)
		}
	}
	panic(line)
}

func hue2rgb(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	switch {
	case t < 1.0/6:
		return p + (q-p)*6*t
	case t < 1.0/2:
		return q
	case t < 1.0/3:
		return p + (q-p)*(2/3-t)*6
	}
	return p
}

func hsl(line string) rgb {
	var hi, si, li int
	fmt.Sscanf(line, "HSL(%d,%d,%d)", &hi, &si, &li)

	h := float64(hi) / 360
	s := float64(si) / 100
	l := float64(li) / 100

	if s == 0 {
		return fromFloat(l, l, l)
	}

	var q float64
	if l < 0.5 {
		q = l * (1 + s)
	} else {
		q = l + s - l*s
	}
	p := 2*l - q

	return fromFloat(
		hue2rgb(p, q, h+1/3),
		hue2rgb(p, q, h),
		hue2rgb(p, q, h-1/3),
	)
}

func r255(f float64) int {
	return int(math.Floor(f*255 + 0.5))
}

func fromFloat(r, g, b float64) rgb {
	return rgb{
		r255(r), r255(g), r255(b),
	}
}

func hsv(line string) rgb {
	var hi, si, vi int
	fmt.Sscanf(line, "HSV(%d,%d,%d)", &hi, &si, &vi)

	h := float64(hi) / 360
	s := float64(si) / 100
	v := float64(vi) / 100

	if si == 0 {
		return fromFloat(v, v, v)
	}

	sector := int(math.Floor(h * 6))
	f := h*6 - float64(sector)
	p := v * (1 - s)
	q := v * (1 - s*f)
	t := v * (1 - s*(1-f))

	var r, g, b float64
	switch sector {
	case 0:
		r, g, b = v, t, p
	case 1:
		r, g, b = q, v, p
	case 2:
		r, g, b = p, v, t
	case 3:
		r, g, b = p, q, v
	case 4:
		r, g, b = t, p, v
	case 5:
		r, g, b = v, p, q
	default:
		panic("eek")
	}
	return fromFloat(r, g, b)
}

func cmyk(s string) rgb {
	var c, m, y, k float64
	fmt.Sscanf(s, "(%f,%f,%f,%f)", &c, &m, &y, &k)

	return fromFloat(
		(1-c)*(1-k),
		(1-m)*(1-k),
		(1-y)*(1-k),
	)
}

func htoi(s string) int {
	n, err := strconv.ParseInt(s, 16, 32)
	if err != nil {
		panic(err)
	}
	return int(n)
}

func hex(s string) rgb {
	return rgb{htoi(s[1:3]), htoi(s[3:5]), htoi(s[5:7])}
}

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
		fmt.Println(convert(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
