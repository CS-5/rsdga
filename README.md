# RSDGA

**R**eally **S**imple **D**omain **G**eneration **A**lgorithm

Short and to the point Go library for generating domain names based on supplied parameters.

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

	// Use the current time to supply the year, month, and day. Use ".com" as the TLD
	gen := rsdga.New(t.Year(), int(t.Month()), t.Day(), "com")

	// Print out 20 domains
	for i := 0; i <= 20; i++ {
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
9a803c0135c0ef9b49cfce8ed7719039.com
e74dc2429f0094366fc5b2dcd47504f4.com
7f24abd8c954f7b965c85634357541b2.com
f7727c68543bc6b18c9720f01ba171aa.com
828f8bb747c58dfdde95cabebb0ebfb7.com
0804b15b01380b1da1785b24d18d7107.com
a2a1fe991980d30f11c8315047b74086.com
591bd58a478ec593b11a66274c783668.com
5f1cf69be4aeef07fc3ee8fc811eb9b3.com
8c4a87e79cd9d72bf999427c646dd550.com
79d93c774adf631aebf3ca573a4558de.com
d0a7c98f9afbe1cb0071b638c3f75fb9.com
733d865d7175c8737206bbeee97cbed7.com
4c77228996d2edbbe9092ce71f39e966.com
f76d2edeec5dc3a87f930444164dd38c.com
1ffd2643f72a6a57f65044442c59721f.com
```

### Make a new Generator

Without a seed:
`generator := rsdga.New(2019, 01, 01, "com")`

With a seed:
`generator := rsdga.NewSeeded(2019, 01, 01, 1234, "com")`

### Get Domain

`domain := generator.Next()`
