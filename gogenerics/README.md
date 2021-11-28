
#### Generic example:

```
package main

import (
	"fmt"
	"strconv"
)

func Map[A, B any](as []A, f func(A) B) []B {
	var bs []B
	for _, a := range as {
		b := f(a)
		bs = append(bs, b)
	}
	return bs
}

func to_i64(i int) int64 {
	return int64(i + i)
}

func to_s(i int) string {
	s := strconv.Itoa(i)
	return s + s
}

func main() {
	ints := []int{1, 2}
	ints64 := Map(ints, to_i64)
	fmt.Println(ints64)

	strs := Map(ints, to_s)
	fmt.Println(strs)

	return
}

```