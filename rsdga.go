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
//		gen, err := rsdga.New(t.Year(), int(t.Month()), t.Day(), "com")
//
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//
//		// Print out 5 domains
//		for i := 0; i <= 5; i++ {
//			fmt.Println(gen.Next())
//		}
//	}
// Example output:
//	7cf3c19ef5871a8bdca3288bac26f615.com
//	813ed870de75f3605f7d4b2a0fe93608.com
//	6d06004feb188a4a463affbdbfff51d3.com
//	e285c151730d8a5ec59369d609df3e07.com
//	bec43b55c54522cfa12eaaef55c6f460.com
package rsdga

/*
 * bulkssh.go by Carson Seese. Created: 09/18/2019. Modified: 09/21/2019.
 * DGA using a year, month, day and an optional seed, hashes with MD5.
 */

import (
	"crypto/md5"
	"fmt"
	"strings"
	"sync"
)

// Generator contains the parameters required to generate domains. Use New() or
// NewSeeded() to initialize.
type Generator struct {
	year, month, day, seed, i int
	tld                       string
	lock                      *sync.Mutex
}

// New initializes a new Generator and returns it.
// Year, month, and day must all be in YYYY, MM, DD format (respectively).
// Note: There is no input validation here.
func New(year, month, day int, tld string) (*Generator, error) {
	return NewSeeded(year, month, day, 0, tld)
}

// NewSeeded initializes a new Generator with a seed and returns it. See New()
// for parameter descriptions.
func NewSeeded(year, month, day, seed int, tld string) (*Generator, error) {
	if !strings.HasPrefix(tld, ".") {
		tld = "." + tld
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

	fmt.Println(
		fmt.Sprintf("%v%v%v%v%v", g.year, g.month, g.day, g.i, g.seed))

	return fmt.Sprintf("%x%s", md5.Sum([]byte(
		fmt.Sprintf("%v%v%v%v%v", g.year, g.month, g.day, g.i, g.seed),
	)), g.tld)
}
