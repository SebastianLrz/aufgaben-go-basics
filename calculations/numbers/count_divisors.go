package numbers

import (
	"math"
)

// Erwartet eine Zahl n >= 1 und liefert die Anzahl der Teiler dieser Zahl zurück.
func CountDivisors(n int) int {

	teilerAnzahl := 0

	for i := 1; i <= n; i++ {

		x := float64(n) / float64(i)
		if x == math.Floor(x) {
			teilerAnzahl++
		}
	}

	return teilerAnzahl

}
