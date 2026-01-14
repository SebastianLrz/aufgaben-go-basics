package numbers

import "math"

// Erwartet eine Zahl n und prüft, ob n eine Primzahl ist.
func IsPrime(n int) bool {
	teilerAnzahl := 0

	for i := 1; i <= n; i++ {

		x := float64(n) / float64(i)
		if x == math.Floor(x) { //math.floor rundet ab, kann aber nur float64
			teilerAnzahl++
		}
	}

	if teilerAnzahl == 2 {
		return true
	} else {
		return false
	}
}
