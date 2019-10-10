package rsdga_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/CS-5/rsdga"
)

var tests = []struct {
	input  [4]int
	domain string
}{
	{[4]int{2000, 1, 1, 0}, "bb601423a9b2055e963831c11071b23f.test"},
	{[4]int{2000, 12, 12, 1234}, "0faac41958610de98ba4a984073a6b1e.test"},
	{[4]int{2001, 1, 1, 0}, "75ad91624945b95701299845ca4892ed.test"},
	{[4]int{2001, 12, 12, 1234}, "99cdb2885d3fe41bc2db260a1f2ce50b.test"},
	{[4]int{0, 0, 0, 0}, "4c68cea7e58591b579fd074bcdaff740.test"},
	{[4]int{1111, 1, 1, 1}, "1bbd886460827015e5d605ed44252251.test"},
	{[4]int{2000, 2, 2, 0}, "fed7023b31e2e8d0474fd6b4fdf5a9d2.test"},
	{[4]int{2001, 3, 2, 1}, "45db80515e9440182a0b8c80bf4878a5.test"},
	{[4]int{2002, 4, 2, 2}, "c62b96ebfc9c775eb47e729903d399e4.test"},
	{[4]int{2003, 5, 2, 3}, "ea6bbe7b9f2344c46480ea79d54e77e1.test"},
}

func ExampleGenerator() {
	t := time.Now()

	/* Use the current time to supply the year, month, day, and seed. Use ".com" as the TLD */
	gen := rsdga.New(t.Year(), int(t.Month()), t.Day(), ".com")

	/* Print out 5 domains */
	for i := 0; i < 5; i++ {
		fmt.Println(gen.Next())
	}
}

func ExampleGenerator_seeded() {
	t := time.Now()

	/* Use the current time to supply the year, month, day, and 1234 as the seed. Use ".com" as the TLD */
	gen := rsdga.NewSeeded(t.Year(), int(t.Month()), t.Day(), 1234, ".com")

	/* Print out 5 domains */
	for i := 0; i < 5; i++ {
		fmt.Println(gen.Next())
	}
}

func TestGenerator(t *testing.T) {
	for _, test := range tests {
		gen := rsdga.NewSeeded(
			test.input[0],
			test.input[1],
			test.input[2],
			test.input[3],
			".test",
		)

		generated := gen.Next()
		control := test.domain

		if generated != control {
			t.Errorf("Domain incorrect, have: %q, want %q", generated, control)
			break
		}
	}
}
