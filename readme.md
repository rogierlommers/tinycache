## tinycache
a cache package to keep the last n elements of any struct types

```
package main

import (
	"fmt"
	"time"

	"github.com/rogierlommers/tinycache"
)

// define your custom struct, you can store everything you want
type download struct {
	timestamp time.Time
	uid       string
}

func main() {
	// create a cache instance
	h := tinycache.NewCache()

	// set max number of elements to keep
	h.SetMaxElements(10)

	// add 20 elements
	for i := 0; i < 20; i++ {
		element := download{
			uid:       fmt.Sprintf("uid-%d", i),
			timestamp: time.Now(),
		}
		h.Add(element)
	}

	// now only the last 10 are kept
	for _, d := range h.GetElements() {
		fmt.Printf("uid: %s, timestamp: %s\n", d.(download).uid, d.(download).timestamp)
	}

}
```