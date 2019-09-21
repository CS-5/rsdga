# RSDGA

**R**eally **S**imple **D**omain **G**eneration **A**lgorithm

Short and to the point Go library for generating domain names based on supplied parameters. A Domain Generation Algorithm is used to circumvent general domain blocklists by generating seemingly random domain names (In this case: using the current date along with an optional seed, hashed with MD5). A better description and overview can be found [here](https://blog.malwarebytes.com/security-world/2016/12/explained-domain-generating-algorithm/).

Written for the _Hands-on Writing Malware in Go_ talk at BSidesDC 2019.

For legal use only.

## Usage

Pretty simple (One might say _really_ simple):

```go
import (
	"fmt"
	"time"

	"github.com/cs-5/rsdga"
)

func main() {
	t := time.Now()

	/* Use the current time to supply the year, month, and day. Use ".com" as the TLD */
	gen, err := rsdga.New(t.Year(), int(t.Month()), t.Day(), "com")

	if err != nil {
		fmt.Println(err)
		return
	}

	/* Print out 5 domains */
	for i := 0; i <= 5; i++ {
		fmt.Println(gen.Next())
	}
}
```

Output:

```
7cf3c19ef5871a8bdca3288bac26f615.com
813ed870de75f3605f7d4b2a0fe93608.com
6d06004feb188a4a463affbdbfff51d3.com
e285c151730d8a5ec59369d609df3e07.com
bec43b55c54522cfa12eaaef55c6f460.com
```

### Make a new Generator

Without a seed:
`generator := rsdga.New(2019, 01, 01, "com")`

With a seed:
`generator := rsdga.NewSeeded(2019, 01, 01, 1234, "com")`

### Get Domain

`domain := generator.Next()`
