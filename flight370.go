package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type point struct{ lat, long float64 }

type region struct {
	radius float64
	point
}

type Placemark struct {
	Id    string `xml:"id,attr"`
	Name  string `xml:"name"`
	When  string `xml:"TimeStamp>when"`
	Where string `xml:"Point>coordinates"`
}

type placemark struct {
	id, name string
	t        time.Time
}

var lineRE = regexp.MustCompile(`([0-9]+); \((-?[0-9.]+), (-?[0-9.]+)\)`)

func readRegions(r *bufio.Reader) ([]region, error) {
	var regions []region
	for {
		peek, err := r.Peek(5)
		if err != nil {
			return nil, err
		}
		if string(peek) == "<?xml" {
			break
		}
		line, err := r.ReadString('\n')
		if err != nil {
			return nil, err
		}
		r, err := parseRegion(line)
		if err != nil {
			log.Fatalf("%s parsing region %q", err, line)
			return nil, err
		}
		regions = append(regions, r)
	}
	return regions, nil
}

func parseRegion(line string) (region, error) {
	var reply region
	m := lineRE.FindStringSubmatch(line)
	if len(m) != 4 {
		return reply, fmt.Errorf("too few matches")
	}
	r, err := strconv.ParseFloat(m[1], 64)
	if err != nil {
		return reply, err
	}
	lat, err := strconv.ParseFloat(m[2], 64)
	if err != nil {
		return reply, err
	}
	long, err := strconv.ParseFloat(m[3], 64)
	if err != nil {
		return reply, err
	}
	return region{r, point{lat, long}}, nil
}

func parsePoint(s string) point {
	chunks := strings.Split(s, ",")
	lat, err := strconv.ParseFloat(chunks[0], 64)
	if err != nil {
		log.Fatalf("%s parsing float in %q", err, s)
	}
	long, err := strconv.ParseFloat(chunks[1], 64)
	if err != nil {
		log.Fatalf("%s parsing float in %q", err, s)
	}
	return point{lat, long}
}

const R = 6371 // km

// distance calculates great circle distance using haversine formula
// from http://www.movable-type.co.uk/scripts/latlong.html
func (p point) distance(q point) float64 {
	φ1 := radians(p.lat)
	φ2 := radians(q.lat)
	Δφ := φ1 - φ2
	Δλ := radians(p.long - q.long)
	a := sqr(math.Sin(Δφ/2)) + math.Cos(φ1)*math.Cos(φ2)*sqr(math.Sin(Δλ/2))

	return R * 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
}
func sqr(a float64) float64           { return a * a }
func radians(degrees float64) float64 { return degrees / 180 * math.Pi }

func (r region) contains(p point) bool {
	return r.distance(p) <= r.radius
}

func readXML(d *xml.Decoder, regions []region) []map[*placemark]int {
	hits := make([]map[*placemark]int, len(regions))
	for k := range regions {
		hits[k] = make(map[*placemark]int)
	}
	for {
		t, err := d.Token()
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		tok, ok := t.(xml.StartElement)
		if !(ok && tok.Name.Local == "Placemark") {
			continue
		}
		var mark Placemark
		d.DecodeElement(&mark, &tok)

		if len(mark.Where) <3 || len(mark.Id) < 1 || len(mark.Where) < 3 {
			continue
		}
		p := parsePoint(mark.Where)
		var pm *placemark
		for k, r := range regions {
			if !r.contains(p) {
				continue
			}
			if pm == nil {
				pm = &placemark{mark.Id, mark.Name, parseTime(mark.When)}
			}
			hits[k][pm]++
		}
	}
	return hits
}

func parseTime(s string) time.Time {
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

func report(counts map[*placemark]int) string {
	if len(counts) == 0 {
		return "None"
	}
	maxCount := 0
	var maxMarks []*placemark
	for p, c := range counts {
		switch {
		case c > maxCount:
			maxCount = c
			maxMarks = []*placemark{p}
		case c == maxCount:
			maxMarks = append(maxMarks, p)
		}
	}
	sort.Sort(BySpecial(maxMarks))
	var names []string
	for _, m := range maxMarks {
		names = append(names, m.name)
	}
	return strings.Join(names, ", ")
}

type BySpecial []*placemark

func (a BySpecial) Len() int      { return len(a) }
func (a BySpecial) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a BySpecial) Less(i, j int) bool {
	if a[i].t.Before(a[j].t) {
		return true
	}
	if a[i].t.After(a[j].t) {
		return false
	}
	return a[i].id < a[j].id
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("expected filename")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)
	regions, err := readRegions(r)
	if err != nil {
		log.Fatal(err, "reading regions")
	}
	hits := readXML(xml.NewDecoder(r), regions)
	for _, m := range hits {
		fmt.Println(report(m))
	}
}
