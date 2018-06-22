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
	// create a cache instance with
	// a max of 10 elements
	h := tinycache.NewCache(10)

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
	// since we store interfaces, you
	// should cast them elements back
	// to the original struct
	
	for _, d := range h.GetElements() {
		fmt.Printf("uid: %s, timestamp: %s\n", d.(download).uid, d.(download).timestamp)
	}

}
```

## output

```
uid: uid-10, timestamp: 2018-06-18 22:16:28.329057694 +0200 CEST m=+0.000420022
uid: uid-11, timestamp: 2018-06-18 22:16:28.329058014 +0200 CEST m=+0.000420342
uid: uid-12, timestamp: 2018-06-18 22:16:28.329058255 +0200 CEST m=+0.000420583
uid: uid-13, timestamp: 2018-06-18 22:16:28.329058511 +0200 CEST m=+0.000420839
uid: uid-14, timestamp: 2018-06-18 22:16:28.329058752 +0200 CEST m=+0.000421080
uid: uid-15, timestamp: 2018-06-18 22:16:28.329058987 +0200 CEST m=+0.000421315
uid: uid-16, timestamp: 2018-06-18 22:16:28.329059211 +0200 CEST m=+0.000421539
uid: uid-17, timestamp: 2018-06-18 22:16:28.329059444 +0200 CEST m=+0.000421772
uid: uid-18, timestamp: 2018-06-18 22:16:28.329059684 +0200 CEST m=+0.000422012
uid: uid-19, timestamp: 2018-06-18 22:16:28.329059923 +0200 CEST m=+0.000422251
```