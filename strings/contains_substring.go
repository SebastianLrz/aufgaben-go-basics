package strings

// Erwartet zwei Strings s und t und prüft, ob t in s als Teilstring vorkommt.
func ContainsSubstring(s, t string) bool {
	for i := 0; i < len(s)-len(t)+1; i++ {
		if s[i:i+len(t)] == t {
			return true
		}
	}
	return false
}

// HINWEIS
// Laufen Sie in einer for-Schleife über alle Positionen in s.
// Prüfen Sie in jedem Schleifendurchlauf, ob s ab der aktuellen Position mit t beginnt.
// Verwenden Sie dazu den Slice-Operator, um einen Teilstring aus s zu extrahieren.
