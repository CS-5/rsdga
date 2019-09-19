package rsdga

import (
	"crypto/md5"
	"fmt"
	"sync"
)

// Generator contains the parameters required to generate domains
type Generator struct {
	year, month, day, seed, i int
	tld                       string
	lock                      sync.Mutex
}

// New initializes a new Generator with the supplied year, month, day, and TLD
func New(year, month, day int, tld string) *Generator {
	return &Generator{
		year:  year,
		month: month,
		day:   day,
		tld:   tld,
	}
}

// NewSeeded initializes a new Generator with the supplied year, month, day, TLD, and seed
func NewSeeded(year, month, day, seed int, tld string) *Generator {
	return &Generator{
		year:  year,
		month: month,
		day:   day,
		tld:   tld,
		seed:  seed,
	}
}

// Next returns the generated domain and increments the iterator
func (g *Generator) Next() string {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.i++

	return fmt.Sprintf("%x.%s", md5.Sum([]byte(fmt.Sprintf("%v%v%v%v%v", g.year, g.month, g.day, g.i, g.seed))), g.tld)
}
