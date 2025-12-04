package numbers

// Erwartet eine Zahl n >= 1 und liefert die Anzahl der Teiler dieser Zahl zurück.
func CountDivisors(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count
}

// HINWEIS
// - Definieren Sie sich eine Zählvariable, die die Anzahl der Teiler speichert.
// - Nutzen Sie eine Schleife, die alle Zahlen von 1 bis n durchgeht.
//   Prüfen Sie mittels des "Modulo"-Operators `%`, ob die aktuelle Zahl ein Teiler von n ist.
//   - Der Modulo-Operator `%` gibt den Rest einer Division zurück.
//   - Eine Zahl i ist genau dann ein Teiler von n, wenn n%i == 0 ist.
