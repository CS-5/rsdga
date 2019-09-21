// Package rsdga (Really Simple Domain Generation Algorithm) is used to generate
// domain names based on the supplied date (year, month, day) and an optional
// seed. MD5 is used to hash the value of the supplied parameters and is
// returned with the supplied TLD appended.
//
// DGA
//
// Domain Genereration Algorithms are used to circumvent general domain blocklists
// by generating seemingly random domain names (In this case: using the current
// date along with an optional seed, hashed with MD5).
//
// Usage
//
// Basic code to initialize the generator and print the domains:
//
//	import (
//		"fmt"
//		"time"
//
//		"github.com/cs-5/rsdga"
//	)
//
//	func main() {
//		t := time.Now()
//
//		// Use the current time to supply the year, month, and day. Use ".com" as the TLD
//		gen := rsdga.New(t.Year(), int(t.Month()), t.Day(), "com")
//
//		// Print out 20 domains
//		for i := 0; i <= 20; i++ {
//			fmt.Println(gen.Next())
//		}
//	}
// Example output:
//	7cf3c19ef5871a8bdca3288bac26f615.com
//	813ed870de75f3605f7d4b2a0fe93608.com
//	6d06004feb188a4a463affbdbfff51d3.com
//	e285c151730d8a5ec59369d609df3e07.com
//	bec43b55c54522cfa12eaaef55c6f460.com
//	9a803c0135c0ef9b49cfce8ed7719039.com
//	e74dc2429f0094366fc5b2dcd47504f4.com
//	7f24abd8c954f7b965c85634357541b2.com
//	f7727c68543bc6b18c9720f01ba171aa.com
//	828f8bb747c58dfdde95cabebb0ebfb7.com
//	0804b15b01380b1da1785b24d18d7107.com
//	a2a1fe991980d30f11c8315047b74086.com
//	591bd58a478ec593b11a66274c783668.com
//	5f1cf69be4aeef07fc3ee8fc811eb9b3.com
//	8c4a87e79cd9d72bf999427c646dd550.com
//	79d93c774adf631aebf3ca573a4558de.com
//	d0a7c98f9afbe1cb0071b638c3f75fb9.com
//	733d865d7175c8737206bbeee97cbed7.com
//	4c77228996d2edbbe9092ce71f39e966.com
//	f76d2edeec5dc3a87f930444164dd38c.com
//	1ffd2643f72a6a57f65044442c59721f.com
package rsdga

/*
 * bulkssh.go by Carson Seese. Created: 09/18/2019. Modified: 09/21/2019.
 * DGA using a year, month, day and an optional seed, hashes with MD5.
 */

import (
	"crypto/md5"
	"fmt"
	"sync"
	"unicode"
)

// Generator contains the parameters required to generate domains. Use New() or
// NewSeeded() to initialize.
type Generator struct {
	year, month, day, seed, i int
	tld                       string
	lock                      *sync.Mutex
}

// New initializes a new Generator and returns it.
// Year, month, and day must all be in YYYY, MM, DD format (respectively) and
// TLD must only contain letters (No numbers or punctuation).
func New(year, month, day int, tld string) (*Generator, error) {
	return NewSeeded(year, month, day, 0, tld)
}

// NewSeeded initializes a new Generator with a seed and returns it. See New()
// for parameter descriptions.
func NewSeeded(year, month, day, seed int, tld string) (*Generator, error) {
	/* Make sure the TLD only contains letters */
	for _, r := range tld {
		if !unicode.IsLetter(r) {
			return new(Generator),
				fmt.Errorf("tld: %v contains invalid characters", tld)
		}
	}

	if year > 10000 {
		return new(Generator),
			fmt.Errorf("year: %v greater than 4 digits", year)
	}

	if month > 12 {
		return new(Generator),
			fmt.Errorf("month: %v > 12 doesn't exist", month)
	}

	if day > 31 {
		return new(Generator),
			fmt.Errorf("day: %v > 31 doesn't exist", day)
	}

	return &Generator{
		year:  year,
		month: month,
		day:   day,
		tld:   tld,
		seed:  seed,
		lock:  new(sync.Mutex),
	}, nil
}

// Next returns the generated domain as a string and increments the iterator
// MD5 is used to hash the generated string before adding the TLD and returning.
func (g *Generator) Next() string {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.i++

	return fmt.Sprintf("%x.%s", md5.Sum([]byte(
		fmt.Sprintf("%v%v%v%v%v", g.year, g.month, g.day, g.i, g.seed),
	)), g.tld)
}
