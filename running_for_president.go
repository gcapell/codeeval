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

var (
	issueCount int
)

func scan(s *bufio.Scanner) string {
	if !s.Scan() {
		log.Fatal("short file")
	}
	return s.Text()
}

func readHeader(s *bufio.Scanner) {
	line := scan(s)
	chunks := strings.Fields(line)
	issueCount = atoi(chunks[2])
	log.Println("issueCount", issueCount)
	blank(s)
}

func blank(s *bufio.Scanner) {
	// consume blank
	line := scan(s)
	if len(line) != 0 {
		log.Fatalf("expected blank line, %q", line)
	}

}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

type issue uint32

var (
	nameToIssue = make(map[string]issue)
	issueCost   []int
)

func splitCost(line string) (string, int) {
	chunks := strings.Split(line, ": ")
	return chunks[0], atoi(strings.TrimSpace(chunks[1]))

}
func readIssues(s *bufio.Scanner) {
	for j := 0; j < issueCount; j++ {
		name, cost := splitCost(scan(s))
		nameToIssue[name] = issue(j)
		issueCost = append(issueCost, cost)
	}
	blank(s)
	log.Println("nameToIssue", nameToIssue)
	log.Println("issueCost", issueCost)
}

type state struct {
	name   string
	votes  int
	issues map[issue]int
}

func (s *state) String() string {
	return fmt.Sprintf("%s %d %v", s.name, s.votes, s.issues)
}

var states []*state

func readState(s *bufio.Scanner) bool {
	name := scan(s)
	_, votes := splitCost(scan(s))
	state := &state{name, votes, make(map[issue]int)}
	states = append(states, state)
	for {
		if !s.Scan() {
			return false
		}
		line := s.Text()
		if len(line) == 0 {
			return true
		}
		name, votes := splitCost(line)
		issue := nameToIssue[name]
		state.issues[issue] = votes
	}
}

type bitmap issue

type platform struct{
	next, prev *platform
	issues bitmap
	electoralVotes int
	stateVotes [50]int
}

func (p *platform) winning() bool {
	return p.electoralVotes >= 270
}

type queue struct {
	head, tail *platform
}

func (q *queue) push(p *platform) {
	p.next = q.head
	if q.head != nil {
		q.head.prev = p
	}
	if q.tail == nil {
		q.tail = p
	} 
}

func (q *queue) empty() bool {
	return q.tail == nil
}

func (q *queue) pop() *platform {
	reply := q.tail
	q.tail := reply.prev
	if q.tail != nil {
		q.tail.next = nil
	}
	return reply
}

func newPlatform(i issue) *platform {
	p := &platform {}
	p.addIssue(i)
	return p
}

func (p *platform) addIssue(i issue){
	p.issues != 1<<i
	for state, votes := range stateVotesByIssue[i] {
		p.stateVotes[state] += votes
		s := states[state]
		if s.victory(p.stateVotes[state], votes){
			p.electoralVotes += s.votes
		}
	}
}

func newSentinel() *platform {
	return &platform{}
}

func (p *platform) isSentinel() bool {
	return p.issues == 0
}

func (p *platform) children() []*platform {
	return nil
}

func chooseIssues() []*platform {
	log.Println("states", states)
	var q queue
	
	for j := 0; j<issueCount; j++ {
		q.push(newPlatform(issue(j)))
	}
	
	q.push(newSentinel())
	var solutions []*platform
	for {
		p := q.pop()
		if p.isSentinel() {
			if len(solutions) != 0 {
				return solutions
			}
			if q.empty() {
				log.Fatal("we lose")
			}
			q.push(newSentinel())
			continue
		}
		if p.winning() {
			solutions = append(solutions, p)
		}
		if len(solutions) > 0 {
			continue
		}
		for _, o := range p.children() {
			q.push(o)
		}
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

	scanner := bufio.NewScanner(f)
	readHeader(scanner)
	readIssues(scanner)
	for readState(scanner) {
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err, "expected EOF")
	}
	chooseIssues()
}
