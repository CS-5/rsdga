# RSDGA
**R**eally **S**imple **D**omain **G**eneration **A**lgorithm library

Short and to the point Go library for generating domain names based on supplied parameters. 

Written for the _Hands-on Writing Malware in Go_ talk at BSidesDC 2019. 

For legal use only.

## Usage

Pretty simple:

```go
import (
	"fmt"
	"time"

	"github.com/cs-5/rsdga"
)

func main() {
	t := time.Now()

	// Use the current time to supply the year, month, and day. Use ".tk" as the TLD
	gen := rsdga.New(t.Year(), int(t.Month()), t.Day(), "tk")

	// Print out 20 domains
	for i := 0; i <= 20; i++ {
		fmt.Println(gen.Next())
	}
}
```
