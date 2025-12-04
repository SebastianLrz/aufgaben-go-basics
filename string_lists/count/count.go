package count

// Erwartet eine Liste von Strings sowie einen String, der gezählt werden soll.
// Liefer die Anzahl der Vorkommen des gesuchten Strings in der Liste.
func Count(strings []string, search string) int {
	count := 0
	for _, s := range strings {
		if s == search {
			count++
		}
	}
	return count
}

// HINWEIS
// - Nutzen Sie eine Schleife, um durch die Liste der Strings zu iterieren.
// - Nutzen Sie außerdem eine `int`-Variable, um die Anzahl der Vorkommen zu zählen.
//   Diese sollte vor der Schleife mit 0 initialisiert
//   und bei jedem Vorkommen des gesuchten Strings um 1 erhöht werden.
// TASK_NOTES
// - einfachere Variante: count
// TAGS: [Arrays/Slices, Loops]
