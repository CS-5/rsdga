package rsdga_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/CS-5/rsdga"
)

func ExampleGenerator() {
	t := time.Now()

	// Use the current time to supply the year, month, and day. Use ".com" as the TLD
	gen, err := rsdga.New(t.Year(), int(t.Month()), t.Day(), "com")

	if err != nil {
		fmt.Println(err)
		return
	}

	// Print out 20 domains
	for i := 0; i <= 20; i++ {
		fmt.Println(gen.Next())
	}
}

func ExampleGenerator_seeded() {
	t := time.Now()

	// Use the current time to supply the year, month, day, and seed. Use ".com" as the TLD
	gen, err := rsdga.NewSeeded(t.Year(), int(t.Month()), t.Day(), 1234, "com")

	if err != nil {
		fmt.Println(err)
		return
	}

	// Print out 20 domains
	for i := 0; i <= 20; i++ {
		fmt.Println(gen.Next())
	}
}

func TestGenerator(t *testing.T) {
	/* Initialize Generator */
	gen, err := rsdga.New(2000, 1, 1, ".com")
	if err != nil {
		t.Error(err.Error())
	}

	/* Get a domain from the generator and compare */
	test := gen.Next()
	control := "0da796a9339d6d45170010aa06375d91.com"

	if test != control {
		t.Errorf("Domain incorrect, have: %q, want: %q", test, control)
	}
}
