//go:build OMIT

package main

import "fmt"

// Index vrátí index x v s, nebo -1, pokud ho nenajde.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v i x jsou typu T, který má omezení
		// comparable, takže tady můžeme použít ==.
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index funguje na výřezu intů
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index funguje také na výřezu řetězců
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}
