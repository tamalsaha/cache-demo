package main

import (
	"fmt"
	"time"

	"github.com/go-pkgz/expirable-cache/v2"
)

func main() {
	// make cache with short TTL and 3 max keys
	c := cache.NewCache[string, string]().WithMaxKeys(3).WithTTL(time.Millisecond * 10)

	// set value under key1.
	// with 0 ttl (last parameter) will use cache-wide setting instead (10ms).
	c.Set("key1", "val1", 0)

	// get value under key1
	r, ok := c.Get("key1")

	// check for OK value, because otherwise return would be nil and
	// type conversion will panic
	if ok {
		rstr := r // convert cached value from interface{} to real type
		fmt.Printf("value before expiration is found: %v, value: %v\n", ok, rstr)
	}

	time.Sleep(time.Millisecond * 11)

	// get value under key1 after key expiration
	r, ok = c.Get("key1")
	// don't convert to string as with ok == false value would be nil
	fmt.Printf("value after expiration is found: %v, value: %v\n", ok, r)

	// set value under key2, would evict old entry because it is already expired.
	// ttl (last parameter) overrides cache-wide ttl.
	c.Set("key2", "val2", time.Minute*5)

	fmt.Printf("%+v\n", c)
	// Output:
	// value before expiration is found: true, value: val1
	// value after expiration is found: false, value: <nil>
	// Size: 1, Stats: {Hits:1 Misses:1 Added:2 Evicted:1} (50.0%)
}
