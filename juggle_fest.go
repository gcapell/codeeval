// Solve juggle fest (stable marriage) using Gale-Shapley algorithm
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var jugglersPerCircuit int

type circuit struct {
	h, e, p  int
	jugglers map[*juggler]int
	minScore *juggler
}

func (c circuit) String() string {
	var jugglers []string
	for j, score := range c.jugglers {
		jugglers = append(jugglers, fmt.Sprintf("%s:%d", j.name, score))
	}
	return strings.Join(jugglers, ",")
}

var circuits = make(map[string]*circuit)

func addCircuit(line string) {
	f := strings.Fields(line)
	circuits[f[1]] = &circuit{
		h:        atoi(f[2][2:]),
		e:        atoi(f[3][2:]),
		p:        atoi(f[4][2:]),
		jugglers: make(map[*juggler]int),
	}
}

func countAssigned()int{
	var total int
	for _, c := range circuits{
		total += len(c.jugglers)
	}
	return total
}

func (c *circuit) propose(j *juggler) (accept bool, displace *juggler) {
	score := j.h*c.h + j.e*c.e + j.p*c.p
	if len(c.jugglers) < jugglersPerCircuit {
		c.jugglers[j] = score
		if len(c.jugglers) == jugglersPerCircuit {
			c.updateMin()
		}
		return true, nil
	}
	if score <= c.jugglers[c.minScore] {
		return false, nil
	}
	delete(c.jugglers, c.minScore)
	c.jugglers[j] = score
	displaced := c.minScore
	c.updateMin()
	return true, displaced
}

// Maybe need a minheap here
func (c *circuit) updateMin() {
	first := true
	var minScore int
	for j, score := range c.jugglers {
		if first || score < minScore {
			c.minScore = j
			first = false
			minScore = score
		}
	}
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

type juggler struct {
	name    string
	h, e, p int
	prefs   []string
	prefPos int
}

var jugglers []*juggler

func addJuggler(line string) {
	f := strings.Fields(line)
	jugglers = append(jugglers, &juggler{
		name:  f[1],
		h:     atoi(f[2][2:]),
		e:     atoi(f[3][2:]),
		p:     atoi(f[4][2:]),
		prefs: strings.Split(f[5], ","),
	})
}

// assign calculates assignments, prints result
func assign() {
	if len(jugglers)%len(circuits) != 0 {
		log.Fatal("doesn't divide", len(jugglers), len(circuits))
	}
	jugglersPerCircuit = len(jugglers) / len(circuits)

	dummy := struct{}{}

	unassigned := make(map[*juggler]struct{})
	for _, j := range jugglers {
		unassigned[j] = dummy
	}

	for len(unassigned) != 0 {
		log.Println("unassigned", len(unassigned), "assigned", countAssigned())
		for j := range unassigned {
			cName := j.nextCircuit()
			win, loser := circuits[cName].propose(j)
			if !win {
				continue
			}
			delete(unassigned, j)
			if loser != nil {
				unassigned[loser] = dummy
			}
		}
	}
	var total int
	c := circuits["C1970"]
	if c == nil {
		log.Fatal("no circuit 1970")
	}
	for j, _ := range c.jugglers {
		total += atoi(j.name[1:])
	}
	fmt.Println(total)

}

func (j *juggler) nextCircuit() string {
	if j.prefPos >= len(j.prefs){
		log.Fatal("prefs", j.prefs, "pos", j.prefPos)
	}
	name := j.prefs[j.prefPos]
	j.prefPos++
	return name
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
		line := scanner.Text()
		if len(line) < 1 {
			continue
		}
		switch line[0] {
		case 'C':
			addCircuit(line)
		case 'J':
			addJuggler(line)
		default:
			log.Fatal(line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	assign()
}
