package rsdga_test

import (
	"crypto/md5"
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
	time := time.Now()

	/* Initialize Generator */
	gen, err := rsdga.New(time.Year(), int(time.Month()), time.Day(), "com")
	if err != nil {
		t.Error(err.Error())
	}

	/* Iterate 5 Times */
	for i := 1; i <= 6; i++ {
		genDomain := gen.Next()
		testDomain := fmt.Sprintf("%x.%s", md5.Sum([]byte(fmt.Sprintf(
			"%v%v%v%v%v", time.Year(), int(time.Month()), time.Day(), i, 0),
		)), "com")

		if genDomain != testDomain {
			t.Errorf("Domain incorrect, got: %q, want: %q", genDomain, testDomain)
			break
		}
	}
}
