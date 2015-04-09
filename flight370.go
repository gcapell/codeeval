package main

import (
	"bufio"
	"encoding/XML"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type region struct {
	id        string
	lat, long float64
}

var lineRE = regexp.MustCompile(`([0-9]+); \(([0-9.]+), ([0-9.]+)\)`)

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
		m := lineRE.FindStringSubmatch(line)
		lat, err := strconv.ParseFloat(m[2], 64)
		if err != nil {
			return nil, err
		}
		lon, err := strconv.ParseFloat(m[3], 64)
		if err != nil {
			return nil, err
		}
		regions = append(regions, region{m[1], lat, lon})
	}
	return regions, nil
}

type Placemark struct {
	Id string `xml:"id,attr"`
	Name string `xml:"name"`
	When string `xml:"TimeStamp>when"`
	Where string `xml:"Point>coordinates"`
	
}

func readXML(d *xml.Decoder) {
	for {
		t, err := d.Token()
		if err != nil {
			log.Println("err", err)
			break
		}
		tok, ok := t.(xml.StartElement)
		if !(ok && tok.Name.Local == "Placemark"){
			continue
		}
		var p Placemark
		d.DecodeElement(&p, &tok)
		fmt.Printf("%#v\n", p)
	}
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
	fmt.Println("regions", regions)
	readXML(xml.NewDecoder(r))

}
