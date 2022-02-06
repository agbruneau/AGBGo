/* A Tour of Go Exercise: Fibonacci closure */
package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	defer timeTrack(time.Now(), "Main")
	f := fibonacci()
	for i := 0; i < 1000; i++ {
		fmt.Println(f())
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

/*
Solution 1: */
func fibonacci() func() int {
	var v1, v2 int = 0, 1
	return func() int {
		v2 = v2 + v1
		v1 = v2 - v1
		return v2 - v1
	}
}

/*
Solution 2:
func fibonacci() func() int {
	v1, v2 := 0, 1

	return func() int {
		defer func() {
			v1, v2 = v2, v1+v2
		}()
		return v1
	}
}
*/

/*
Solution 3:
func fibonacci() func() int {
	previous := 0
	last := 0
	index := 0

	return func() int {
		var result int
		if index == 1 {
			result = 1
		} else {
			result = last + previous
		}

		previous = last
		last = result
		index++

		return int(result)
	}
}
*/

/*
Solution 4:
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		result := a // Store the result value before translations
		c := a + b  // Fibonacci
		a = b       // Translation
		b = c       // Translation
		return result
	}
}
*/

/*
Solution 5 :
func fibonacci() func() int {
	n := 0
	a := 0
	b := 1
	c := a + b
	return func() int {
		var ret int
		switch {
		case n == 0:
			n++
			ret = 0
		case n == 1:
			n++
			ret = 1
		default:
			ret = c
			a = b
			b = c
			c = a + b
		}
		return ret
	}
}
*/

/*
Solution 6, with Slice:
func fibonacci() func() int {
	var i []int
	return func() int {
		if len(i) == 0 {
			i = append(i, 0)
		} else if len(i) == 1 {
			i = append(i, 1)
		} else {
			i = append(i, i[len(i)-2]+i[len(i)-1])
		}
		return i[len(i)-1]
	}
}
*/
